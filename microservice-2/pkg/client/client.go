package client

import (
	"context"

	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/client/interfaces"
	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/config"
	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userClient struct {
	Server pb.UserServiceClient
}

func InitClient(c *config.Config) (pb.UserServiceClient, error) {
	usercc, usererr := grpc.Dial(c.UserService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if usererr != nil {
		return nil, usererr
	}
	return pb.NewUserServiceClient(usercc), nil

}

func NewUserClient(server pb.UserServiceClient) interfaces.UserClient {
	return &userClient{
		Server: server,
	}
}

func (c *userClient) ListUsers(ctx context.Context, body *pb.MethodReq) (*pb.MethodRes, error) {
	res, err := c.Server.Users(ctx, body)
	if err != nil {
		return nil, err
	}
	return res, nil
}
