package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/ghodss/yaml"
	nats "github.com/nats-io/go-nats"

	events "github.com/penguinpowernz/go-geonet-events"
)

type config struct {
	NATSUser     string `json:"nats_user"`
	NATSPassword string `json:"nats_password"`
	NATSURL      string `json:"nats_url"`
	MQTTURL      string `json:"mqtt_url"`
	WSPort       string `json:"ws_port"`
}

func (c config) LoadFromFile(fn string) error {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &c); err != nil {
		return err
	}

	return nil
}

func (c *config) LoadFromEnv() error {
	c.MQTTURL = os.Getenv("MQTT_URL")
	c.NATSURL = os.Getenv("NATS_URL")
	c.NATSUser = os.Getenv("NATS_USER")
	c.NATSPassword = os.Getenv("NATS_PASSWORD")
	c.WSPort = os.Getenv("WS_PORT")
	return nil
}

func main() {
	var cfgFile string
	flag.StringVar(&cfgFile, "c", "", "the path to the config file")
	flag.Parse()

	var cfg config
	var err error
	if cfgFile == "" {
		err = cfg.LoadFromEnv()
	} else {
		err = cfg.LoadFromFile(cfgFile)
	}

	if err != nil {
		log.Fatalf("failed to read config file: %s", err)
	}

	if cfg.NATSURL == "" && cfg.MQTTURL == "" && cfg.WSPort == "" {
		log.Fatalf("no output servers specified")
	}

	ntfr := &events.Notifier{}

	if cfg.NATSURL != "" {
		bus := createNATSBus(cfg.NATSURL, cfg.NATSUser, cfg.NATSPassword)
		ntfr.AddBus(bus)
	}

	// if cfg.MQTTURL == "" {
	// 	bus := createMQTTBus(cfg.MQTTURL)
	// 	ntfr.AddBus(bus)
	// }

	if cfg.WSPort != "" {
		svr, bus := events.NewWebsocketNotifier(cfg.WSPort)
		ntfr.AddBus(bus)
		log.Printf("ERROR: Websocket server stopped: %s", svr.ListenAndServe())
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

func createNATSBus(url, user, pass string) events.EventBus {
	nc, err := nats.Connect(url, nats.UserInfo(user, pass))
	if err != nil {
		panic(err)
	}

	return events.NatsNotifier(nc)
}

// func createMQTTBus(url string) events.EventBus {
// 	return func(events.Event) {

// 	}
// }
