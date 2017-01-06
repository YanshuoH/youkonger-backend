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

var Event *event

func initEvent(conn *gorm.DB) {
	Event = &event{conn}
}
