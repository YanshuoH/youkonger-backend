package dao_test

import (
	. "github.com/YanshuoH/youkonger/dao"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/YanshuoH/youkonger/models"
	"github.com/jinzhu/gorm"
)

var _ = Describe("User", func() {
	var user models.User

	BeforeEach(func() {
		u := models.User{}
		Expect(Conn.Create(&u).Error).ToNot(HaveOccurred())
		user = u
	})

	Describe("FindById", func() {
		Context("With wrong id", func() {
			It("Should return an error", func() {
				_, err := User.FindById(666)
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With right id", func() {
			It("Should return the expected user", func() {
				res, err := User.FindById(user.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.UUID).To(Equal(user.UUID))
			})
		})
	})

	Describe("FindByUUID", func() {
		Context("With wrong uuid", func() {
			It("Should return an error", func() {
				_, err := User.FindByUUID("")
				Expect(err.Error()).To(Equal(gorm.ErrRecordNotFound.Error()))
			})
		})

		Context("With right uuid", func() {
			It("Should return the expected user", func() {
				res, err := User.FindByUUID(user.UUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(res.ID).To(Equal(user.ID))
			})
		})
	})
})
