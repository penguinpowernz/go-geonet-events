package geonet

// Volcano represents a volcano status
type Volcano struct {
	volcanoProperties
	Coordinates []float64 `json:"coordinates"`
}

// Volcanos will return a list containing the status of
// all volcanos in the Zelandian continent
func (c *Client) Volcanos() ([]Volcano, error) {
	resp := volcanoAPIResponse{}
	if err := c.Get("/volcano/val", &resp); err != nil {
		return []Volcano{}, err
	}

	return resp.Volcanos(), nil
}

type volcanoProperties struct {
	// a unique identifier for the volcano
	ID string `json:"volcanoID"`

	// the volcano title
	Title string `json:"volcanoTitle"`

	// volcanic alert level
	Level int `json:"level"`

	// volcanic activity
	Activity string `json:"activity"`

	// most likely hazards
	Hazards string `json:"hazards"`
}

type volcanoFeature struct {
	feature
	Properties volcanoProperties `json:"properties"`
}

type volcanoAPIResponse struct {
	Type     string           `json:"type"`
	Features []volcanoFeature `json:"features"`
}

func (res volcanoAPIResponse) Volcanos() []Volcano {
	reports := []Volcano{}

	for _, f := range res.Features {
		r := Volcano{f.Properties, f.Geometry.Coordinates}
		reports = append(reports, r)
	}

	return reports
}
