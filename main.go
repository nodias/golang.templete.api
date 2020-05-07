package main

import (
	"./database"
	"./router"
	"github.com/nodias/golang.templete.common/logger"
	"github.com/nodias/golang.templete.common/middleware"
	"github.com/nodias/golang.templete.common/model"

	"github.com/urfave/negroni"
)

var config model.TomlConfig

func init() {
	model.Load("config/%s/config.toml")
	config = *model.GetConfig()
	logger.Init()
	database.Init()
	database.NewOpenDB()
}

func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.Logging(config.Logconfig.Logpath)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["ApmExam3"].PORT)
}
