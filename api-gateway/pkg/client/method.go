package client

import (
	"context"

	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/client/interfaces"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/pb"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/service"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/utils"
)

type methodClient struct {
	Server pb.MethodServiceClient
}

func NewMethodClient(server service.Clients) interfaces.MethodCli {
	return &methodClient{
		Server: server.MethodCli,
	}
}

func (c *methodClient) Methods(ctx context.Context, body utils.MethodReq) (*pb.MethodRes, error) {
	res, err := c.Server.ListUsers(ctx, &pb.MethodReq{
		Method:   int32(body.Method),
		WaitTime: int32(body.WaitTime),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
