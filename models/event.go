package models

import (
	"github.com/satori/go.uuid"
	"time"
)

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
	e.AdminHash = uuid.NewV4().String()
	e.UUID = uuid.NewV4().String()

	now := time.Now()
	e.CreatedAt = now
	e.UpdatedAt = now
}