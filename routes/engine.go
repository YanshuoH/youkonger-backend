package routes

import (
	"github.com/YanshuoH/youkonger/conf"
	"github.com/YanshuoH/youkonger/controllers/api"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	config := conf.Config
	gin.SetMode(config.AppConf.GinMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// loading views
	router.LoadHTMLGlob("views/*")

	eventRouter := router.Group("/event")
	{
		eventRouter.GET("/get", api.ApiEventGet)
		eventRouter.POST("/create", api.ApiEventUpsert)
		eventRouter.POST("/update", api.ApiEventUpsert)
	}

	return router
}
