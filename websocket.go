package events

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// NewWebsocketNotifier will create a new websocket notifier returning the server
// that will serve the websockets, and an EventBus to send events to
func NewWebsocketNotifier(port string) (*http.Server, EventBus) {
	mgr := &wsManager{}

	return &http.Server{
		Addr:           port,
		Handler:        mgr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}, mgr.handleEvent
}

type wsManager struct {
	conns []*websocket.Conn
}

func (mgr *wsManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		rootHandler(w, r)
	case "/events":
		mgr.webSocketHandler(w, r)
	default:
		http.NotFoundHandler().ServeHTTP(w, r)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", indexHTML)
}

func (mgr *wsManager) webSocketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	mgr.conns = append(mgr.conns, conn)
}

// handleEvent will handle the event passed to it by writing it out
// to all the connections.  It is an EventBus function
func (mgr *wsManager) handleEvent(evt Event) {
	for _, conn := range mgr.conns {
		if err := conn.WriteJSON(evt); err != nil {
			fmt.Println(err)
		}
	}
}
