# Geonet Quake Events Forwarder

**This is completely unofficial and unaffiliated with Geonet**

This is a free service which forwards events from the [Geonet API](https://api.geonet.org.nz/) in an evented manner.  This means you
don't need to poll the [Geonet API](https://api.geonet.org.nz/) for quakes anymore, they can be delivered via the following protocols:

- [x] NATS via `nats://quakes.nz:4222` on subjects `geonet.quakes.new` and `geonet.quakes.updated` (username of `client` and empty password)
- [x] WebSockets via `ws://quakes.nz/events`
- [ ] MQTT via `mqtt://quakes.nz:8883` (COMING SOON)

## How does it work

Basically it does the polling for you, every second, of the Geonet API.  Then your scripts can just connect to one of the protocols at
[quakes.nz](http://quakes.nz) and sit back and wait to be alerted.

## Contributing

We would love and appreciate contributions to the code but would prefer if the only running instance was at [quakes.nz](http://quakes.nz) given
that many people running it could overload Geonets servers.