package form

import (
	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	"github.com/YanshuoH/youkonger/utils"
	"github.com/pkg/errors"
)

type EventParticipantForm struct {
	UUID          string `json:"uuid"`
	EventDateUUID string `json:"eventDateUuid" binding:"required"`
	Remove        bool   `json:"remove"`

	// transients
	EM               *dao.Manager             `json:"-"`
	EventDate        *models.EventDate        `json:"-"`
	EventParticipant *models.EventParticipant `json:"-"`
	ParticipantUser  *models.ParticipantUser  `json:"-"`
}

type EventParticipantForms struct {
	Forms []EventParticipantForm `json:"eventParticipantList"`
	Name  string                 `json:"name" binding:"required"`

	// user's uuid
	UUID      string `json:"participantUserUuid"`
	EventUUID string `json:"eventUuid" binding:"required"`

	// internal
	EM              *dao.Manager            `json:"-"`
	ParticipantUser *models.ParticipantUser `json:"-"`
	Event *models.Event `json:"-"`
}

func (f *EventParticipantForm) validate() *utils.CommonError {
	if f.EM == nil {
		return utils.NewCommonError(consts.NoEntityManagerInForm, nil)
	}
	if f.ParticipantUser == nil {
		return utils.NewCommonError(consts.NoParticipantUserInForm, nil)
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
		ParticipantUserId: f.ParticipantUser.ID,
		EventDateID:       f.EventDate.ID,
	}
	if err := f.EM.Create(ep).Error; err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}

	return ep, nil
}

func (f *EventParticipantForm) update() (*models.EventParticipant, *utils.CommonError) {
	err := f.EM.Model(f.EventParticipant).Where("uuid = ?", f.UUID).Update("removed", f.Remove).Error

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

func (f *EventParticipantForms) validate() *utils.CommonError {
	if f.EM == nil {
		return utils.NewCommonError(consts.NoEntityManagerInForm, nil)
	}

	e, err := f.EM.Event().FindByUUID(f.EventUUID)
	if err != nil {
		return utils.NewCommonError(consts.EventNotFound, err)
	}
	f.Event = e

	if f.UUID != "" {
		pu, err := f.EM.ParticipantUser().FindByUUIDAndEventUUID(f.UUID, f.Event.UUID)
		if err != nil {
			return utils.NewCommonError(consts.ParticipantUserNotFound, err)
		}
		f.ParticipantUser = pu
	}

	return nil
}

func (f *EventParticipantForms) insert() (*models.ParticipantUser, *utils.CommonError) {
	m := &models.ParticipantUser{
		Name: f.Name,
		EventId: f.Event.ID,
	}
	if err := f.EM.Create(m).Error; err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}

	return m, nil
}

func (f *EventParticipantForms) update() (*models.ParticipantUser, *utils.CommonError) {
	if err := f.EM.Model(f.ParticipantUser).Update("name", f.Name).Error; err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}

	return f.ParticipantUser, nil
}

func (f *EventParticipantForms) Handle() (res []models.EventParticipant, cErr *utils.CommonError) {
	// not allow empty slice
	if len(f.Forms) == 0 {
		return res, utils.NewCommonError(consts.FormInvalid, errors.New("Expected a none-zero length form"))
	}
	invalidCount := 0
	for _, epf := range f.Forms {
		if epf.UUID == "" && epf.Remove {
			invalidCount++
		}
	}
	if invalidCount == len(f.Forms) {
		return res, utils.NewCommonError(consts.FormInvalid, errors.New("Expected a none-zero length participant form"))
	}

	// create participant user
	if cErr := f.validate(); cErr != nil {
		return res, cErr
	}
	if f.UUID != "" && !f.EM.NewRecord(f.ParticipantUser) {
		if _, cErr := f.update(); cErr != nil {
			return res, cErr
		}
	} else {
		pu, cErr := f.insert()
		if cErr != nil {
			return res, cErr
		}
		f.ParticipantUser = pu
	}

	for _, epf := range f.Forms {
		// nothing to do if is insert but set to removed
		if epf.UUID == "" && epf.Remove {
			continue
		}

		var ep *models.EventParticipant
		epf.EM = f.EM
		epf.ParticipantUser = f.ParticipantUser
		if ep, cErr = epf.Handle(); cErr != nil {
			return
		}
		// inject participant user
		ep.ParticipantUser = f.ParticipantUser

		res = append(res, *ep)
	}
	return
}
