package jrenders

import (
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	"time"
)

type eventDate struct{}

type JEventDate struct {
	UUID       string    `json:"uuid"`
	EventUUID  string    `json:"eventUuid"`
	Time       time.Time `json:"time"`
	TimeInUnix int64     `json:"timeInUnix"`
	IsDDay     bool      `json:"isDDay"`
	JEventParticipants
}

type JEventDates struct {
	JList []JEventDate `json:"eventDateList"`
}

func (r *eventDate) Itemize(ed *models.EventDate, e *models.Event) JEventDate {
	j := JEventDate{
		UUID:       ed.UUID,
		Time:       ed.Time,
		TimeInUnix: ed.Time.Unix(),
		IsDDay:     ed.IsDDay,
		EventUUID:  e.UUID,
	}

	dao.EventDate.LoadEventParticipants(ed)
	j.JEventParticipants = EventParticipant.List(ed.EventParticipants, ed)

	return j
}

func (r *eventDate) List(eds []models.EventDate, e *models.Event) JEventDates {
	jList := make([]JEventDate, len(eds))
	for i, ed := range eds {
		jList[i] = r.Itemize(&ed, e)
	}

	return JEventDates{JList: jList}
}

var EventDate *eventDate

func initEventDate() {
	EventDate = &eventDate{}
}
