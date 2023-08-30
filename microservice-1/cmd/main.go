package main

import (
	"log"

	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/config"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/di"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config", err.Error())
	}
	server, err1 := di.InitializeServe(c)
	if err1 != nil {
		log.Fatal("failed to init server", err1.Error())
	}
	if err := server.Start(); err != nil {
		log.Fatal("couldn't start the server")
	}
}
