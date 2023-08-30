//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/api"
	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/api/service"
	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/client"
	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/config"
)

func InitializeServe(c *config.Config) (*api.Server, error) {
	wire.Build(client.InitClient,
		client.NewUserClient,
		service.NewMethodServer,
		api.NewgrpcServe)
	return &api.Server{}, nil
}
