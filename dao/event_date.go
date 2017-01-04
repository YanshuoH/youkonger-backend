package dao

import "github.com/jinzhu/gorm"

type eventDate struct {
	*gorm.DB
}

var EventDate *eventDate

func initEventDate(conn *gorm.DB) {
	EventDate = &eventDate{conn}
}
