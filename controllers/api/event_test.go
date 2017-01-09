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

var _ = Describe("Event", func() {
	var engine *gin.Engine
	var event models.Event

	BeforeEach(func() {
		// start gin test mode engine
		r := gin.New()
		gin.SetMode(gin.TestMode)
		r.GET("/test", ApiEventGet)
		r.POST("/test", ApiEventUpsert)
		r.PUT("/test", ApiEventUpsert)

		engine = r

		// get the event
		e := models.Event{}
		Expect(dao.Event.First(&e, eventSet[0].ID).Error).NotTo(HaveOccurred())
		event = e
	})

	Describe("ApiEventGet", func() {
		Context("Without mandatory query params", func() {
			It("Should return BadRequest", func() {
				resp := test.PerformRequest("GET", "/test", engine)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				jresp := test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.FormInvalid))
			})
		})

		Context("With wrong event uuid", func() {
			It("Should return EventNotFound", func() {
				resp := test.PerformRequest("GET", "/test?uuid=a", engine)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				jresp := test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.EventNotFound))
			})
		})

		Context("With wrong admin uuid", func() {
			It("Should return InvalidAdminHash", func() {
				resp := test.PerformRequest("GET", fmt.Sprintf("/test?uuid=%s&hash=%s", event.UUID, "h"), engine)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				jresp := test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.InvalidAdminHash))
			})
		})

		Context("With right uuid and right admin hash", func() {
			It("Should return OK", func() {
				resp := test.PerformRequest("GET", fmt.Sprintf("/test?uuid=%s&hash=%s", event.UUID, event.AdminHash), engine)
				Expect(resp.Code).To(Equal(http.StatusOK))
				jresp := test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.OK))
			})
		})
	})

	Describe("ApiEventUpsert", func() {
		Context("Without mandatory fields or with invalid json form", func() {
			It("Should failed on binding step", func() {
				By("Invalid json form")
				j := `{jj}`
				resp := test.PerformRequest("POST", "/test", engine, j)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				jresp := test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.FormInvalid))

				By("Missing mandatory field")
				j = `
					{
						"description": "ha",
						"location": "yo"
					}
				`
				resp = test.PerformRequest("POST", "/test", engine, j)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				jresp = test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.FormInvalid))
			})
		})

		Context("With wrong form datas, eg. uuid invalid", func() {
			It("Should failed on form handler", func() {
				j := `
					{
						"title": "title",
						"uuid": "lol",
						"hash": "lol",
						"description": "ha",
						"location": "yo"
					}
				`
				resp := test.PerformRequest("PUT", "/test", engine, j)
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
				jresp := test.ReadJsonResponse(resp)
				Expect(jresp.ResultCode).To(Equal(consts.EventNotFound))
			})
		})

		Context("With correct form", func() {
			It("Should return ok as code", func() {
				// create it
				j := `
					{
						"title": "title",
						"description": "ha",
						"location": "yo",
						"eventDateList": [
							{
								"timeInUnix": 45678
							}
						]
					}
				`
				resp := test.PerformRequest("POST", "/test", engine, j)
				//Expect(resp.Code).To(Equal(http.StatusOK))
				jresp := test.ReadJsonResponse(resp)
				fmt.Println(jresp.Data)
				Expect(jresp.ResultCode).To(Equal(consts.OK))
			})
		})
	})
})
