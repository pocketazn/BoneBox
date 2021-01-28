package main

import (
	"github.com/pocketazn/BoneBox/internal/application"
	"github.com/pocketazn/BoneBox/internal/configuration"
	"log"
)

func main() {
	c, err := configuration.Configure()
	if err != nil {
		log.Fatal(err)
	}

	app := application.NewAPIApplication(c)
	app.Run()
}
