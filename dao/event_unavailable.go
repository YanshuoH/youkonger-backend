package dao

import "github.com/jinzhu/gorm"

type eventUnavailable struct {
	*gorm.DB
}

var EventUnavailable *eventUnavailable

func initEventUnavailable(conn *gorm.DB) {
	EventUnavailable = &eventUnavailable{conn}
}
