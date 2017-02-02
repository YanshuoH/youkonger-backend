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

func ApiDDay(c *gin.Context) {
	f := form.DDayForm{}
	if err := binding.JSON.Bind(c.Request, &f); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewJSONResponse(consts.FormInvalid, err.Error()))
		return
	}

	// begin a tx
	em := dao.GetManager(dao.Conn.Begin())
	f.EM = em
	event, cErr := f.Handle()
	if cErr != nil {
		em.Rollback()
		c.JSON(http.StatusBadRequest, utils.NewJSONResponse(cErr.Code, cErr.Err.Error()))
		return
	}
	if err := f.EM.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewJSONResponse(consts.DefaultErrorMsg, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.NewOKJSONResponse(
		jrenders.Event.Itemize(event, jrenders.EventParam{true, nil})))
}
