package api

import (
	"github.com/gin-gonic/gin"
	"github.com/YanshuoH/youkonger/forms"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/utils"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/go-playground/log"
	"github.com/YanshuoH/youkonger/jrenders"
)

func ApiEventParticipantUpsert(c *gin.Context) {
	var f form.EventParticipantForms
	if err := binding.JSON.Bind(c.Request, &f); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewJSONResponse(consts.FormInvalid, err.Error()))
		return
	}

	// begin a transaction
	em := dao.GetManager(dao.Conn.Begin())
	f.EM = em
	eps, cErr := f.Handle()
	if cErr != nil {
		em.Rollback()
		c.JSON(http.StatusBadRequest, utils.NewJSONResponse(cErr.Code, cErr.Err.Error()))
		return
	}
	if err := em.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewJSONResponse(consts.DefaultErrorMsg, err.Error()))
		return
	}

	// trace back the event
	ep := eps[0]
	e, err := dao.Event.FindByEventParticipant(&ep)
	if err != nil {
		log.Error("Cannot retrieve event by event participant")
		c.JSON(http.StatusInternalServerError, utils.NewJSONResponse(consts.DefaultErrorMsg, err.Error()))
		return
	}
	// fully return the event
	c.JSON(http.StatusOK, utils.NewOKJSONResponse(
		jrenders.Event.Itemize(e, jrenders.EventParam{false})))
}
