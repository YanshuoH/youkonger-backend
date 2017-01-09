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

var _ = Describe("EventDateForm", func() {

	Describe("Handle", func() {
		Context("Without entity manager in form", func() {
			It("Should return an error", func() {
				f := EventDateForm{}
				ed, cErr := f.Handle()
				Expect(ed).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.NoEntityManagerInForm))
			})
		})

		Context("With negative timeInUnix", func() {
			It("Should return an error", func() {
				f := EventDateForm{
					TimeInUnix: -1,
					EM:         dao.GetManager(),
				}
				ed, cErr := f.Handle()
				Expect(ed).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.IncorrectUnixTime))
			})
		})

		Context("Without eventUuid", func() {
			It("Should return an error", func() {
				f := EventDateForm{
					TimeInUnix: 123,
					EM:         dao.GetManager(),
				}
				ed, cErr := f.Handle()
				Expect(ed).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.FormInvalid))
			})
		})

		Context("With incorrect eventDate uuid", func() {
			It("Should return an error", func() {
				f := EventDateForm{
					EM:         dao.GetManager(),
					TimeInUnix: time.Now().Unix(),
					UUID:       "s",
					EventUUID:  "d",
				}
				ed, cErr := f.Handle()
				Expect(ed).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.EventDateNotFound))
			})
		})

		Context("With incorrect event uuid", func() {
			It("Should return an error", func() {
				f := EventDateForm{
					EM:         dao.GetManager(),
					TimeInUnix: time.Now().Unix(),
					EventUUID:  "s",
				}
				ed, cErr := f.Handle()
				Expect(ed).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.EventNotFound))
			})
		})

		Context("With correct form", func() {
			var event models.Event

			BeforeEach(func() {
				e := models.Event{
					Title:       "t",
					Description: "d",
					Location:    "l",
				}
				Expect(dao.Conn.Create(&e).Error).NotTo(HaveOccurred())
				event = e
			})

			It("Should return the newly created eventDate when inserting", func() {
				f := EventDateForm{
					EM:         dao.GetManager(),
					TimeInUnix: time.Now().Unix(),
					EventUUID:  event.UUID,
				}
				ed, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(ed.EventID).To(Equal(event.ID))
				Expect(ed.Time.Unix()).To(Equal(f.TimeInUnix))
			})

			It("Should return the updated eventDate when updating", func() {
				// create a new eventDate
				ed := models.EventDate{
					Time:    time.Now(),
					EventID: event.ID,
				}
				Expect(dao.Conn.Create(&ed).Error).NotTo(HaveOccurred())
				t, _ := time.Parse("2006-01-02", "2017-01-06")
				f := EventDateForm{
					EM:         dao.GetManager(),
					TimeInUnix: t.Unix(),
					UUID:       ed.UUID,
					EventUUID:  event.UUID,
				}
				updatedED, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(updatedED.Time.Format("2006-01-02")).To(Equal("2017-01-06"))
				Expect(updatedED.ID).To(Equal(ed.ID))
			})
		})
	})
})

var _ = Describe("EventDateForms", func() {
	Describe("Handle", func() {
		var event models.Event
		var eventDate models.EventDate

		Context("Without entity manager or event entity in form", func() {
			It("Should return an error", func() {
				f := EventDateForms{}
				res, cErr := f.Handle()
				Expect(res).To(HaveLen(0))
				Expect(cErr.Code).To(Equal(consts.DefaultErrorMsg))
			})
		})

		Context("With invalid eventDate in list", func() {
			It("Should return an error", func() {
				// create a new event and eventDate
				e := models.Event{
					Title:       "t",
					Description: "d",
					Location:    "l",
				}
				Expect(dao.Conn.Create(&e).Error).NotTo(HaveOccurred())

				f := EventDateForms{
					Forms: []*EventDateForm{
						&EventDateForm{
							TimeInUnix: -123,
						},
					},
					EM:    dao.GetManager(),
					Event: &e,
				}
				res, cErr := f.Handle()
				Expect(res).To(HaveLen(0))
				Expect(cErr.Code).To(Equal(consts.IncorrectUnixTime))
			})
		})

		Context("With valid eventDate list", func() {
			BeforeEach(func() {
				// create a new event and eventDate
				e := models.Event{
					Title:       "t",
					Description: "d",
					Location:    "l",
				}
				Expect(dao.Conn.Create(&e).Error).NotTo(HaveOccurred())
				ed := models.EventDate{
					EventID: e.ID,
					Time:    time.Now(),
				}
				Expect(dao.Conn.Create(&ed).Error).NotTo(HaveOccurred())
				event = e
				eventDate = ed
			})

			It("Should return the list of eventDates", func() {
				var initialCount int
				dao.Conn.Raw("SELECT count(id) FROM event_date").Row().Scan(&initialCount)
				f := EventDateForms{
					Forms: []*EventDateForm{
						&EventDateForm{
							UUID:       eventDate.UUID,
							EventUUID:  event.UUID,
							TimeInUnix: 123,
						},
						&EventDateForm{
							EventUUID:  event.UUID,
							TimeInUnix: 456,
						},
					},
					Event: &event,
					EM:    dao.GetManager(),
				}
				res, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(res).To(HaveLen(len(f.Forms)))

				var finalCount int
				dao.Conn.Raw("SELECT count(id) FROM event_date").Row().Scan(&finalCount)

				Expect(finalCount).To(Equal(initialCount + 1))
			})
		})
	})
})
