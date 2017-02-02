package models

// ParticipantUser stores the user who's participated an event
type ParticipantUser struct {
	BaseModel
	Name        string
	EventID     uint64
	Event       *Event
	UserID      uint64
	User        *User
	Unavailable bool `gorm:"default:FALSE"`
}
