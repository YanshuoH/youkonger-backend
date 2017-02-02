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

var _ = Describe("DDayForm", func() {
	Describe("Handle", func() {
		var event models.Event
		var eventDate models.EventDate
		var dDay models.EventDate
		BeforeEach(func() {
			e := models.Event{
				Title: "something",
			}
			Expect(dao.Conn.Create(&e).Error).ToNot(HaveOccurred())
			event = e

			ed := models.EventDate{
				Time:    time.Now(),
				EventID: e.ID,
			}
			Expect(dao.Conn.Create(&ed).Error).ToNot(HaveOccurred())
			eventDate = ed

			dday := models.EventDate{
				Time:    time.Now(),
				IsDDay:  true,
				EventID: e.ID,
			}
			Expect(dao.Conn.Create(&dday).Error).ToNot(HaveOccurred())
			dDay = dday
		})

		Context("Without entity manager in form", func() {
			It("Should return an error", func() {
				f := DDayForm{}
				res, cErr := f.Handle()
				Expect(res).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.NoEntityManagerInForm))
			})
		})

		Context("With invalid uuid and hash id", func() {
			It("Should return an error", func() {
				f := DDayForm{
					UUID: event.UUID,
					Hash: "ha",
					EM:   dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(res).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.EventNotFound))
			})
		})

		Context("With invalid eventDateUUID", func() {
			It("Should return an error", func() {
				f := DDayForm{
					UUID:          event.UUID,
					Hash:          event.AdminHash,
					EventDateUUID: "something",
					EM:            dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(res).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.EventDateNotFound))
			})
		})

		Context("With valid eventDateUUID but it does not belong to the event", func() {
			It("Should return an error", func() {
				ed := &models.EventDate{
					EventID: 666,
					Time:    time.Now(),
				}
				Expect(dao.Conn.Create(ed).Error).ToNot(HaveOccurred())
				f := DDayForm{
					UUID:          event.UUID,
					Hash:          event.AdminHash,
					EventDateUUID: ed.UUID,
					EM:            dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(res).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.EventDateNotFound))
			})
		})

		Context("With correct form", func() {
			It("Should set the specified event date to dday and unset the others", func() {
				f := DDayForm{
					UUID:          event.UUID,
					Hash:          event.AdminHash,
					EventDateUUID: eventDate.UUID,
					EM:            dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(res.ID).To(Equal(event.ID))

				// reload event date
				ed, err := dao.EventDate.FindByUUID(eventDate.UUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(ed.IsDDay).To(BeTrue())

				// reload the dday
				dday, err := dao.EventDate.FindByUUID(dDay.UUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(dday.IsDDay).To(BeFalse())
			})
		})
	})
})
