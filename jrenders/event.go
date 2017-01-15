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
}

type EventParam struct {
	ShowHash bool
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

	// load event dates
	dao.Event.LoadEventDates(e)
	j.JEventDates = EventDate.List(e.EventDates, e)

	return j
}

var Event *event

func initEvent() {
	Event = &event{}
}
