package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/YanshuoH/youkonger/conf"
	"github.com/YanshuoH/youkonger/controllers/api"
	"github.com/YanshuoH/youkonger/controllers"
	"github.com/YanshuoH/youkonger/controllers/middlewares"
)

func Setup() *gin.Engine {
	config := conf.Config
	gin.SetMode(config.AppConf.GinMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.NoRoute(middlewares.RedirectOn404())

	// loading views
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./public/assets")

	router.GET("/", controllers.Index)
	router.GET("/create", controllers.RedirectCreate)
	router.GET("/404", controllers.NotFound)

	viewRouter := router.Group("/event")
	{
		viewRouter.GET("/:eventUuid", controllers.ParticipateEvent)
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
