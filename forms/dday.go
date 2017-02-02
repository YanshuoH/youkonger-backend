package form

import (
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	"github.com/YanshuoH/youkonger/utils"
	"github.com/YanshuoH/youkonger/consts"
)

type DDayForm struct {
	UUID          string `json:"uuid" binding:"required"`
	Hash          string `json:"hash" binding:"required"`
	EventDateUUID string `json:"eventDateUuid" binding:"required"`

	EM    *dao.Manager  `json:"-"`
	Event *models.Event `json:"-"`
	EventDate *models.EventDate `json:"-"`
}

func (f *DDayForm) validate() *utils.CommonError {
	if f.EM == nil {
		return utils.NewCommonError(consts.NoEntityManagerInForm, nil)
	}
	e, err := f.EM.Event().FindByUUIDAndAdminHash(f.UUID, f.Hash)
	if err != nil {
		return utils.NewCommonError(consts.EventNotFound, err)
	}
	f.Event = e

	ed, err := f.EM.EventDate().FindByUUID(f.EventDateUUID)
	if err != nil {
		return utils.NewCommonError(consts.EventDateNotFound, err)
	}
	f.EventDate = ed

	if f.EventDate.EventID != f.Event.ID {
		return utils.NewCommonError(consts.EventDateNotFound, nil)
	}

	return nil
}

func (f *DDayForm) update() (*models.Event, *utils.CommonError) {
	// find all event dates under that event
	// set isDDay => false
	if err := f.EM.
		Model(&models.EventDate{}).
		Where("event_id = ? AND removed = FALSE", f.Event.ID).
		Update("is_d_day", false).Error; err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}
	// set true to the specified date
	if err := f.EM.Model(f.EventDate).Update("is_d_day", true).Error; err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}

	return f.Event, nil
}

func (f *DDayForm) Handle() (*models.Event, *utils.CommonError) {
	if cErr := f.validate(); cErr != nil {
		return nil, cErr
	}

	return f.update()
}
