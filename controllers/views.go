package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/YanshuoH/youkonger/jrenders"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	"encoding/json"
)

func OK(c *gin.Context) {
	e := &models.Event{}
	dao.Event.First(e)

	// @TODO: handle exception
	jsonByte, _ := json.Marshal(jrenders.Event.Itemize(e, jrenders.EventParam{ShowHash: true}))
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Initial": string(jsonByte),
	})
}
