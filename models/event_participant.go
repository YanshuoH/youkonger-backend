package models

type EventParticipant struct {
	BaseModel
	Name        string
	EventDateID uint64
}
