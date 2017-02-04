package main

import (
	"flag"
	"github.com/YanshuoH/youkonger/conf"
	"github.com/YanshuoH/youkonger/consts"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/jrenders"
	"github.com/YanshuoH/youkonger/routes"
	"github.com/YanshuoH/youkonger/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
)

func main() {
	env := flag.String("env", "loc", "Specify your env")
	file := flag.String("conf", "./conf/conf_loc.toml", "Specify config file")
	workspaceFlag := flag.String("workspace", "./", "Specify workspace, useful for static/view files")
	listenPort := flag.String("port", consts.DefaultPort, "Specify the port")
	flag.Parse()

	// configure workspace
	var err error
	var workspace string
	if workspace, err = utils.GetAbsFilePath(*workspaceFlag); err != nil {
		panic(err)
	}

	// setup logger
	cLog := console.New()
	log.RegisterHandler(cLog, log.AllLevels...)

	log.Noticef("Current workspace = %s", workspace)
	log.Infof("Running in %s", *env)

	c, err := conf.Setup(*file)
	if err != nil {
		panic(err)
	}
	dao.Connect(c.DbConf.Dsn)
	// optional
	dao.AutoMigration()

	// enable in debug mode
	dao.Conn.LogMode(c.AppConf.GinMode == gin.DebugMode)

	// register jrenders
	jrenders.Register()

	routes.Setup(workspace).Run(":" + *listenPort)
}
