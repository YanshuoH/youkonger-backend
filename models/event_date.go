package models

import "time"

type EventDate struct {
	BaseModel
	EventID uint64
	Time    time.Time
}
