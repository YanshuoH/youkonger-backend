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

	JParticipantUser
	UnavailableParticipantList []JParticipantUser `json:"unavailableParticipantList"`
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
		j.JParticipantUser = ParticipantUser.Itemize(p.ParticipantUser)
	}

	// load event dates
	dao.Event.LoadEventDates(e)
	j.JEventDates = EventDate.List(e.EventDates, e)

	puList, _ := dao.ParticipantUser.FindUnavailableByEventID(e.ID)
	j.UnavailableParticipantList = ParticipantUser.List(puList)

	return j
}

var Event *event

func initEvent() {
	Event = &event{}
}
