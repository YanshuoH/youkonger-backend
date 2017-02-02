package models

import (
	"time"
)

type EventDate struct {
	BaseModel
	EventID uint64
	Time    time.Time
	IsDDay  bool `gorm:"default:FALSE"`

	EventParticipants []EventParticipant
}
