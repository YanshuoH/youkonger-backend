package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/YanshuoH/youkonger/models"
)

type eventParticipant struct {
	*gorm.DB
}

func (ep *eventParticipant) FindByUUID(uuid string) (*models.EventParticipant, error) {
	res := &models.EventParticipant{}
	err := ep.Where("uuid = ? AND removed = FALSE", uuid).First(res).Error
	return res, err
}

var EventParticipant *eventParticipant

func initEventParticipant (conn *gorm.DB) {
	EventParticipant = &eventParticipant{conn}
}
