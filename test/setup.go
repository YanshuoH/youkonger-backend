package test

import (
	"github.com/YanshuoH/youkonger/models"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/conf"
)

func Setup()  {
	// load conf
	c, err := conf.Setup("../conf/conf_test.toml")
	if err != nil {
		panic(err)
	}
	// connect mysql test db
	dao.Connect(c.DbConf.Dsn)
	if err != nil {
		panic(err)
	}

	// drop tables
	dao.Conn.
		DropTableIfExists(&models.Event{}).
		DropTableIfExists(&models.EventDate{}).
		DropTableIfExists(&models.EventParticipant{}).
		DropTableIfExists(&models.EventUnavailable{})

	// migration tables
	dao.AutoMigration()
	dao.Conn.LogMode(true)
}

func Teardown() {
	dao.Conn.Close()
}