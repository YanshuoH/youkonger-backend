package form_test

import (
	. "github.com/YanshuoH/youkonger/forms"

	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("EventParticipantForm", func() {
	Describe("Handle", func() {
		var participantUser = models.ParticipantUser{
			BaseModel: models.BaseModel{
				ID: 1,
			},
			Name: "someone",
		}
		Context("Without EM", func() {
			It("Should return an error", func() {
				f := EventParticipantForm{}
				res, cErr := f.Handle()
				Expect(res).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.NoEntityManagerInForm))
			})
		})

		Context("Without participant user in form", func() {
			It("should return an error", func() {
				f := EventParticipantForm{
					EM: dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(res).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.NoParticipantUserInForm))
			})
		})

		Context("With none existed uuid", func() {
			It("Should return an error", func() {
				f := EventParticipantForm{
					EM:              dao.GetManager(),
					ParticipantUser: &participantUser,
					UUID:            "s",
				}
				res, cErr := f.Handle()
				Expect(res).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.EventParticipantNotFound))
			})
		})

		Context("With none existed eventDate uuid", func() {
			It("Should return an error", func() {
				f := EventParticipantForm{
					EM:              dao.GetManager(),
					ParticipantUser: &participantUser,
					EventDateUUID:   "s",
				}
				res, cErr := f.Handle()
				Expect(res).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.EventDateNotFound))
			})
		})

		Context("With insert form", func() {
			var eventDate models.EventDate
			BeforeEach(func() {
				ed := models.EventDate{
					Time: time.Now(),
				}
				Expect(dao.Conn.Create(&ed).Error).NotTo(HaveOccurred())
				eventDate = ed
			})

			It("Should return the created/updated event participant", func() {
				By("Insertion form")
				f := EventParticipantForm{
					EventDateUUID:   eventDate.UUID,
					ParticipantUser: &participantUser,
					EM:              dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(res.ParticipantUserID).To(Equal(participantUser.ID))

				By("Update form")
				f = EventParticipantForm{
					EventDateUUID:   eventDate.UUID,
					ParticipantUser: &participantUser,
					Remove:          true,
					UUID:            res.UUID,
					EM:              dao.GetManager(),
				}
				res, cErr = f.Handle()
				Expect(cErr).To(BeNil())
				Expect(res.UUID).To(Equal(f.UUID))
				Expect(res.Removed).To(Equal(true))
				Expect(res.ParticipantUserID).To(Equal(participantUser.ID))
			})
		})
	})
})

var _ = Describe("EventParticipantForms", func() {
	Describe("Handle", func() {
		var event models.Event
		var eventDate models.EventDate
		var eventDate2 models.EventDate
		var eventParticipant models.EventParticipant
		var user models.User
		var participantUser models.ParticipantUser

		BeforeEach(func() {
			e := models.Event{
				Title: "a title",
			}
			Expect(dao.Conn.Create(&e).Error).NotTo(HaveOccurred())
			ed := models.EventDate{
				Time: time.Now(),
			}
			ed2 := models.EventDate{
				Time: time.Now(),
			}
			Expect(dao.Conn.Create(&ed).Error).NotTo(HaveOccurred())
			Expect(dao.Conn.Create(&ed2).Error).NotTo(HaveOccurred())
			u := models.User{}
			Expect(dao.Conn.Create(&u).Error).NotTo(HaveOccurred())
			pu := models.ParticipantUser{
				Name:    "someone",
				EventID: e.ID,
				UserID:  u.ID,
			}
			Expect(dao.Conn.Create(&pu).Error).NotTo(HaveOccurred())
			ep := models.EventParticipant{
				EventDateID:       ed.ID,
				ParticipantUserID: pu.ID,
			}
			Expect(dao.Conn.Create(&ep).Error).NotTo(HaveOccurred())
			event = e
			eventDate = ed
			user = u
			participantUser = pu
			eventParticipant = ep
		})

		Context("Without entity manager in form", func() {
			It("Should return an error", func() {
				f := EventParticipantForms{
					Forms: []EventParticipantForm{
						EventParticipantForm{},
					},
				}
				res, _, cErr := f.Handle()
				Expect(cErr.Code).To(Equal(consts.NoEntityManagerInForm))
				Expect(res).To(HaveLen(0))
			})
		})

		Context("With form length equals to zero", func() {
			It("Should return an error", func() {
				f := EventParticipantForms{}
				res, _, cErr := f.Handle()
				Expect(cErr.Code).To(Equal(consts.FormInvalid))
				Expect(res).To(HaveLen(0))
			})
		})

		Context("With all nothing-to-do participant form", func() {
			It("Should return an error", func() {
				f := EventParticipantForms{
					Forms: []EventParticipantForm{
						EventParticipantForm{
							Remove: true,
						},
					},
					EM: dao.GetManager(),
				}
				res, _, cErr := f.Handle()
				Expect(cErr.Code).To(Equal(consts.FormInvalid))
				Expect(res).To(HaveLen(0))
			})
		})

		Context("With correct form", func() {
			It("Should return the created/updated event participants", func() {
				By("Update the participant user")
				var initialCount int
				dao.Conn.Raw("SELECT count(id) FROM participant_user").Row().Scan(&initialCount)

				f := EventParticipantForms{
					Forms: []EventParticipantForm{
						EventParticipantForm{
							EventDateUUID: eventDate.UUID,
							UUID:          eventParticipant.UUID,
						},
						EventParticipantForm{
							EventDateUUID: eventDate.UUID,
						},
					},
					ParticipantUserForm: ParticipantUserForm{
						Name:      "a new name",
						EventUUID: event.UUID,
						UserUUID:  user.UUID,
					},
					EM: dao.GetManager(),
				}
				res, pu, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(pu).ToNot(BeNil())
				Expect(res).To(HaveLen(len(f.Forms)))
				Expect(res[0].ParticipantUser).NotTo(BeNil())
				Expect(res[0].ParticipantUserID).To(Equal(participantUser.ID))
				Expect(res[0].ParticipantUser.Name).To(Equal(f.Name))

				var finalCount int
				dao.Conn.Raw("SELECT count(id) FROM participant_user").Row().Scan(&finalCount)

				Expect(finalCount).To(Equal(initialCount))

				By("Insert the partcipant user")
				f.ParticipantUserForm = ParticipantUserForm{
					Name:      "even newer",
					EventUUID: event.UUID,
				}
				res, cErr = f.Handle()
				dao.Conn.Raw("SELECT count(id) FROM participant_user").Row().Scan(&finalCount)
				Expect(cErr).To(BeNil())
				Expect(res[0].ParticipantUser.Name).To(Equal(f.Name))
				Expect(finalCount).To(Equal(initialCount + 1))
			})
		})

		Context("With unavailable participant user form", func() {
			var eventParticipants []models.EventParticipant
			BeforeEach(func() {
				ep1 := models.EventParticipant{
					EventDateID:       eventDate.ID,
					ParticipantUserID: participantUser.ID,
				}
				ep2 := models.EventParticipant{
					EventDateID:       eventDate2.ID,
					ParticipantUserID: participantUser.ID,
				}
				Expect(dao.GetManager().Create(&ep1).Error).NotTo(HaveOccurred())
				Expect(dao.GetManager().Create(&ep2).Error).NotTo(HaveOccurred())
				eventParticipants = append(eventParticipants, ep1, ep2)
			})
			It("Should create a unavailable participant user and remove all event_participant", func() {
				var initialCount int
				dao.Conn.Raw("SELECT COUNT(id) FROM event_participant WHERE participant_user_id = ? AND removed = FALSE",
					participantUser.ID).Row().Scan(&initialCount)
				Expect(initialCount).ToNot(Equal(0))

				f := EventParticipantForms{
					ParticipantUserForm: ParticipantUserForm{
						Name:        "a new name",
						EventUUID:   event.UUID,
						UserUUID:    user.UUID,
						Unavailable: true,
					},
					EM: dao.GetManager(),
				}
				res, pu, err := f.Handle()
				Expect(err).To(BeNil())
				Expect(pu).ToNot(BeNil())
				Expect(pu.Unavailable).To(BeTrue())
				Expect(res).To(HaveLen(0))

				var finalCount int
				dao.Conn.Raw("SELECT COUNT(id) FROM event_participant WHERE participant_user_id = ? AND removed = FALSE",
					participantUser.ID).Row().Scan(&finalCount)
				Expect(finalCount).To(Equal(0))
			})
		})
	})
})
