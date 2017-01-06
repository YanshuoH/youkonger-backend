package consts

const (
	DefaultErrorMsg                   = "DefaultErrorMsg"
	FormSaveError                     = "FormSaveError"
	NoEntityManagerInForm             = "NoEntityManagerInForm"
	IncorrectUnixTime                 = "IncorrectUnixTime"
	EventNotFound                     = "EventNotFound"
	EventDateNotFound                 = "EventDateNotFound"
	EventDateRequiredWhenInitializing = "EventDateRequiredWhenInitializing"
	EventParticipantNotFound          = "EventParticipantNotFound"
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
		DefaultErrorMsg:                   "发生了未知错误",
		NoEntityManagerInForm:             "提交过程中产生了内部错误, 请稍后再试",
		IncorrectUnixTime:                 "发送的时间有误",
		EventNotFound:                     "未找到指定聚会信息",
		EventDateNotFound:                 "未找到指定的聚会日期",
		EventDateRequiredWhenInitializing: "第一次创建时必须要填至少一个聚会日期",
		EventParticipantNotFound:          "未找到指定的参与人员",
		FormSaveError:                     "保存时发生了意外错误, 请稍候重试",
	},
}
