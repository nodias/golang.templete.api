package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/nodias/golang.templete.api/interface/rest"
	"github.com/nodias/golang.templete.api/registry"
	"github.com/nodias/golang.templete.common/models"
	"github.com/nodias/golang.templete.common/shared/logger"
	"github.com/nodias/golang.templete.common/shared/repository"

	"github.com/urfave/negroni"
)

var config models.TomlConfig

func init() {
	config.Load("config/%s/config.toml")
	config = *models.GetConfig()
	logger.Init()
	repository.Init()
	repository.NewOpenDB()
}

func main() {
	log := logger.New(context.Background())

	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container %v", err)
	}

	router := mux.NewRouter()

	rest.Apply(router, ctn)

	n := negroni.New()
	n.UseHandler(router)
	log.Info("Server - Server On!")
	n.Run(config.Servers["api"].PORT)
}


