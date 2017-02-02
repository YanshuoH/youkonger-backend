package dao_test

import (
	. "github.com/YanshuoH/youkonger/dao"

	"github.com/YanshuoH/youkonger/models"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParticipantUser", func() {
	var user models.User
	var event models.Event
	var participantUser models.ParticipantUser

	BeforeEach(func() {
		u := models.User{}
		Expect(Conn.Create(&u).Error).ToNot(HaveOccurred())
		user = u

		e := models.Event{
			Title: "a title",
		}
		Expect(Conn.Create(&e).Error).ToNot(HaveOccurred())
		event = e

		pu := models.ParticipantUser{
			Name:    "someone",
			UserID:  u.ID,
			EventID: e.ID,
		}
		Expect(Conn.Create(&pu).Error).ToNot(HaveOccurred())
		participantUser = pu

	})

	Describe("FindById", func() {
		Context("With invalid id", func() {
			It("Should return an error", func() {
				_, err := ParticipantUser.FindById(666)
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With valid id", func() {
			It("Should return the expected participant user", func() {
				res, err := ParticipantUser.FindByUUID(participantUser.UUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.Name).To(Equal(participantUser.Name))
				Expect(res.ID).To(Equal(participantUser.ID))
			})
		})
	})

	Describe("FindByUUID", func() {
		Context("With invalid uuid", func() {
			It("Should return an error", func() {
				_, err := ParticipantUser.FindByUUID("")
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With valid uuid", func() {
			It("Should return the expected participant user", func() {
				res, err := ParticipantUser.FindByUUID(participantUser.UUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.Name).To(Equal(participantUser.Name))
			})
		})
	})

	Describe("FindByUUIDAndEventUUID", func() {
		Context("With invalid uuid or eventUuid", func() {
			It("Should return an error", func() {
				_, err := ParticipantUser.FindByUserUUIDAndEventUUID("", "")
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With valid uuid and eventUuid", func() {
			It("Should return the expected participant user", func() {
				res, err := ParticipantUser.FindByUserUUIDAndEventUUID(user.UUID, event.UUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.Name).To(Equal(participantUser.Name))
				Expect(res.ID).To(Equal(participantUser.ID))
			})
		})
	})

	Describe("FindUnavailableByEventID", func() {
		Context("With unexisted event id", func() {
			It("Should return an empty slice", func() {
				res, err := ParticipantUser.FindUnavailableByEventID(666)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(HaveLen(0))
			})
		})

		Context("With existed event id", func() {
			It("Should return all unavailable participant users", func() {
				// insert some
				unavailablePU := models.ParticipantUser{
					Name: "a",
					EventID: event.ID,
					UserID: 999,
					Unavailable: true,
				}
				Expect(Conn.Create(&unavailablePU).Error).ToNot(HaveOccurred())
				res, err := ParticipantUser.FindUnavailableByEventID(event.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(HaveLen(1))
			})
		})
	})
})
