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
	eventUUID := c.Param("eventUUID")
	// if event not found
	e, err := dao.Event.FindByUUID(eventUUID)
	if err != nil {
		// redirect to 404 page
		NotFound(c)
		return
	}

	// load participant user for this event
	var participantUser *models.ParticipantUser = nil
	if userUUID, err := c.Cookie(consts.UserUUIDCookieKey); err == nil {
		pu, err := dao.ParticipantUser.FindByUserUUIDAndEventUUID(userUUID, eventUUID)
		if err != nil {
			log.Infof("No participant user with user.uuid = %s, event.uuid = %s", userUUID, eventUUID)
		} else {
			participantUser = pu
		}
	}

	jsonByte, err := json.Marshal(jrenders.Event.Itemize(e, jrenders.EventParam{
		ShowHash:        true,
		ParticipantUser: participantUser,
	}))
	if err != nil {
		log.Errorf("Error occured while encoding event: +%v", err)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"InitialParticipate": string(jsonByte),
		"IsSSR":              true,
	})
}

func AdminEvent(c *gin.Context) {
	eventUUID := c.Param("eventUUID")
	eventHash := c.Param("eventHash")
	e, err := dao.Event.FindByUUIDAndAdminHash(eventUUID, eventHash)
	if err != nil {
		// redirect to 404
		NotFound(c)
		return
	}

	// render to html
	jsonByte, err := json.Marshal(jrenders.Event.Itemize(e, jrenders.EventParam{
		ShowHash: true,
	}))
	if err != nil {
		log.Errorf("Error occured while encoding event: +%v", err)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"InitialAdmin": string(jsonByte),
		"IsSSR":        true,
	})
}
