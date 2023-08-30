package service

import (
	"context"

	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/client/interfaces"
	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/pb"
)

type MethodServer struct {
	UserClient interfaces.UserClient
	pb.UnimplementedMethodServiceServer
}

func NewMethodServer(client interfaces.UserClient) pb.MethodServiceServer {
	return &MethodServer{
		UserClient: client,
	}
}

func (c *MethodServer) ListUsers(ctx context.Context, body *pb.MethodReq) (*pb.MethodRes, error) {
	res, err := c.UserClient.ListUsers(ctx, body)
	if err != nil {
		return nil, err
	}
	return res, nil
}
