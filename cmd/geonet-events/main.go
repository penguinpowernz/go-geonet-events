package main

import (
	"time"

	nats "github.com/nats-io/go-nats"

	events "github.com/penguinpowernz/go-geonet-events"
)

func main() {
	ntfr := &events.Notifier{}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	ntfr.AddBus(events.NatsNotifier(nc))

	getQuakes := events.NewQuakeGetter()
	processor := events.NewProcessor()

	for {
		qks, err := getQuakes()
		if err != nil {
			panic(err)
		}

		evts := processor.Process(qks)
		ntfr.Notify(evts...)

		time.Sleep(time.Second)
	}
}
