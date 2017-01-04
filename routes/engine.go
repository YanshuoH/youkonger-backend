package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/YanshuoH/youkonger/conf"
	"github.com/YanshuoH/youkonger/controllers"
)

func Setup() *gin.Engine {
	config := conf.Config
	gin.SetMode(config.AppConf.GinMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// loading views
	router.LoadHTMLGlob("views/*")

	router.GET("/ok", controllers.OK)

	return router
}
