package geonet

type intensityAPIResponse struct {
	Type     string             `json:"type"`
	Features []intensityFeature `json:"features"`
}

func (res intensityAPIResponse) Reports() []FeltReport {
	reports := []FeltReport{}

	for _, f := range res.Features {
		r := FeltReport{f.Properties, f.Geometry.Coordinates}
		reports = append(reports, r)
	}

	return reports
}

type intensityFeature struct {
	feature
	Properties intensityProperties `json:"properties"`
}

type intensityProperties struct {
	MMI   int `json:"mmi"`
	Count int `json:"count"`
}

// FeltReport represents a report from a user who felt shaking
type FeltReport struct {
	intensityProperties
	Coordinates []float64 `json:"coordinates"`
}

// FeltReports will return reports from people who felt the quake
// using the given API client
func (q Quake) FeltReports(c *Client) ([]FeltReport, error) {
	resp := intensityAPIResponse{}
	err := c.Get("/intensity?type=reported&publicID="+q.PublicID, &resp)
	if err != nil {
		return []FeltReport{}, err
	}

	return resp.Reports(), nil
}
