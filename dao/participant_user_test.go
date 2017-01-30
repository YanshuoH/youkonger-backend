package dao_test

import (
	. "github.com/YanshuoH/youkonger/dao"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jinzhu/gorm"
	"github.com/YanshuoH/youkonger/models"
)

var _ = Describe("ParticipantUser", func() {
	var user models.ParticipantUser
	var event models.Event

	BeforeEach(func() {
		e := models.Event{
			Title: "a title",
		}
		Expect(Conn.Create(&e).Error).ToNot(HaveOccurred())
		event = e

		pu := models.ParticipantUser{
			Name: "someone",
			EventId: e.ID,
		}
		Expect(Conn.Create(&pu).Error).ToNot(HaveOccurred())
		user = pu

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
				res, err := ParticipantUser.FindByUUID(user.UUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.Name).To(Equal(user.Name))
				Expect(res.ID).To(Equal(user.ID))
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
				res, err := ParticipantUser.FindByUUID(user.UUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.Name).To(Equal(user.Name))
			})
		})
	})

	Describe("FindByUUIDAndEventUUID", func() {
		Context("With invalid uuid or eventUuid", func() {
			It("Should return an error", func() {
				_, err := ParticipantUser.FindByUUIDAndEventUUID("", "")
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With valid uuid and eventUuid", func() {
			It("Should return the expected participant user", func() {
				res, err := ParticipantUser.FindByUUIDAndEventUUID(user.UUID, event.UUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.Name).To(Equal(user.Name))
				Expect(res.ID).To(Equal(user.ID))
			})
		})
	})
})
