package events

import (
	"time"

	"github.com/cenkalti/backoff"
	geonet "github.com/penguinpowernz/go-geonet"
)

// NewQuakeGetter will return a function that can be called to get
// a list of quakes from the API
func NewQuakeGetter() func() ([]geonet.Quake, error) {
	cl := geonet.NewClient()
	expo := backoff.NewExponentialBackOff()
	expo.MaxElapsedTime = time.Minute * 2

	return func() ([]geonet.Quake, error) {
		var err error
		var qks []geonet.Quake

		err = backoff.RetryNotify(func() error {
			qks, err = cl.Quakes()
			return err
		}, expo, func(err error, d time.Duration) {
			// TODO: log
		})

		return qks, err
	}
}
