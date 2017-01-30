package jrenders

import (
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
)

type event struct{}

type JEvent struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	AdminHash   string `json:"hash"`
	JEventDates

	Name                string `json:"name,omitempty"`
	ParticipantUserUUID string `json:"participantUserUuid,omitempty"`
}

type EventParam struct {
	ShowHash        bool
	ParticipantUser *models.ParticipantUser
}

func (r *event) Itemize(e *models.Event, p EventParam) JEvent {
	j := JEvent{
		UUID:        e.UUID,
		Title:       e.Title,
		Description: e.Description,
		Location:    e.Location,
	}

	if p.ShowHash {
		j.AdminHash = e.AdminHash
	}
	if p.ParticipantUser != nil {
		j.Name = p.ParticipantUser.Name
		j.ParticipantUserUUID = p.ParticipantUser.UUID
	}

	// load event dates
	dao.Event.LoadEventDates(e)
	j.JEventDates = EventDate.List(e.EventDates, e)

	return j
}

var Event *event

func initEvent() {
	Event = &event{}
}
