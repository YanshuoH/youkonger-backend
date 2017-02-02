package models

type EventParticipant struct {
	BaseModel
	EventDateID       uint64
	ParticipantUserID uint64
	ParticipantUser   *ParticipantUser
}
