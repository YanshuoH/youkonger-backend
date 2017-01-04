package models

import "github.com/satori/go.uuid"

type Event struct {
	BaseModel
	Title            string
	Description      string
	Location         string
	AdminHash        string `gorm:"unique_index"`
	EventDates       []EventDate
	EventParticipant []EventParticipant
	EventUnavailable []EventUnavailable
}

func (e *Event) BeforeCreate() {
	e.UUID = uuid.NewV4().String()
}
