# go-geonet-events

A server that polls the Geonet.org.nz API every second publishes new and updated quakes to NATS and MQTT.

## Building

Build it like so:

    go build ./cmd/quakepub

## Usage

Just run the server:

    ./quakepub
