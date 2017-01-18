package middlewares_test

import (
	. "github.com/YanshuoH/youkonger/controllers/middlewares"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/gin-gonic/gin"
	"github.com/YanshuoH/youkonger/test"
	"net/http"
	"github.com/YanshuoH/youkonger/controllers"
)

var _ = Describe("ErrorRedirect", func() {
	Describe("RedirectOn404", func() {
		Context("With calling this handler", func() {
			var router *gin.Engine
			BeforeEach(func() {
				router = gin.New()
				gin.SetMode(gin.TestMode)
				router.Use(RedirectOn404())
				router.GET("/404", controllers.NotFound)
			})
			It("Should render 404 page", func() {
				resp := test.PerformRequest("GET", "/something", router)
				Expect(resp.Code).To(Equal(http.StatusTemporaryRedirect))
			})
		})
	})
})
