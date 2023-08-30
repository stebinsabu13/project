//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/api"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/api/service"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/config"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/db"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/repository"
)

func InitializeServe(c *config.Config) (*api.Server, error) {
	wire.Build(db.Initdb, db.InitRedis,
		repository.NewUserRepo,
		service.NewUserServer,
		api.NewgrpcServe)
	return &api.Server{}, nil
}
