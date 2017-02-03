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

func (e *event) FindByUUIDAndAdminHash(uuid, hash string) (*models.Event, error) {
	res := &models.Event{}
	err := e.Where("uuid = ? AND admin_hash = ? AND removed = FALSE", uuid, hash).First(res).Error
	return res, err
}

func (e *event) LoadEventDates(event *models.Event) error {
	return e.Where("event_id = ? AND removed = FALSE", event.ID).Find(&event.EventDates).Error
}

func (e *event) FindByEventParticipant(ep *models.EventParticipant) (*models.Event, error) {
	res := &models.Event{}
	err := e.Select("event.*").
		Joins("INNER JOIN event_date ed ON ed.event_id = event.id").
		Joins("INNER JOIN event_participant ep ON ep.event_date_id = ed.id").
		Where("ep.id = ?", ep.ID).
		First(res).
		Error

	return res, err
}

func (e *event) IsFinished(event *models.Event) (bool, error) {
	err := e.Where("event_id = ? AND removed = FALSE", event.ID).Find(&event.EventDates).Error
	if err != nil {
		return false, err
	}

	for _, ed := range event.EventDates {
		if ed.IsDDay {
			return true, nil
		}
	}

	return false, nil
}

var Event *event

func initEvent(conn *gorm.DB) {
	Event = &event{conn}
}
