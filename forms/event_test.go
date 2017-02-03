package form_test

import (
	. "github.com/YanshuoH/youkonger/forms"

	"bytes"
	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"time"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

var _ = Describe("EventForm", func() {
	Describe("Handle", func() {
		var event models.Event
		BeforeEach(func() {
			toInsert := models.Event{
				Title:       "t",
				Description: "d",
				Location:    "l",
			}
			Expect(dao.Conn.Create(&toInsert).Error).NotTo(HaveOccurred())
			event = toInsert
		})

		Context("Without entity manager in form", func() {
			It("Should reutnr an error", func() {
				f := EventForm{}
				e, cErr := f.Handle()
				Expect(e).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.NoEntityManagerInForm))
			})
		})

		Context("With fields too long", func() {
			It("Should return an error", func() {
				By("Title too long")
				titleLength := consts.TitleLengthConstraint + 1
				var titleBuf bytes.Buffer
				for i := 0; i < titleLength; i++ {
					titleBuf.WriteString("啊")
				}
				f := EventForm{
					EM:    dao.GetManager(),
					Title: titleBuf.String(),
					UUID:  "s",
				}
				e, cErr := f.Handle()
				Expect(e).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.TileTooLong))

				By("Description too long")
				descLength := consts.DescriptionLengthConstraint + 1
				var descBuf bytes.Buffer
				for i := 0; i < descLength; i++ {
					descBuf.WriteString("哈")
				}
				f = EventForm{
					EM:          dao.GetManager(),
					Title:       "title",
					UUID:        "s",
					Description: descBuf.String(),
				}
				e, cErr = f.Handle()
				Expect(e).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.DescriptionTooLong))

				By("Location too long")
				locLength := consts.LocationLengthConstraint + 1
				var locBuf bytes.Buffer
				for i := 0; i < locLength; i++ {
					locBuf.WriteString("哟")
				}
				f = EventForm{
					EM:          dao.GetManager(),
					Title:       "title",
					UUID:        "s",
					Description: "d",
					Location:    locBuf.String(),
				}
				e, cErr = f.Handle()
				Expect(e).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.LocationTooLong))
			})
		})

		Context("With uuid but without admin hash in form", func() {
			It("Should return an error", func() {
				f := EventForm{
					EM:   dao.GetManager(),
					UUID: "s",
				}
				e, cErr := f.Handle()
				Expect(e).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.InvalidAdminHash))
			})
		})

		Context("With wrong uuid", func() {
			It("Should return an error", func() {
				f := EventForm{
					UUID:      "ha",
					AdminHash: "yo",
					EM:        dao.GetManager(),
				}
				e, cErr := f.Handle()
				Expect(e).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.EventNotFound))
			})
		})

		Context("With correct uuid but wrong admin hash", func() {
			It("Should return an error", func() {
				f := EventForm{
					UUID:      event.UUID,
					AdminHash: "yo",
					EM:        dao.GetManager(),
				}
				e, cErr := f.Handle()
				Expect(e).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.InvalidAdminHash))
			})
		})

		Context("With valid form and update symptom", func() {
			It("Should update the given event with title, desc and loc", func() {
				f := EventForm{
					UUID:        event.UUID,
					AdminHash:   event.AdminHash,
					Title:       "title",
					Description: "description",
					Location:    "",
					EM:          dao.GetManager(),
				}

				e, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(e.Title).To(Equal("title"))
				Expect(e.Location).To(Equal(""))
				Expect(e.ID).To(Equal(event.ID))
			})
		})

		Context("With valid insertion form but without eventDates", func() {
			It("Should return an error", func() {
				f := EventForm{
					Title:       "title",
					Description: "desc",
					Location:    "loc",
					EM:          dao.GetManager(),
				}
				e, cErr := f.Handle()
				Expect(e).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.EventDateRequiredWhenInitializing))
			})
		})

		Context("With valid insertion form and broken eventDates", func() {
			It("Should return an error", func() {
				f := EventForm{
					Title:       "title",
					Description: "desc",
					Location:    "location",
					EM:          dao.GetManager(),
					EventDateForms: EventDateForms{
						Forms: []*EventDateForm{
							&EventDateForm{
								TimeInUnix: -1,
							},
						},
					},
				}
				e, cErr := f.Handle()
				Expect(e).To(BeNil())
				Expect(cErr.Code).To(Equal(consts.IncorrectUnixTime))
			})
		})

		Context("With valid insertion form and valid eventDates", func() {
			It("Should return new event and its eventDates", func() {
				f := EventForm{
					Title:       "title",
					Description: "desc",
					Location:    "location",
					EM:          dao.GetManager(),
					EventDateForms: EventDateForms{
						Forms: []*EventDateForm{
							&EventDateForm{
								TimeInUnix: time.Now().Unix(),
							},
						},
					},
				}
				e, cErr := f.Handle()
				Expect(cErr).To(BeNil())
				Expect(e.Title).To(Equal("title"))
				Expect(len(e.EventDates)).To(Equal(1))
			})
		})
	})
})
