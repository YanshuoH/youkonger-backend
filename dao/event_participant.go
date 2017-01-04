package dao

import "github.com/jinzhu/gorm"

type eventParticipant struct {
	*gorm.DB
}

var EventParticipant *eventParticipant

func initEventParticipant (conn *gorm.DB) {
	EventParticipant = &eventParticipant{conn}
}
