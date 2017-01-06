package main

import (
	"flag"
	"github.com/YanshuoH/youkonger/consts"
	"github.com/go-playground/log"
	"github.com/YanshuoH/youkonger/conf"
	"github.com/YanshuoH/youkonger/routes"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	env := flag.String("env", "loc", "Specify your env")
	file := flag.String("conf", "./conf/conf_loc.toml", "Specify config file")

	listenPort := flag.String("port", consts.DefaultPort, "Specify the port")
	flag.Parse()

	log.Info("Running in %s", env)

	c, err := conf.Setup(*file)
	if err != nil {
		panic(err)
	}
	dao.Connect(c.DbConf.Dsn)
	// optional
	dao.AutoMigration()

	// enable in debug mode
	dao.Conn.LogMode(c.AppConf.GinMode == gin.DebugMode)

	routes.Setup().Run(":" + *listenPort)
}
