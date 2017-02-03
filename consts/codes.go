package consts

import "fmt"

const (
	OK                                = "OK"
	DefaultErrorMsg                   = "DefaultErrorMsg"
	FormSaveError                     = "FormSaveError"
	FormInvalid                       = "FormInvalid"
	NoEntityManagerInForm             = "NoEntityManagerInForm"
	NoParticipantUserInForm           = "NoParticipantUserInForm"
	IncorrectUnixTime                 = "IncorrectUnixTime"
	EventNotFound                     = "EventNotFound"
	EventDateNotFound                 = "EventDateNotFound"
	EventDateRequiredWhenInitializing = "EventDateRequiredWhenInitializing"
	EventParticipantNotFound          = "EventParticipantNotFound"
	UserNotFound                      = "UserNotFound"
	ParticipantUserNotFound           = "ParticipantUserNotFound"
	InvalidAdminHash                  = "InvalidAdminHash"
	TileTooLong                       = "TileTooLong"
	DescriptionTooLong                = "DescriptionTooLong"
	LocationTooLong                   = "LocationTooLong"
	NameTooLong                       = "NameTooLong"
)

type messenger struct {
	messages map[string]string
}

func (m messenger) Get(code string) string {
	if val, ok := m.messages[code]; ok {
		return val
	}

	return m.messages[DefaultErrorMsg]
}

var Messenger = &messenger{
	messages: map[string]string{
		OK:                                "OK",
		DefaultErrorMsg:                   "发生了未知错误",
		NoEntityManagerInForm:             "提交过程中产生了内部错误, 请稍后再试",
		NoParticipantUserInForm:           "提交参与者的过程中产生了内部错误, 请稍后再试",
		IncorrectUnixTime:                 "发送的时间有误",
		EventNotFound:                     "未找到指定聚会信息",
		EventDateNotFound:                 "未找到指定的聚会日期",
		EventDateRequiredWhenInitializing: "第一次创建时必须要填至少一个聚会日期",
		EventParticipantNotFound:          "未找到指定的参与人员",
		UserNotFound:                      "未找到指定用户",
		ParticipantUserNotFound:           "未找到指定的参与人员信息",
		InvalidAdminHash:                  "管理员序列有误, 请检查",
		FormSaveError:                     "保存时发生了意外错误, 请稍候重试",
		FormInvalid:                       "表单填写有误, 请检查",
		TileTooLong:                       fmt.Sprintf("标题过长, 请保持在%d字之内", TitleLengthConstraint),
		DescriptionTooLong:                fmt.Sprintf("附加信息过长, 请保持在%d字之内", DescriptionLengthConstraint),
		LocationTooLong:                   fmt.Sprintf("地址过长, 请保持在%d字之内", LocationLengthConstraint),
		NameTooLong:                       fmt.Sprintf("名字过长, 请保持在%d字之内", NameLengthConstraint),
	},
}
