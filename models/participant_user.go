package models

// ParticipantUser stores the user who's participated an event
type ParticipantUser struct {
	BaseModel
	Name    string
	EventId uint64
	Event   *Event
	UserId  uint64
	User    *User
}
