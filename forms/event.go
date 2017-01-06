package form

import (
	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	"github.com/YanshuoH/youkonger/utils"
)

type EventForm struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"title"`
	Location    string `json:"location"`

	EventDateForms

	// transients
	Event *models.Event
	EM    *dao.Manager
}

func (f *EventForm) validate() *utils.CommonError {
	if f.EM == nil {
		return utils.NewCommonError(consts.NoEntityManagerInForm, nil)
	}

	if f.UUID != "" {
		ed, err := f.EM.Event().FindByUUID(f.UUID)
		if err != nil {
			return utils.NewCommonError(consts.EventNotFound, nil)
		}
		f.Event = ed
	}

	return nil
}

func (f *EventForm) update() (*models.Event, *utils.CommonError) {
	m := map[string]interface{}{
		"title":       f.Title,
		"description": f.Description,
		"location":    f.Location,
	}
	err := f.EM.Model(f.Event).Where("uuid = ?", f.UUID).Updates(m).Error
	if err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}

	return f.Event, nil
}

func (f *EventForm) insert() (*models.Event, *utils.CommonError) {
	e := &models.Event{
		Title:       f.Title,
		Description: f.Description,
		Location:    f.Location,
	}

	if err := f.EM.Create(e).Error; err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}

	return e, nil
}

func (f *EventForm) Handle() (*models.Event, *utils.CommonError) {
	if cErr := f.validate(); cErr != nil {
		return nil, cErr
	}

	var e *models.Event
	var cErr *utils.CommonError
	isCreate := false

	if f.UUID != "" && !f.EM.NewRecord(f.Event) {
		e, cErr = f.update()
	} else {
		e, cErr = f.insert()
		isCreate = true
	}
	if cErr != nil {
		return nil, cErr
	}

	// the first insertion must have more than one dates
	if isCreate && len(f.EventDateForms.Forms) == 0 {
		return nil, utils.NewCommonError(consts.EventDateRequiredWhenInitializing, nil)
	}

	if len(f.EventDateForms.Forms) > 0 {
		// try to insert event dates
		f.EventDateForms.EM = f.EM
		f.EventDateForms.Event = e
		eds, cErr := f.EventDateForms.Handle()
		if cErr != nil {
			return nil, cErr
		}

		e.EventDates = eds
	}

	return e, cErr
}
