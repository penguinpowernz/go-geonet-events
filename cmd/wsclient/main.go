package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://quakes.nz/events"
	ws, _, err := new(websocket.Dialer).Dial(url, http.Header{"Origin": []string{"http://quakes.nz"}})
	if err != nil {
		panic(err)
	}

	for {
		msg := map[string]interface{}{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}

		fmt.Printf("Received: %+v.\n", msg)
	}
}
