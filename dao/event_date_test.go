package dao_test

import (
	. "github.com/YanshuoH/youkonger/dao"

	"github.com/YanshuoH/youkonger/models"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("EventDate", func() {

	Describe("FindByUUID", func() {
		Context("With not existed uuid", func() {
			It("Should return an error with record not found", func() {
				_, err := EventDate.FindByUUID("")
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With existed uuid", func() {
			It("Should return the right eventDate entity", func() {
				toInsert :=
					models.EventDate{
						Time: time.Now(),
					}
				Expect(Conn.Create(&toInsert).Error).ToNot(HaveOccurred())

				ed, err := EventDate.FindByUUID(toInsert.UUID)
				Expect(err).To(BeNil())
				// format equal
				Expect(ed.Time.Format(time.RFC822)).To(Equal(toInsert.Time.Format(time.RFC822)))
			})
		})
	})
})
