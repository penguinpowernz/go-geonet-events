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

## Payloads

There are two event types; `new` and `updated`. The latter is only sent when there is an update to a previous quake that was sent out.  This
is usually due to revisions to the magnitude and depth of the quakes, the `quality` field can be used to determine the usefulness of the event.

### New

```json
{
    "type": "new",
    "quake": {
        "publicID": "2020p203673",
        "time": "2020-03-16T08:37:10.376Z",
        "depth": 10.83877039,
        "magnitude": 1.021546187,
        "locality": "10 km north-east of Matawai",
        "mmi": -1,
        "quality": "preliminary",
        "coordinates": [
            177.6768036,
            -38.28015518
        ]
    }
}
```

### Updated

```json
{
    "type": "update",
    "quake": {
        "publicID": "2020p203673",
        "time": "2020-03-16T08:37:10.376Z",
        "depth": 23.87388039,
        "magnitude": 1.222926187,
        "locality": "10 km north-east of Matawai",
        "mmi": -1,
        "quality": "best",
        "coordinates": [
            177.6768036,
            -38.28015518
        ]
    },
    "updated_fields": [
        "quality",
        "magnitude",
        "depth"
    ]
}
```

# Usage

You can see the clients written in golang in the `cmd` directory.