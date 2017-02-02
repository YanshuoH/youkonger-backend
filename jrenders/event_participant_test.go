package jrenders_test

import (
	. "github.com/YanshuoH/youkonger/jrenders"

	"github.com/YanshuoH/youkonger/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("EventParticipant", func() {
	var participantUser models.ParticipantUser
	var eventParticipants []models.EventParticipant
	var eventDate models.EventDate
	BeforeEach(func() {
		eps := []models.EventParticipant{
			models.EventParticipant{
				ParticipantUserID: participantUserSet[0].ID,
			},
			models.EventParticipant{
				ParticipantUserID: participantUserSet[0].ID,
			},
		}
		ed := models.EventDate{
			Time: time.Now(),
			BaseModel: models.BaseModel{
				UUID: "123-345",
			},
		}

		participantUser = participantUserSet[0]
		eventParticipants = eps
		eventDate = ed
	})

	Describe("Itemize", func() {
		Context("With eventParticipant and eventDate provided", func() {
			It("Should return the struct for json rendering", func() {
				ep := eventParticipants[0]
				j := EventParticipant.Itemize(&ep, &eventDate)
				Expect(j.EventDateUUID).To(Equal(eventDate.UUID))
				Expect(j.Name).To(Equal(participantUser.Name))
			})
		})
	})

	Describe("List", func() {
		Context("With eventParticipants and eventDate provided", func() {
			It("Should return the struct for json rendering", func() {
				j := EventParticipant.List(eventParticipants, &eventDate)
				Expect(j.JList).To(HaveLen(len(eventParticipants)))
				Expect(j.JList[0].Name).To(Equal(participantUser.Name))
			})
		})
	})
})
