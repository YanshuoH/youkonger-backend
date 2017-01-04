package dao

import (
	"github.com/YanshuoH/youkonger/models"
	"github.com/go-playground/log"
	"github.com/jinzhu/gorm"
)

var Conn *gorm.DB

func Connect(dsnString string) *gorm.DB {
	log.Info("Dial mysql: " + dsnString)
	conn, err := gorm.Open("mysql", dsnString)
	if err != nil {
		panic(err)
	}

	Conn = conn
	Conn.SingularTable(true)

	return Conn
}

func AutoMigration() {
	Conn.
		AutoMigrate(&models.Event{}).
		AutoMigrate(&models.EventDate{}).
		AutoMigrate(&models.EventParticipant{}).
		AutoMigrate(&models.EventUnavailable{})
}