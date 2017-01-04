package dao

import "github.com/jinzhu/gorm"

type event struct {
	*gorm.DB
}

var Event *event

func initEvent(conn *gorm.DB) {
	Event = &event{conn}
}
