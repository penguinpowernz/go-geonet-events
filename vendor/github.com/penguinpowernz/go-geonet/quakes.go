package geonet

import (
	"fmt"
	"time"
)

const DefaultMMI = 0

type Quake struct {
	QuakeProperties
	Coordinates []float64 `json:"coordinates"`
}

type FeltReport struct {
	IntensityProperties
	Coordinates []float64 `json:"coordinates"`
}

func (q Quake) FeltReports(c *Client) ([]FeltReport, error) {
	resp := IntensityAPIResponse{}
	err := c.Get("/intensity?type=reported&publicID="+q.PublicID, &resp)
	if err != nil {
		return []FeltReport{}, err
	}

	return resp.Reports(), nil
}

func (c *Client) Quake(id string) (Quake, error) {
	resp := QuakeAPIResponse{}
	if err := c.Get(sfmt("/quake/%s", id), &resp); err != nil {
		return Quake{}, err
	}

	if len(resp.Quakes()) == 0 {
		return Quake{}, fmt.Errorf("quake could not be parsed")
	}

	return resp.Quakes()[0], nil
}

func (c *Client) Quakes(mmis ...int) ([]Quake, error) {
	mmi := DefaultMMI
	if len(mmis) > 0 {
		mmi = mmis[0]
	}

	resp := QuakeAPIResponse{}
	if err := c.Get(sfmt("/quake?MMI=%d", mmi), &resp); err != nil {
		return []Quake{}, err
	}

	return resp.Quakes(), nil
}

type QuakeAPIResponse struct {
	Type     string         `json:"type"`
	Features []QuakeFeature `json:"features"`
}

type IntensityAPIResponse struct {
	Type     string             `json:"type"`
	Features []IntensityFeature `json:"features"`
}

type Feature struct {
	Type     string   `json:"type"`
	Geometry Geometry `json:"geometry"`
}

type QuakeFeature struct {
	Feature
	Properties QuakeProperties `json:"properties"`
}

type IntensityFeature struct {
	Feature
	Properties IntensityProperties `json:"properties"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type QuakeProperties struct {
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

type IntensityProperties struct {
	MMI   int `json:"mmi"`
	Count int `json:"count"`
}

func (res QuakeAPIResponse) Quakes() []Quake {
	quakes := []Quake{}

	for _, qf := range res.Features {
		q := Quake{qf.Properties, qf.Geometry.Coordinates}
		quakes = append(quakes, q)
	}

	return quakes
}

func (res IntensityAPIResponse) Reports() []FeltReport {
	reports := []FeltReport{}

	for _, f := range res.Features {
		r := FeltReport{f.Properties, f.Geometry.Coordinates}
		reports = append(reports, r)
	}

	return reports
}
