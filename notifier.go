package events

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	nats "github.com/nats-io/go-nats"
)

// Notifier sends events to event buses
type Notifier struct {
	buses []EventBus
}

// EventBus is a simple function that can send events on a bus
type EventBus func(Event)

// AddBus will add an event bus to the notifier
func (ntfr *Notifier) AddBus(bus EventBus) {
	ntfr.buses = append(ntfr.buses, bus)
}

// Notify will send a quake notification to all event buses attached to the notifier
func (ntfr *Notifier) Notify(evts ...Event) {
	for _, x := range ntfr.buses {
		for _, e := range evts {
			x(e)
		}
	}
}

// NatsNotifier will return an events bus that publishes events through a NATS connection
func NatsNotifier(nc *nats.Conn) EventBus {
	return func(evt Event) {
		data, err := json.Marshal(evt)
		if err != nil {
			log.Printf("ERROR: failed to marshal data for %s.%s event: %s", evt.Quake.PublicID, evt.Type, err)
			return
		}
		if err = nc.Publish("geonet.quakes."+evt.Type, data); err != nil {
			log.Printf("ERROR: failed send %s.%s event %s", evt.Quake.PublicID, evt.Type, err)
		}
	}
}

// MQTTNotifier returns an event bus that forwards events over MQTT
func MQTTNotifier(cl mqtt.Client) EventBus {
	return func(evt Event) {
		data, err := json.Marshal(evt)
		if err != nil {
			log.Printf("ERROR: failed to marshal data for %s.%s event: %s", evt.Quake.PublicID, evt.Type, err)
			return
		}
		t := cl.Publish("geonet/events/"+evt.Type, 0, false, data)
		if t.WaitTimeout(time.Second) && t.Error() != nil {
			log.Printf("ERROR: failed send %s.%s event %s", evt.Quake.PublicID, evt.Type, t.Error())
		}
	}
}

// WebsocketNotifier returns an event bus that forwards events over websockets
func WebsocketNotifier(svr *http.Server) EventBus {
	return svr.Handler.(*wsManager).handleEvent
}
