package geonet

type Volcano struct {
	VolcanoProperties
	Coordinates []float64 `json:"coordinates"`
}

type VolcanoProperties struct {
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

type VolcanoFeature struct {
	Feature
	Properties VolcanoProperties `json:"properties"`
}

type VolcanoAPIResponse struct {
	Type     string           `json:"type"`
	Features []VolcanoFeature `json:"features"`
}

func (c *Client) Volcanos() ([]Volcano, error) {
	resp := VolcanoAPIResponse{}
	if err := c.Get("/volcano/val", &resp); err != nil {
		return []Volcano{}, err
	}

	return resp.Volcanos(), nil
}

func (res VolcanoAPIResponse) Volcanos() []Volcano {
	reports := []Volcano{}

	for _, f := range res.Features {
		r := Volcano{f.Properties, f.Geometry.Coordinates}
		reports = append(reports, r)
	}

	return reports
}
