package events

import (
	"time"

	cache "github.com/patrickmn/go-cache"
	geonet "github.com/penguinpowernz/go-geonet"
)

// NewProcessor will return a new processor with a cache already instantiated
func NewProcessor() *Processor {
	quakes := cache.New(time.Hour*24, time.Hour*24)
	return &Processor{quakes}
}

// Processor will process quakes given to it
type Processor struct {
	cache *cache.Cache
}

// Process will process new quakes and return events
func (pr *Processor) Process(qks []geonet.Quake) Events {
	evts := Events{}

	for _, qk := range qks {
		qki, found := pr.cache.Get(qk.PublicID)
		if !found {
			pr.cache.SetDefault(qk.PublicID, qk)
			evts.Add("new", qk)
			continue
		}

		xqk := qki.(geonet.Quake)

		fields := compareQuake(xqk, qk)
		if len(fields) == 0 {
			continue
		}

		pr.cache.SetDefault(qk.PublicID, qk)
		evts.Add("updated", qk, fields)
	}

	return evts
}

func compareQuake(a, b geonet.Quake) []string {
	fields := []string{}
	if a.Magnitude != b.Magnitude {
		fields = append(fields, "magnitude")
	}

	if a.Depth != b.Depth {
		fields = append(fields, "depth")
	}

	if a.Quality != b.Quality {
		fields = append(fields, "quality")
	}

	if a.Locality != b.Locality {
		fields = append(fields, "locality")
	}

	if a.MMI != b.MMI {
		fields = append(fields, "mmi")
	}

	return fields
}
