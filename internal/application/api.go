package application

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pocketazn/BoneBox/internal/configuration"
	"github.com/pocketazn/BoneBox/internal/framework"
	"github.com/pocketazn/BoneBox/internal/routers"
	"github.com/pocketazn/BoneBox/server"
	"log"
)

type APIApplication struct {
	config      *configuration.AppConfig
	Server      *server.Server
	Router      *mux.Router
	BoneBoxRepo framework.DataAccessor
}

func NewAPIApplication(c *configuration.AppConfig) *APIApplication {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		//TODO HANDLE THIS
	}
	repo := framework.NewDataAccess(*db)

	rootRouter := mux.NewRouter()
	r := routers.NewV1Router(c, &repo)
	r.Register(rootRouter)

	srv := server.New(rootRouter)
	srv.Setup()

	return &APIApplication{
		config:      c,
		Server:      &srv,
		BoneBoxRepo: &repo,
		Router:      rootRouter,
	}
}

func (a *APIApplication) Run() {
	log.Println("booting server...")
	a.Server.Run()
}
