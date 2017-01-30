package jrenders_test

import (
	. "github.com/YanshuoH/youkonger/jrenders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/YanshuoH/youkonger/models"
	"github.com/YanshuoH/youkonger/dao"
)

var _ = Describe("ParticipantUser", func() {
	var user models.User
	var participantUser models.ParticipantUser

	BeforeEach(func() {
		u := models.User{}
		Expect(dao.GetManager().Create(&u).Error).ToNot(HaveOccurred())
		user = u

		pu := models.ParticipantUser{
			Name: "someone",
			UserId: user.ID,
		}
		Expect(dao.GetManager().Create(&pu).Error).ToNot(HaveOccurred())
		participantUser = pu
	})

	Describe("Itemize", func() {
		Context("With participant user", func() {
			It("Should return JParticipantUser object", func() {
				j := ParticipantUser.Itemize(&participantUser)
				Expect(j.Name).To(Equal(participantUser.Name))
				Expect(j.UserUUID).To(Equal(user.UUID))
			})
		})
	})
})
