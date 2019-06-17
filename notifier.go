package events

import (
	"encoding/json"

	nats "github.com/nats-io/go-nats"
)

// Notifier sends events to event buses
type Notifier struct {
	buses []EventBus
}

// EventBus is a simple function that can send events on a bus
type EventBus func(Event)

func (ntfr *Notifier) AddBus(bus EventBus) {
	ntfr.buses = append(ntfr.buses, bus)
}

func (ntfr *Notifier) Notify(evts ...Event) {
	for _, x := range ntfr.buses {
		for _, e := range evts {
			x(e)
		}
	}
}

func NatsNotifier(nc *nats.Conn) EventBus {
	return func(evt Event) {
		data, err := json.Marshal(evt)
		if err != nil {

		}
		err = nc.Publish("geonet.quakes."+evt.Type, data)
	}
}
