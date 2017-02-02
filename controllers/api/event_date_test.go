package api_test

import (
	. "github.com/YanshuoH/youkonger/controllers/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/gin-gonic/gin"
	"github.com/YanshuoH/youkonger/models"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/test"
	"github.com/YanshuoH/youkonger/consts"
	"net/http"
	"fmt"
)

var _ = Describe("EventDate", func() {
	var engine *gin.Engine
	var event models.Event
	var eventDate models.EventDate

	BeforeEach(func() {
		r := gin.New()
		gin.SetMode(gin.TestMode)
		r.POST("/test", ApiDDay)
		engine = r

		// get the event
		e := models.Event{}
		Expect(dao.Event.First(&e, eventSet[0].ID).Error).NotTo(HaveOccurred())
		event = e

		ed := models.EventDate{}
		Expect(dao.Event.First(&ed, eventDateSet[0].ID).Error).NotTo(HaveOccurred())
		eventDate = ed
	})

	Describe("ApiDDay", func() {
		Context("With invalid form", func() {
			It("Should return an error", func() {
				By("Invalid json")
				j := `56789dfdg`
				resp := test.PerformRequest("POST", "/test", engine, j)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				jresp := test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.FormInvalid))

				By("Invalid field value")
				j = `
{
	"uuid": "%s",
	"hash": "hh",
	"eventDateUuid": "%s"
}`
				resp = test.PerformRequest("POST", "/test", engine, fmt.Sprintf(j, event.UUID, eventDate.UUID))
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				jresp = test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.EventNotFound))
			})
		})

		Context("With correct form", func() {
			It("Should return ok", func() {
				j := `
{
	"uuid": "%s",
	"hash": "%s",
	"eventDateUuid": "%s"
}`
				resp := test.PerformRequest("POST", "/test", engine, fmt.Sprintf(j, event.UUID, event.AdminHash, eventDate.UUID))
				Expect(resp.Code).To(Equal(http.StatusOK))
				jresp := test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.OK))
			})
		})
	})
})
