package main

import (
	"log"

	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/config"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/di"
)

func main() {
	c, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("cannot load config:", configerr)
	}
	log.Println("method service", c.MethodService, "port", c.Port, "user service", c.UserService)
	server, dierr := di.InitializeAPI(c)
	if dierr != nil {
		log.Fatal("cannot initialize server", dierr)
	}
	server.Start()
}
