package form

import (
	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	"github.com/YanshuoH/youkonger/utils"
	"time"
)

type EventDateForm struct {
	UUID       string `json:"uuid"`
	EventUUID  string `json:"eventUuid" binding:"required"`
	TimeInUnix int64  `json:"timeInUnix" binding:"required"`

	// transients
	Event     *models.Event     `json:"-"`
	EventDate *models.EventDate `json:"-"`
	EM        *dao.Manager      `json:"-"`
}

type EventDateForms struct {
	Forms []*EventDateForm `json:"eventDateList"`
	Event *models.Event   `json:"-"`
	EM    *dao.Manager    `json:"-"`
}

func (f *EventDateForm) validate() *utils.CommonError {
	if f.EM == nil {
		return utils.NewCommonError(consts.NoEntityManagerInForm, nil)
	}

	if f.TimeInUnix < 0 {
		return utils.NewCommonError(consts.IncorrectUnixTime, nil)
	}

	// try to load EventDate if uuid provided
	if f.UUID != "" {
		ed, err := f.EM.EventDate().FindByUUID(f.UUID)
		if err != nil {
			return utils.NewCommonError(consts.EventDateNotFound, nil)
		}
		f.EventDate = ed
	}

	// try to load event if not provided in form
	if f.Event == nil {
		e, err := f.EM.Event().FindByUUID(f.EventUUID)
		if err != nil {
			return utils.NewCommonError(consts.EventNotFound, nil)
		}
		f.Event = e
	}

	return nil
}

func (f *EventDateForm) insert() (*models.EventDate, *utils.CommonError) {
	ed := models.EventDate{
		EventID: f.Event.ID,
		Time:    time.Unix(f.TimeInUnix, 0),
	}
	if err := f.EM.Create(&ed).Error; err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}

	return &ed, nil
}

func (f *EventDateForm) update() (*models.EventDate, *utils.CommonError) {
	err := f.EM.Model(f.EventDate).Where("uuid = ?", f.UUID).Update("time", f.TimeInUnix).Error
	if err != nil {
		return nil, utils.NewCommonError(consts.FormSaveError, err)
	}
	return f.EventDate, nil
}

func (f *EventDateForm) Handle() (*models.EventDate, *utils.CommonError) {
	if cErr := f.validate(); cErr != nil {
		return nil, cErr
	}

	if f.UUID != "" && !f.EM.NewRecord(f.EventDate) {
		return f.update()
	}

	return f.insert()
}

func (f *EventDateForms) Handle() (res []models.EventDate, cErr *utils.CommonError) {
	// must have event entity in form
	if f.Event == nil || f.EM == nil {
		return res, utils.NewCommonError(consts.DefaultErrorMsg, nil, "Coding error! Check EventDateForms")
	}
	for _, edf := range f.Forms {
		// inject em and event
		edf.EM = f.EM
		edf.Event = f.Event

		if cErr = edf.validate(); cErr != nil {
			return res, cErr
		}
	}

	for _, edf := range f.Forms {
		ed, cErr := edf.Handle()
		if cErr != nil {
			return res, cErr
		}
		res = append(res, *ed)
	}

	return res, nil
}
