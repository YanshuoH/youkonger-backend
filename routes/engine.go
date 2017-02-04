package routes

import (
	"github.com/YanshuoH/youkonger/conf"
	"github.com/YanshuoH/youkonger/controllers"
	"github.com/YanshuoH/youkonger/controllers/api"
	"github.com/YanshuoH/youkonger/controllers/middlewares"
	"github.com/gin-gonic/gin"
	"path"
)

func Setup(workspace string) *gin.Engine {
	config := conf.Config
	gin.SetMode(config.AppConf.GinMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.NoRoute(middlewares.RedirectOn404())

	// loading views
	router.LoadHTMLGlob(path.Join(workspace, "views/*.html"))
	router.Static("/assets", path.Join(workspace, "public/assets"))

	router.GET("/", controllers.Index)
	router.GET("/create", controllers.RedirectCreate)
	router.GET("/404", controllers.NotFound)

	viewRouter := router.Group("/event")
	{
		viewRouter.GET("/:eventUUID", controllers.ParticipateEvent)
		viewRouter.GET("/:eventUUID/admin/:eventHash", controllers.AdminEvent)
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

		eventDateRouter := apiRouter.Group("/eventdate")
		{
			eventDateRouter.POST("/dday", api.ApiDDay)
		}

		eventParticipantRouter := apiRouter.Group("/eventparticipant")
		{
			eventParticipantRouter.POST("/create", api.ApiEventParticipantUpsert)
			eventParticipantRouter.PUT("/update", api.ApiEventParticipantUpsert)
			eventParticipantRouter.POST("/upsert", api.ApiEventParticipantUpsert)
		}

		participantUserRouter := apiRouter.Group("/participantuser")
		{
			participantUserRouter.POST("/upsert", api.ApiEventParticipantUpsert)
		}
	}

	return router
}
