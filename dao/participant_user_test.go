package dao_test

import (
	. "github.com/YanshuoH/youkonger/dao"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jinzhu/gorm"
	"github.com/YanshuoH/youkonger/models"
)

var _ = Describe("ParticipantUser", func() {
	Describe("FindByUUID", func() {
		Context("With invalid uuid", func() {
			It("Should return an error", func() {
				_, err := ParticipantUser.FindByUUID("")
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With valid uuid", func() {
			var user models.ParticipantUser

			BeforeEach(func() {
				m := models.ParticipantUser{
					Name: "someone",
				}
				Expect(Conn.Create(&m).Error).ToNot(HaveOccurred())
				user = m
			})

			It("Should return the expected participant user", func() {
				res, err := ParticipantUser.FindByUUID(user.UUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.Name).To(Equal(user.Name))
			})
		})
	})
})
