package dao_test

import (
	. "github.com/YanshuoH/youkonger/dao"

	"github.com/YanshuoH/youkonger/models"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
)

var _ = Describe("Event", func() {

	Describe("FindByUUID", func() {
		Context("With not existed uuid", func() {
			It("Should return an error with record not found", func() {
				_, err := Event.FindByUUID("")
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With existed uuid", func() {
			It("Should return the right event entity", func() {
				toInsert :=
					models.Event{
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
})
