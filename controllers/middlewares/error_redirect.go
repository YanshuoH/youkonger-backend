package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RedirectOn404() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/404")
	}
}
