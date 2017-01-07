package jrenders_test

import (
	. "github.com/YanshuoH/youkonger/jrenders"

	"github.com/YanshuoH/youkonger/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("EventParticipant", func() {
	var eventParticipants = []models.EventParticipant{
		models.EventParticipant{
			Name: "bigbro",
		},
		models.EventParticipant{
			Name: "littlebro",
		},
	}
	var eventDate = models.EventDate{
		Time: time.Now(),
		BaseModel: models.BaseModel{
			UUID: "123-345",
		},
	}

	Describe("Itemize", func() {
		Context("With eventParticipant and eventDate provided", func() {
			It("Should return the struct for json rendering", func() {
				ep := eventParticipants[0]
				j := EventParticipant.Itemize(&ep, &eventDate)
				Expect(j.EventDateUUID).To(Equal(eventDate.UUID))
				Expect(j.Name).To(Equal(ep.Name))
			})
		})
	})

	Describe("List", func() {
		Context("With eventParticipants and eventDate provided", func() {
			It("Should return the struct for json rendering", func() {
				j := EventParticipant.List(eventParticipants, &eventDate)
				Expect(j.JList).To(HaveLen(len(eventParticipants)))
				Expect(j.JList[0].Name).To(Equal(eventParticipants[0].Name))
			})
		})
	})
})
