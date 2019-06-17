package events

import (
	"encoding/json"
	"log"

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
			log.Printf("ERROR: failed to marshal data for %s.%s event", evt.Quake.PublicID, evt.Type)
			return
		}
		err = nc.Publish("geonet.quakes."+evt.Type, data)
	}
}
