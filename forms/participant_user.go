package form

import (
	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	"github.com/YanshuoH/youkonger/utils"
)

type ParticipantUserForm struct {
	Name        string `json:"name" binding:"required"`
	EventUUID   string `json:"eventUuid" binding:"required"`
	UserUUID    string `json:"userUuid"`
	Remove      bool   `json:"remove"`
	Unavailable bool   `json:"unavailable"`

	EM              *dao.Manager            `json:"-"`
	Event           *models.Event           `json:"-"`
	User            *models.User            `json:"-"`
	ParticipantUser *models.ParticipantUser `json:"-"`
}

func (f *ParticipantUserForm) validate() *utils.CommonError {
	if f.EM == nil {
		return utils.NewCommonError(consts.NoEntityManagerInForm, nil)
	}

	e, err := f.EM.Event().FindByUUID(f.EventUUID)
	if err != nil {
		return utils.NewCommonError(consts.EventNotFound, err)
	}
	f.Event = e

	if f.UserUUID != "" {
		u, err := f.EM.User().FindByUUID(f.UserUUID)
		if err != nil {
			return utils.NewCommonError(consts.UserNotFound, err)
		}
		f.User = u

		pu, err := f.EM.ParticipantUser().FindByUserUUIDAndEventUUID(f.User.UUID, f.EventUUID)
		// it's ok to not found
		if err == nil {
			f.ParticipantUser = pu
		}
	}
	return nil
}

func (f *ParticipantUserForm) insert() (*models.ParticipantUser, *utils.CommonError) {
	if f.UserUUID == "" {
		// create user first
		u := &models.User{}
		if err := f.EM.Create(u).Error; err != nil {
			return nil, utils.NewCommonError(consts.FormSaveError, err)
		}
		f.User = u
	}

	m := &models.ParticipantUser{
		Name:        f.Name,
		EventID:     f.Event.ID,
		Event:       f.Event,
		UserID:      f.User.ID,
		User:        f.User,
		Unavailable: f.Unavailable,
	}
	if err := f.EM.Create(m).Error; err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}
	return m, nil
}

func (f *ParticipantUserForm) update() (*models.ParticipantUser, *utils.CommonError) {
	updateMap := map[string]interface{}{
		"unavailable": f.Unavailable,
		"name":        f.Name,
	}
	if err := f.EM.Model(f.ParticipantUser).Updates(updateMap).Error; err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}
	f.ParticipantUser.User = f.User
	f.ParticipantUser.Event = f.Event

	return f.ParticipantUser, nil
}

func (f *ParticipantUserForm) Handle() (*models.ParticipantUser, *utils.CommonError) {
	if cErr := f.validate(); cErr != nil {
		return nil, cErr
	}
	if f.ParticipantUser != nil {
		return f.update()
	}
	return f.insert()
}
