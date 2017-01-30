package controllers

import (
	"encoding/json"
	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/jrenders"
	"github.com/YanshuoH/youkonger/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/log"
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
	// if event not found,
	e, err := dao.Event.FindByUUID(eventUuid)
	if err != nil {
		// redirect to 404 page
		NotFound(c)
		return
	}

	// load participant user for this event
	var participantUser *models.ParticipantUser = nil
	if participantUserUuid, err := c.Cookie(consts.ParticipantUserCookieKey); err == nil {
		pu, err := dao.ParticipantUser.FindByUUIDAndEventUUID(participantUserUuid, eventUuid)
		if err != nil {
			log.Infof("No participant user with uuid %s", participantUserUuid)
		} else {
			participantUser = pu
		}
	}

	// @TODO: handle exception
	jsonByte, _ := json.Marshal(jrenders.Event.Itemize(e, jrenders.EventParam{
		ShowHash:        true,
		ParticipantUser: participantUser,
	}))
	c.HTML(http.StatusOK, "index.html", gin.H{
		"InitialParticipate": string(jsonByte),
		"IsSSR":              true,
	})
}
