package geonet

import (
	"fmt"
	"time"
)

// DefaultMMI is the mercali modified index score that
// is used by default in API calls where none is specified
const DefaultMMI = 2

// Quake represents an Earthquake event
type Quake struct {
	quakeProperties
	Coordinates []float64 `json:"coordinates"`
}

type quakeProperties struct {
	// the unique public identifier for this quake
	PublicID string `json:"publicID"`

	// the origin time of the quake
	Time time.Time `json:"time"`

	// the depth of the quake in km
	Depth float64 `json:"depth"`

	// the summary magnitude for the quake
	Magnitude float64 `json:"magnitude"`

	// distance and direction to the nearest locality
	Locality string `json:"locality"`

	// the calculated MMI shaking at the closest locality in the New Zealand region, -1..8
	MMI int `json:"mmi"`

	// the quality of this information; best, good, caution, deleted
	Quality string `json:"quality"`
}

// Quake will get a specific quake by the ID
func (c *Client) Quake(id string) (Quake, error) {
	resp := quakeAPIResponse{}
	if err := c.Get(sfmt("/quake/%s", id), &resp); err != nil {
		return Quake{}, err
	}

	if len(resp.Quakes()) == 0 {
		return Quake{}, fmt.Errorf("quake could not be parsed")
	}

	return resp.Quakes()[0], nil
}

// Quakes will return any recent quakes with the given MMI or higher
func (c *Client) Quakes(mmis ...int) ([]Quake, error) {
	mmi := DefaultMMI
	if len(mmis) > 0 {
		mmi = mmis[0]
	}

	resp := quakeAPIResponse{}
	if err := c.Get(sfmt("/quake?MMI=%d", mmi), &resp); err != nil {
		return []Quake{}, err
	}

	return resp.Quakes(), nil
}

type quakeAPIResponse struct {
	Type     string         `json:"type"`
	Features []quakeFeature `json:"features"`
}

func (res quakeAPIResponse) Quakes() []Quake {
	quakes := []Quake{}

	for _, qf := range res.Features {
		q := Quake{qf.Properties, qf.Geometry.Coordinates}
		quakes = append(quakes, q)
	}

	return quakes
}

type feature struct {
	Type     string   `json:"type"`
	Geometry geometry `json:"geometry"`
}

type quakeFeature struct {
	feature
	Properties quakeProperties `json:"properties"`
}

type geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
