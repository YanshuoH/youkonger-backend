package form_test

import (
	. "github.com/YanshuoH/youkonger/forms"

	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParticipantUser", func() {
	var event models.Event
	var user models.User
	var participateUser models.ParticipantUser
	BeforeEach(func() {
		e := models.Event{
			Title: "a titile",
		}
		Expect(dao.GetManager().Create(&e).Error).ToNot(HaveOccurred())
		event = e

		u := models.User{}
		Expect(dao.GetManager().Create(&u).Error).ToNot(HaveOccurred())
		user = u

		pu := models.ParticipantUser{
			Name:    "someone",
			EventID: e.ID,
			UserID:  u.ID,
		}
		Expect(dao.GetManager().Create(&pu).Error).ToNot(HaveOccurred())
		participateUser = pu
	})
	Describe("ParticipantUserForm", func() {
		Context("Without entity manager in form", func() {
			It("Should return an error", func() {
				f := ParticipantUserForm{}
				res, cErr := f.Handle()
				Expect(cErr.Code).To(Equal(consts.NoEntityManagerInForm))
				Expect(res).To(BeNil())
			})
		})

		Context("With unexisted event uuid", func() {
			It("Should return an error", func() {
				f := ParticipantUserForm{
					EM:        dao.GetManager(),
					EventUUID: "abc",
				}
				res, cErr := f.Handle()
				Expect(cErr.Code).To(Equal(consts.EventNotFound))
				Expect(res).To(BeNil())
			})
		})

		Context("With defined user uuid but unexisted", func() {
			It("Should return an error", func() {
				f := ParticipantUserForm{
					EM:        dao.GetManager(),
					EventUUID: event.UUID,
					UserUUID:  "abc",
				}
				res, cErr := f.Handle()
				Expect(cErr.Code).To(Equal(consts.UserNotFound))
				Expect(res).To(BeNil())
			})
		})

		Context("With correct update form", func() {
			It("Should update the participate user", func() {
				f := ParticipantUserForm{
					EM:        dao.GetManager(),
					EventUUID: event.UUID,
					UserUUID:  user.UUID,
					Name:      "anotherone",
				}
				res, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(res.Name).To(Equal(f.Name))
				Expect(res.UUID).To(Equal(participateUser.UUID))
				Expect(res.User).ToNot(BeNil())
				Expect(res.Event).ToNot(BeNil())
			})
		})

		Context("With correct insertion form", func() {
			It("Should insert a participate user", func() {
				By("User not exists yet")
				var initialUserCount int
				var initialParticipateUserCount int
				dao.Conn.Raw("SELECT count(id) FROM user").Row().Scan(&initialUserCount)
				dao.Conn.Raw("SELECT count(id) FROM participant_user").Row().Scan(&initialParticipateUserCount)

				f := ParticipantUserForm{
					EM:        dao.GetManager(),
					EventUUID: event.UUID,
					Name:      "thatuser",
				}
				res, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(res.Name).To(Equal(f.Name))
				Expect(res.Event).ToNot(BeNil())
				Expect(res.User).ToNot(BeNil())

				var finalUserCount int
				var finalParticipateUserCount int
				dao.Conn.Raw("SELECT count(id) FROM user").Row().Scan(&finalUserCount)
				dao.Conn.Raw("SELECT count(id) FROM participant_user").Row().Scan(&finalParticipateUserCount)

				Expect(finalUserCount).To(Equal(initialUserCount + 1))
				Expect(finalParticipateUserCount).To(Equal(initialParticipateUserCount + 1))
			})

			It("Should insert a participate user", func() {
				By("User already exists")
				// create another user
				u := models.User{}
				Expect(dao.GetManager().Create(&u).Error).ToNot(HaveOccurred())
				f := ParticipantUserForm{
					EM:        dao.GetManager(),
					EventUUID: event.UUID,
					UserUUID:  u.UUID,
					Name:      "thatuser",
				}

				var initialUserCount int
				var initialParticipateUserCount int
				dao.Conn.Raw("SELECT count(id) FROM user").Row().Scan(&initialUserCount)
				dao.Conn.Raw("SELECT count(id) FROM participant_user").Row().Scan(&initialParticipateUserCount)

				res, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(res.Name).To(Equal(f.Name))
				Expect(res.Event).ToNot(BeNil())
				Expect(res.User).ToNot(BeNil())

				var finalUserCount int
				var finalParticipateUserCount int
				dao.Conn.Raw("SELECT count(id) FROM user").Row().Scan(&finalUserCount)
				dao.Conn.Raw("SELECT count(id) FROM participant_user").Row().Scan(&finalParticipateUserCount)

				Expect(finalUserCount).To(Equal(initialUserCount))
				Expect(finalParticipateUserCount).To(Equal(initialParticipateUserCount + 1))
			})
		})
	})
})
