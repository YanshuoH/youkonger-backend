package models

type EventParticipant struct {
	BaseModel
	EventDateID       uint64
	ParticipantUserId int64
	ParticipantUser   *ParticipantUser
}
