package dao_test

import (
	. "github.com/YanshuoH/youkonger/dao"

	"github.com/YanshuoH/youkonger/models"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventParticipant", func() {
	Describe("FindByUUID", func() {
		Context("With not existed uuid", func() {
			It("Should return an error with record not found", func() {
				_, err := EventParticipant.FindByUUID("")
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With existed uuid", func() {
			It("Should return the right eventDate entity", func() {
				toInsert := models.EventParticipant{
					EventDateID: 51,
				}
				Expect(Conn.Create(&toInsert).Error).ToNot(HaveOccurred())

				ep, err := EventParticipant.FindByUUID(toInsert.UUID)
				Expect(err).To(BeNil())
				Expect(ep.EventDateID).To(Equal(toInsert.EventDateID))
			})
		})
	})
})
