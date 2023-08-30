package interfaces

import (
	"context"

	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/domain"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/pb"
)

type UserRepo interface {
	CreateUser(context.Context, domain.User) (string, error)
	GetUser(context.Context, string) (domain.User, error)
	UpdateUser(context.Context, domain.User) (domain.User, error)
	DeleteUser(context.Context, string) error
	Users() ([]*pb.User, error)
}
