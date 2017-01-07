package dao

import (
	"github.com/YanshuoH/youkonger/models"
	"github.com/jinzhu/gorm"
)

type event struct {
	*gorm.DB
}

func (e *event) FindByUUID(uuid string) (*models.Event, error) {
	res := &models.Event{}
	err := e.Where("uuid = ? AND removed = FALSE", uuid).First(res).Error
	return res, err
}

func (e *event) LoadEventDates(event *models.Event) error {
	return e.Where("event_id = ? AND removed = FALSE", event.ID).Find(&event.EventDates).Error
}

var Event *event

func initEvent(conn *gorm.DB) {
	Event = &event{conn}
}
