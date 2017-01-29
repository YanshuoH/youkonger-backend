package models

type EventParticipant struct {
	BaseModel
	EventDateID       uint64
	ParticipantUserId uint64
	ParticipantUser   *ParticipantUser
}
