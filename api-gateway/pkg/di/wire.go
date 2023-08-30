//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/api"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/api/handler"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/client"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/config"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/service"
)

func InitializeAPI(c *config.Config) (*api.Server, error) {
	wire.Build(service.InitClient,
		client.NewUserClient, client.NewMethodClient,
		handler.NewUserHandler, handler.NewMethodHandler,
		api.NewServerHTTP)
	return &api.Server{}, nil
}
