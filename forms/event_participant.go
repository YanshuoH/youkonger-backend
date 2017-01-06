package form

import (
	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	"github.com/YanshuoH/youkonger/utils"
)

type EventParticipantForm struct {
	UUID          string `json:"uuid"`
	EventDateUUID string `json:"eventDateUuid" binding:"required"`
	Name          string `json:"name" binding:"required"`

	// transients
	EM               *dao.Manager             `json:"-"`
	EventDate        *models.EventDate        `json:"-"`
	EventParticipant *models.EventParticipant `json:"-"`
}

type EventParticipantForms struct {
	Forms []EventParticipantForm `json:"eventParticipantFormList"`
}

func (f *EventParticipantForm) validate() *utils.CommonError {
	if f.EM == nil {
		return utils.NewCommonError(consts.NoEntityManagerInForm, nil)
	}

	if f.UUID != "" {
		ep, err := f.EM.EventParticipant().FindByUUID(f.UUID)
		if err != nil {
			return utils.NewCommonError(consts.EventParticipantNotFound, nil)
		}
		f.EventParticipant = ep
	}

	ed, err := f.EM.EventDate().FindByUUID(f.EventDateUUID)
	if err != nil {
		return utils.NewCommonError(consts.EventDateNotFound, nil)
	}
	f.EventDate = ed

	return nil
}

func (f *EventParticipantForm) insert() (*models.EventParticipant, *utils.CommonError) {
	ep := &models.EventParticipant{
		Name: f.Name,
		EventDateID: f.EventDate.ID,
	}
	if err := f.EM.Create(ep).Error; err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}

	return ep, nil
}

func (f *EventParticipantForm) update() (*models.EventParticipant, *utils.CommonError) {
	err := f.EM.Model(f.EventParticipant).Where("uuid = ?", f.UUID).Update("name", f.Name).Error
	if err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}
	return f.EventParticipant, nil
}

func (f *EventParticipantForm) Handle() (*models.EventParticipant, *utils.CommonError) {
	if cErr := f.validate(); cErr != nil {
		return nil, cErr
	}

	if f.UUID != "" && !f.EM.NewRecord(f.EventParticipant) {
		return f.update()
	}

	return f.insert()
}