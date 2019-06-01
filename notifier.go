package events

import (
	"encoding/json"

	nats "github.com/nats-io/go-nats"
)

type Notifier struct {
	buses []EventBus
}

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
