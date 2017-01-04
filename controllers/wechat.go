package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}