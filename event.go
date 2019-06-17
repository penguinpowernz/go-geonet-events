package events

import geonet "github.com/penguinpowernz/go-geonet"

// Event represents a new quake event.  It will have a type of either 'new', or 'updated'
// which idicates if it's a new quake or an update to an already delivered one
type Event struct {
	Type          string       `json:"type"`
	Quake         geonet.Quake `json:"quake"`
	UpdatedFields []string     `json:"updated_fields,omitempty"`
}

// Events is a collection of events
type Events []Event

func (evts Events) Add(t string, qk geonet.Quake, updatedFields ...[]string) {
// Add will add an event to a collection
	evt := Event{
		Type:  t,
		Quake: qk,
	}

	if len(updatedFields) > 0 {
		evt.UpdatedFields = updatedFields[0]
	}

	evts = append(evts, evt)
}
