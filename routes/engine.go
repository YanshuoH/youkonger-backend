package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/YanshuoH/youkonger/conf"
	"github.com/YanshuoH/youkonger/controllers/api"
	"github.com/YanshuoH/youkonger/controllers"
)

func Setup() *gin.Engine {
	config := conf.Config
	gin.SetMode(config.AppConf.GinMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// loading views
	router.LoadHTMLGlob("views/index.html")
	router.Static("/assets", "./public/assets")

	viewRouter := router.Group("/")
	{
		viewRouter.GET("/", controllers.OK)
	}

	apiRouter := router.Group("/api")
	{
		eventRouter := apiRouter.Group("/event")
		{
			eventRouter.GET("/get", api.ApiEventGet)
			eventRouter.POST("/create", api.ApiEventUpsert)
			eventRouter.PUT("/update", api.ApiEventUpsert)
			eventRouter.POST("/upsert", api.ApiEventUpsert)
		}

		eventParticipantRouter := apiRouter.Group("/eventparticipant")
		{
			eventParticipantRouter.POST("/create", api.ApiEventParticipantUpsert)
			eventParticipantRouter.PUT("/update", api.ApiEventParticipantUpsert)
		}
	}

	return router
}
