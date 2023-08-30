package client

import (
	"context"
	"log"

	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/client/interfaces"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/pb"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/service"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/utils"
)

type userClient struct {
	Server pb.UserServiceClient
}

func NewUserClient(server service.Clients) interfaces.UserClient {
	return &userClient{
		Server: server.UserCli,
	}
}

func (c *userClient) UserCreate(ctx context.Context, body utils.UserDetails) (*pb.CreateUserRes, error) {
	log.Println("inside client", body)
	res, err := c.Server.CreateUser(ctx, &pb.CreateUserReq{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Phone:     body.MobileNum,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userClient) GetUserById(ctx context.Context, userid string) (*pb.UserRes, error) {
	res, err := c.Server.GetUser(ctx, &pb.UserReq{
		Userid: userid,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userClient) UpdateUser(ctx context.Context, userid int, body utils.UserDetails) (*pb.UserRes, error) {
	log.Println("inside update client", userid, body)
	res, err := c.Server.UpdateUser(ctx, &pb.User{
		ID:        uint32(userid),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Phone:     body.MobileNum,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userClient) DeleteUser(ctx context.Context, userid string) (*pb.DeleteUserRes, error) {
	res, err := c.Server.DeleteUser(ctx, &pb.UserReq{
		Userid: userid,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
