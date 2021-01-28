package application

import (
	"github.com/gorilla/mux"
	"github.com/pocketazn/BoneBox/internal/configuration"
	"github.com/pocketazn/BoneBox/internal/framework"
	"github.com/pocketazn/BoneBox/internal/routers"
	"github.com/pocketazn/BoneBox/server"
	"log"
)

type APIApplication struct {
	config *configuration.AppConfig
	Server *server.Server
	Router *mux.Router
	BoneBoxRepo framework.BoneRepoType
}

func NewAPIApplication(c *configuration.AppConfig) *APIApplication {
	bRepo := framework.NewBoneRepository()

	rootRouter := mux.NewRouter()
	r := routers.NewV1Router(c, bRepo)
	r.Register(rootRouter)

	srv := server.New(rootRouter)
	srv.Setup()

	return &APIApplication{
		config:       c,
		Server: &srv,
		BoneBoxRepo: bRepo,
		Router:       rootRouter,
	}
}

func (a *APIApplication) Run() {
	log.Println("booting server...")
	a.Server.Run()
}