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
				Expect(res.ParticipantUserId).To(Equal(participantUser.ID))

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
				Expect(res.ParticipantUserId).To(Equal(participantUser.ID))
			})
		})
	})
})

var _ = Describe("EventParticipantForms", func() {
	Describe("Handle", func() {
		var event models.Event
		var eventDate models.EventDate
		var eventParticipant models.EventParticipant
		var participantUser models.ParticipantUser
		BeforeEach(func() {
			e := models.Event{
				Title: "a title",
			}
			Expect(dao.Conn.Create(&e).Error).NotTo(HaveOccurred())
			ed := models.EventDate{
				Time: time.Now(),
			}
			Expect(dao.Conn.Create(&ed).Error).NotTo(HaveOccurred())
			pu := models.ParticipantUser{
				Name:    "someone",
				EventId: e.ID,
			}
			Expect(dao.Conn.Create(&pu).Error).NotTo(HaveOccurred())
			ep := models.EventParticipant{
				EventDateID:       ed.ID,
				ParticipantUserId: pu.ID,
			}
			Expect(dao.Conn.Create(&ep).Error).NotTo(HaveOccurred())
			event = e
			eventDate = ed
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
				res, cErr := f.Handle()
				Expect(cErr.Code).To(Equal(consts.NoEntityManagerInForm))
				Expect(res).To(HaveLen(0))
			})
		})

		Context("With form length equals to zero", func() {
			It("Should return an error", func() {
				f := EventParticipantForms{}
				res, cErr := f.Handle()
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
					Name: "some",
					UUID: "abc",
					EM:   dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(cErr.Code).To(Equal(consts.FormInvalid))
				Expect(res).To(HaveLen(0))
			})
		})

		Context("With unexisted event uuid", func() {
			It("Should return an error", func() {
				f := EventParticipantForms{
					Forms: []EventParticipantForm{
						EventParticipantForm{},
					},
					Name:      "some",
					UUID:      "abc",
					EventUUID: "cbd",
					EM:        dao.GetManager(),
				}
				_, cErr := f.Handle()
				Expect(cErr.Code).To(Equal(consts.EventNotFound))
			})
		})

		Context("With unexisted participant user", func() {
			It("Should return an error", func() {
				f := EventParticipantForms{
					Forms: []EventParticipantForm{
						EventParticipantForm{},
					},
					Name:      "some",
					UUID:      "abc",
					EventUUID: event.UUID,
					EM:        dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(cErr.Code).To(Equal(consts.ParticipantUserNotFound))
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
					Name:      "a new name",
					EventUUID: event.UUID,
					UUID:      participantUser.UUID,
					EM:        dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(res).To(HaveLen(len(f.Forms)))
				Expect(res[0].ParticipantUser).NotTo(BeNil())
				Expect(res[0].ParticipantUserId).To(Equal(participantUser.ID))
				Expect(res[0].ParticipantUser.Name).To(Equal(f.Name))

				var finalCount int
				dao.Conn.Raw("SELECT count(id) FROM participant_user").Row().Scan(&finalCount)

				Expect(finalCount).To(Equal(initialCount))

				By("Insert the partcipant user")
				f.UUID = ""
				res, cErr = f.Handle()
				dao.Conn.Raw("SELECT count(id) FROM participant_user").Row().Scan(&finalCount)
				Expect(cErr).To(BeNil())
				Expect(res[0].ParticipantUser.Name).To(Equal(f.Name))
				Expect(finalCount).To(Equal(initialCount + 1))
			})
		})
	})
})
