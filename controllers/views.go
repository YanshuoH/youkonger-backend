package controllers

import (
	"encoding/json"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/jrenders"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{})
}

func RedirectCreate(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func ParticipateEvent(c *gin.Context) {
	eventUuid := c.Param("eventUuid")

	e, err := dao.Event.FindByUUID(eventUuid)
	if err != nil {
		// redirect to 404 page
		NotFound(c)
		return
	}

	// @TODO: handle exception
	jsonByte, _ := json.Marshal(jrenders.Event.Itemize(e, jrenders.EventParam{ShowHash: true}))
	c.HTML(http.StatusOK, "index.html", gin.H{
		"InitialParticipate": string(jsonByte),
		"IsSSR":              true,
	})
}
