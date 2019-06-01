package events

import geonet "github.com/penguinpowernz/go-geonet"

type Event struct {
	Type          string       `json:"type"`
	Quake         geonet.Quake `json:"quake"`
	UpdatedFields []string     `json:"updated_fields,omitempty"`
}

type Events []Event

func (evts Events) Add(t string, qk geonet.Quake, updatedFields ...[]string) {
	evt := Event{
		Type:  t,
		Quake: qk,
	}

	if len(updatedFields) > 0 {
		evt.UpdatedFields = updatedFields[0]
	}

	evts = append(evts, evt)
}
