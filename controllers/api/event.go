package api

import (
	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/forms"
	"github.com/YanshuoH/youkonger/jrenders"
	"github.com/YanshuoH/youkonger/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func ApiEventUpsert(c *gin.Context) {
	var f form.EventForm
	if err := binding.JSON.Bind(c.Request, &f); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewJSONResponse(consts.FormInvalid, err.Error()))
		return
	}

	// begin a transaction
	em := dao.GetManager(dao.Conn.Begin())
	f.EM = em
	event, cErr := f.Handle()
	if cErr != nil {
		em.Rollback()
		c.JSON(http.StatusBadRequest, utils.NewJSONResponse(cErr.Code, cErr.Err.Error()))
		return
	}
	if err := em.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewJSONResponse(consts.DefaultErrorMsg, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.NewOKJSONResponse(
		jrenders.Event.Itemize(event, jrenders.EventParam{true})))
}

func ApiEventGet(c *gin.Context) {
	f := struct {
		UUID string `form:"uuid" binding:"required"`
		Hash string `form:"hash"`
	}{}
	if err := binding.Form.Bind(c.Request, &f); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewJSONResponse(consts.FormInvalid, err.Error()))
		return
	}

	e, err := dao.Event.FindByUUID(f.UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewJSONResponse(consts.EventNotFound, err.Error()))
		return
	}

	if f.Hash != "" && f.Hash != e.AdminHash {
		c.JSON(http.StatusBadRequest, utils.NewJSONResponse(consts.InvalidAdminHash))
		return
	}

	c.JSON(http.StatusOK, utils.NewOKJSONResponse(
		jrenders.Event.Itemize(e, jrenders.EventParam{f.Hash == e.AdminHash})))
}
