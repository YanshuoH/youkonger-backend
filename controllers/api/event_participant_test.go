package api_test

import (
	. "github.com/YanshuoH/youkonger/controllers/api"

	"fmt"
	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	"github.com/YanshuoH/youkonger/test"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("EventParticipant", func() {
	var engine *gin.Engine
	var event models.Event
	var eventDate models.EventDate

	BeforeEach(func() {
		// start gin test mode engine
		r := gin.New()
		gin.SetMode(gin.TestMode)
		r.POST("/test", ApiEventParticipantUpsert)
		r.PUT("/test", ApiEventParticipantUpsert)

		engine = r

		// get the event
		e := models.Event{}
		Expect(dao.Event.First(&e, eventSet[0].ID).Error).NotTo(HaveOccurred())
		event = e

		ed := models.EventDate{}
		Expect(dao.Event.First(&ed, eventDateSet[0].ID).Error).NotTo(HaveOccurred())
		eventDate = ed
	})

	Describe("ApiEventParticipantUpsert", func() {
		Context("With invalid json form", func() {
			It("Should faild on binding step", func() {
				By("wrong json")
				j := `56789dfdg`
				resp := test.PerformRequest("POST", "/test", engine, j)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				jresp := test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.FormInvalid))

				By("missing mandatory fields")
				j = `
					{
						"eventParticipantList": [
							{
								name: "who"
							}
						]
					}
				`
				resp = test.PerformRequest("POST", "/test", engine, j)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				jresp = test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.FormInvalid))
			})
		})

		Context("With wrong data inside form", func() {
			It("Should failed on form handler", func() {
				j := `
					{
						"eventParticipantList": []
					}
				`
				resp := test.PerformRequest("POST", "/test", engine, j)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				jresp := test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.FormInvalid))
			})
		})

		Context("With correct form", func() {
			It("Should return ok", func() {
				j := `
					{
						"eventParticipantList": [
							{
								"name": "who",
								"eventDateUuid": "%s"
							}
						],
						"name": "haha"
					}
				`
				resp := test.PerformRequest("POST", "/test", engine, fmt.Sprintf(j, eventDate.UUID))
				//Expect(resp.Code).To(Equal(http.StatusOK))
				jresp := test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.OK))
			})
		})
	})
})
