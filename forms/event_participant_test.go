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
		Context("Without EM", func() {
			It("Should return an error", func() {
				f := EventParticipantForm{}
				res, cErr := f.Handle()
				Expect(res).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.NoEntityManagerInForm))
			})
		})

		Context("With none existed uuid", func() {
			It("Should return an error", func() {
				f := EventParticipantForm{
					EM:   dao.GetManager(),
					UUID: "s",
				}
				res, cErr := f.Handle()
				Expect(res).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.EventParticipantNotFound))
			})
		})

		Context("With none existed eventDate uuid", func() {
			It("Should return an error", func() {
				f := EventParticipantForm{
					EM:            dao.GetManager(),
					EventDateUUID: "s",
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
					EventDateUUID: eventDate.UUID,
					Name:          "yo",
					EM:            dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(res.Name).To(Equal("yo"))

				By("Update form")
				f = EventParticipantForm{
					EventDateUUID: eventDate.UUID,
					Name:          "haha",
					UUID:          res.UUID,
					EM:            dao.GetManager(),
				}
				res, cErr = f.Handle()
				Expect(cErr).To(BeNil())
				Expect(res.Name).To(Equal("haha"))
			})
		})
	})
})

var _ = Describe("EventParticipantForms", func() {
	Describe("Handle", func() {
		var eventDate models.EventDate
		var eventParticipant models.EventParticipant
		BeforeEach(func() {
			ed := models.EventDate{
				Time: time.Now(),
			}
			Expect(dao.Conn.Create(&ed).Error).NotTo(HaveOccurred())
			ep := models.EventParticipant{
				Name:        "bigbro",
				EventDateID: ed.ID,
			}
			Expect(dao.Conn.Create(&ep).Error).NotTo(HaveOccurred())
			eventDate = ed
			eventParticipant = ep
		})

		Context("With incorrect form", func() {
			It("Should return an error", func() {
				f := EventParticipantForms{
					Forms: []*EventParticipantForm{
						&EventParticipantForm{},
					},
				}
				res, cErr := f.Handle()
				Expect(cErr.Code).To(Equal(consts.NoEntityManagerInForm))
				Expect(res).To(HaveLen(0))
			})
		})

		Context("With correct form", func() {
			It("Should return the created/updated event participants", func() {
				var initialCount int
				dao.Conn.Raw("SELECT count(id) FROM event_participant").Row().Scan(&initialCount)

				f := EventParticipantForms{
					Forms: []*EventParticipantForm{
						&EventParticipantForm{
							EventDateUUID: eventDate.UUID,
							UUID: eventParticipant.UUID,
							Name: "hoho",
						},
						&EventParticipantForm{
							EventDateUUID: eventDate.UUID,
							Name: "haha",
						},
					},
					EM: dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(res).To(HaveLen(len(f.Forms)))

				var finalCount int
				dao.Conn.Raw("SELECT count(id) FROM event_participant").Row().Scan(&finalCount)

				Expect(finalCount).To(Equal(initialCount + 1))
			})
		})
	})
})
