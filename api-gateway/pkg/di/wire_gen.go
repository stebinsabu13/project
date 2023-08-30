// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/api"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/api/handler"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/client"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/config"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/service"
)

// Injectors from wire.go:

func InitializeAPI(c *config.Config) (*api.Server, error) {
	clients, err := service.InitClient(c)
	if err != nil {
		return nil, err
	}
	userClient := client.NewUserClient(clients)
	userHandler := handler.NewUserHandler(userClient)
	methodCli := client.NewMethodClient(clients)
	methodHandler := handler.NewMethodHandler(methodCli)
	server, err := api.NewServerHTTP(c, userHandler, methodHandler)
	if err != nil {
		return nil, err
	}
	return server, nil
}
