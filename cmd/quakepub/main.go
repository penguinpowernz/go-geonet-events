package main

import (
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/ghodss/yaml"
	nats "github.com/nats-io/go-nats"

	events "github.com/penguinpowernz/go-geonet-events"
)

type config struct {
	NATSURL string `json:"nats_url"`
	MQTTURL string `json:"mqtt_url"`
}

func main() {
	var cfgFile string
	flag.StringVar(&cfgFile, "c", "", "the path to the config file")
	flag.Parse()

	cfg := config{}
	data, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("failed to read config file: %s", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to read config file: %s", err)
	}

	if cfg.NATSURL == "" && cfg.MQTTURL == "" {
		log.Fatalf("no output servers specified")
	}

	ntfr := &events.Notifier{}

	if cfg.NATSURL != "" {
		bus := createNATSBus(cfg.NATSURL)
		ntfr.AddBus(bus)
	}

	if cfg.MQTTURL == "" {
		bus := createMQTTBus(cfg.MQTTURL)
		ntfr.AddBus(bus)
	}

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

func createNATSBus(url string) events.EventBus {
	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}

	return events.NatsNotifier(nc)
}

func createMQTTBus(url string) events.EventBus {
	return func(events.Event) {

	}
}
