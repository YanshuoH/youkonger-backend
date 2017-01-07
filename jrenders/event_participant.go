package jrenders

import "github.com/YanshuoH/youkonger/models"

type eventParticipant struct{}

type JEventParticipant struct {
	Name          string `json:"name"`
	EventDateUUID string `json:"eventDateUuid"`
}

type JEventParticipants struct {
	JList []JEventParticipant `json:"eventParticipantList"`
}

func (r *eventParticipant) Itemize(ep *models.EventParticipant, ed *models.EventDate) JEventParticipant {
	j := JEventParticipant{
		Name:          ep.Name,
		EventDateUUID: ed.UUID,
	}
	return j
}

func (r *eventParticipant) List(eps []models.EventParticipant, ed *models.EventDate) JEventParticipants {
	jList := make([]JEventParticipant, len(eps))
	for i, ep := range eps {
		jList[i] = r.Itemize(&ep, ed)
	}

	return JEventParticipants{JList: jList}
}

var EventParticipant *eventParticipant

func initEventParticipant() {
	EventParticipant = &eventParticipant{}
}
