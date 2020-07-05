package main

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://quakes.nz:4222", nats.UserInfo("client", "quakes"))
	if err != nil {
		panic(err)
	}

	sub, err := nc.SubscribeSync("geonet.quakes.>")
	if err != nil {
		panic(err)
	}

	for {
		msg, err := sub.NextMsgWithContext(context.Background())
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		fmt.Printf("Received: %s", msg.Data)
	}
}
