package dao_test

import (
	. "github.com/YanshuoH/youkonger/dao"

	"github.com/YanshuoH/youkonger/models"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	"time"
)

var _ = Describe("Event", func() {
	var event models.Event
	var eventDates []models.EventDate
	BeforeEach(func() {
		eToInsert := models.Event{
			Title:       "t",
			Description: "d",
			Location:    "l",
		}
		Expect(Conn.Create(&eToInsert).Error).ToNot(HaveOccurred())
		edToInserts := []models.EventDate{
			models.EventDate{
				Time: time.Now(),
			},
			models.EventDate{
				Time: time.Now(),
			},
		}
		for _, ed := range edToInserts {
			ed.EventID = eToInsert.ID
			Expect(Conn.Create(&ed).Error).ToNot(HaveOccurred())
		}
		event = eToInsert
		eventDates = edToInserts
	})

	Describe("FindByUUID", func() {
		Context("With not existed uuid", func() {
			It("Should return an error with record not found", func() {
				_, err := Event.FindByUUID("")
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With existed uuid", func() {
			It("Should return the right event entity", func() {
				toInsert := models.Event{
					Title:       "title",
					Description: "description",
					Location:    "beijing",
					AdminHash:   uuid.NewV4().String(),
				}

				Expect(Conn.Create(&toInsert).Error).ToNot(HaveOccurred())

				e, err := Event.FindByUUID(toInsert.UUID)
				Expect(err).To(BeNil())
				Expect(e.Description).To(Equal("description"))
			})
		})
	})

	Describe("FindByUUIDAndAdminHash", func() {
		var event models.Event
		BeforeEach(func() {
			e := models.Event{
				Title: "thing",
			}
			Expect(Conn.Create(&e).Error).ToNot(HaveOccurred())
			event = e
		})
		Context("With none existed uuid and hash", func() {
			It("Should return an error", func() {
				_, err := Event.FindByUUIDAndAdminHash(event.UUID, "something")
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With correct uuid and hash", func() {
			It("Should return the expected event", func() {
				res, err := Event.FindByUUIDAndAdminHash(event.UUID, event.AdminHash)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.ID).To(Equal(event.ID))
			})
		})
	})

	Describe("LoadEventDate", func() {
		Context("With given event", func() {
			It("Should return the related eventDates", func() {
				err := Event.LoadEventDates(&event)
				Expect(err).ToNot(HaveOccurred())
				Expect(event.EventDates).To(HaveLen(len(eventDates)))
			})
		})
	})

	Describe("FindByEventParticipant", func() {
		Context("With unreachable participant id", func() {
			It("Should return an not found error", func() {
				_, err := Event.FindByEventParticipant(&models.EventParticipant{})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With right participant id", func() {
			It("Should return the right event", func() {
				e := models.Event{
					Title: "thing",
				}
				Expect(Conn.Create(&e).Error).NotTo(HaveOccurred())
				ed := models.EventDate{
					Time:    time.Now(),
					EventID: e.ID,
				}
				Expect(Conn.Create(&ed).Error).NotTo(HaveOccurred())
				ep := models.EventParticipant{
					EventDateID: ed.ID,
				}
				Expect(Conn.Create(&ep).Error).NotTo(HaveOccurred())

				res, err := Event.FindByEventParticipant(&ep)
				Expect(err).NotTo(HaveOccurred())
				Expect(res.ID).To(Equal(e.ID))
			})
		})
	})

	Describe("IsFinished", func() {
		Context("Without dday", func() {
			It("Should return false", func() {
				finished, err := Event.IsFinished(&event)
				Expect(err).ToNot(HaveOccurred())
				Expect(finished).To(BeFalse())
			})
		})

		Context("With dday", func() {
			It("Should return true", func() {
				// insert a dday event/eventDate set
				e := models.Event{
					Title: "thing",
				}
				Expect(Conn.Create(&e).Error).NotTo(HaveOccurred())
				ed := models.EventDate{
					Time:    time.Now(),
					EventID: e.ID,
					IsDDay:  true,
				}
				Expect(Conn.Create(&ed).Error).NotTo(HaveOccurred())

				finished, err := Event.IsFinished(&e)
				Expect(err).ToNot(HaveOccurred())
				Expect(finished).To(BeTrue())
			})
		})
	})
})
