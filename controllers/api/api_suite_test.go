package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/YanshuoH/youkonger/test"
	"github.com/YanshuoH/youkonger/jrenders"
	"github.com/YanshuoH/youkonger/models"
	"time"
	"github.com/YanshuoH/youkonger/dao"
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Suite")
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
		Time: time.Now(),
		EventID: 1,
	},
}

var participantUserSet = []models.ParticipantUser{
	models.ParticipantUser{
		BaseModel: models.BaseModel{
			ID: 1,
		},
		Name: "someone",
	},
}

var eventParticipantSet = []models.EventParticipant{
	models.EventParticipant{
		BaseModel: models.BaseModel{
			ID: 1,
		},
		EventDateID: 1,
		ParticipantUserID: 1,
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
	for _, pu := range participantUserSet {
		conn.Create(&pu)
	}
	for _, ep := range eventParticipantSet {
		conn.Create(&ep)
	}
})

var _ = AfterSuite(func() {
	test.Teardown()
})
