package jrenders_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/jrenders"
	"github.com/YanshuoH/youkonger/models"
	"github.com/YanshuoH/youkonger/test"
	"testing"
	"time"
)

func TestJrenders(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jrenders Suite")
}

// write some data for sharing
var eventSet = []models.Event{
	models.Event{
		BaseModel: models.BaseModel{
			ID: 1,
		},
		Title:       "t",
		Description: "d",
		Location:    "l",
	},
}

var eventDateSet = []models.EventDate{
	models.EventDate{
		BaseModel: models.BaseModel{
			ID: 1,
		},
		Time:    time.Now(),
		EventID: 1,
		IsDDay:  true,
	},
}

var eventParticipantSet = []models.EventParticipant{
	models.EventParticipant{
		BaseModel: models.BaseModel{
			ID: 1,
		},
		EventDateID: 1,
	},
}

var participantUserSet = []models.ParticipantUser{
	models.ParticipantUser{
		BaseModel: models.BaseModel{
			ID: 1,
		},
		Name: "bigbro",
	},
}

var _ = BeforeSuite(func() {
	test.Setup()
	jrenders.Register()

	// insert dataset
	conn := dao.Conn
	for _, e := range eventSet {
		conn.Create(&e)
	}
	for _, ed := range eventDateSet {
		conn.Create(&ed)
	}
	for _, ep := range eventParticipantSet {
		conn.Create(&ep)
	}
	for _, pu := range participantUserSet {
		conn.Create(&pu)
	}
})

var _ = AfterSuite(func() {
	test.Teardown()
})
