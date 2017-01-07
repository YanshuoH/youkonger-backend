package dao

import (
	"github.com/YanshuoH/youkonger/models"
	"github.com/jinzhu/gorm"
)

type eventDate struct {
	*gorm.DB
}

func (ed *eventDate) FindByUUID(uuid string) (*models.EventDate, error) {
	res := &models.EventDate{}
	err := ed.Where("uuid = ? AND removed = FALSE", uuid).First(res).Error
	return res, err
}

func (ed *eventDate) LoadEventParticipants(eventDate *models.EventDate) error {
	return ed.Where("event_date_id = ? AND removed = FALSE", eventDate.ID).Find(&eventDate.EventParticipants).Error
}

var EventDate *eventDate

func initEventDate(conn *gorm.DB) {
	EventDate = &eventDate{conn}
}
