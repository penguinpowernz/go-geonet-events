package main

import (
	"encoding/json"
	"time"

	"github.com/cenkalti/backoff"
	nats "github.com/nats-io/go-nats"
	cache "github.com/patrickmn/go-cache"

	"github.com/penguinpowernz/go-geonet"
)

func main() {
	cl := geonet.NewClient()
	expo := backoff.NewExponentialBackOff()
	expo.MaxElapsedTime = time.Minute * 2

	quakes := cache.New(time.Hour*24, time.Hour*24)

	ntfr := &Notifier{}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
	}
	ntfr.buses = append(ntfr.buses, natsNotifier(nc))

	for {
		var err error
		var qks []geonet.Quake

		err = backoff.RetryNotify(func() error {
			qks, err = cl.Quakes()
			return err
		}, expo, func(err error, d time.Duration) {
			// TODO: log
		})

		for _, qk := range qks {
			qki, found := quakes.Get(qk.PublicID)
			if !found {
				quakes.SetDefault(qk.PublicID, qk)
				ntfr.Notify("new", qk)
				continue
			}

			xqk := qki.(geonet.Quake)

			fieldsUpdated := []string{}
			if xqk.Magnitude != qk.Magnitude {
				fieldsUpdated = append(fieldsUpdated, "magnitude")
			}

			if xqk.Depth != qk.Depth {
				fieldsUpdated = append(fieldsUpdated, "depth")
			}

			if xqk.Quality != qk.Quality {
				fieldsUpdated = append(fieldsUpdated, "quality")
			}

			if xqk.Locality != qk.Locality {
				fieldsUpdated = append(fieldsUpdated, "locality")
			}

			if xqk.MMI != qk.MMI {
				fieldsUpdated = append(fieldsUpdated, "mmi")
			}

			if len(fieldsUpdated) == 0 {
				continue
			}

			quakes.SetDefault(qk.PublicID, qk)
			ntfr.Notify("update", qk, fieldsUpdated)
		}

		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

type Notifier struct {
	buses []eventBus
}

func (ntfr *Notifier) Notify(t string, qk geonet.Quake, fieldsUpdated ...[]string) {
	evt := Event{
		Type:  t,
		Quake: qk,
	}

	if len(fieldsUpdated) > 0 {
		evt.UpdatedFields = fieldsUpdated[0]
	}
}

type Event struct {
	Type          string       `json:"type"`
	Quake         geonet.Quake `json:"quake"`
	UpdatedFields []string     `json:"updated_fields,omitempty"`
}

type eventBus func(Event)

func natsNotifier(nc *nats.Conn) eventBus {
	return func(evt Event) {
		data, err := json.Marshal(evt)
		if err != nil {

		}
		err = nc.Publish("geonet.quakes."+evt.Type, data)
	}
}
