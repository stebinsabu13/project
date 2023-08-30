package interfaces

import (
	"context"

	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/pb"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/utils"
)

type UserClient interface {
	UserCreate(context.Context, utils.UserDetails) (*pb.CreateUserRes, error)
	GetUserById(context.Context, string) (*pb.UserRes, error)
	UpdateUser(context.Context, int, utils.UserDetails) (*pb.UserRes, error)
	DeleteUser(context.Context, string) (*pb.DeleteUserRes, error)
}
