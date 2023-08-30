package service

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/domain"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/pb"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/repository/interfaces"
)

type UserServer struct {
	Repo  interfaces.UserRepo
	Mutex sync.Mutex
	pb.UnimplementedUserServiceServer
}

func NewUserServer(repo interfaces.UserRepo) pb.UserServiceServer {
	return &UserServer{
		Repo:  repo,
		Mutex: sync.Mutex{},
	}
}

func (c *UserServer) CreateUser(ctx context.Context, body *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	log.Println("inside server", body)
	res, err := c.Repo.CreateUser(ctx, domain.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		MobileNum: body.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserRes{
		Status:     http.StatusOK,
		Responseid: res,
	}, nil
}

func (c *UserServer) GetUser(ctx context.Context, body *pb.UserReq) (*pb.UserRes, error) {
	log.Println("inside service get user; userid=", body.Userid)
	res, err := c.Repo.GetUser(ctx, body.Userid)
	if err != nil {
		return nil, err
	}
	log.Println("user details", res)
	return &pb.UserRes{
		Status: http.StatusOK,
		User: &pb.User{
			ID:        uint32(res.ID),
			FirstName: res.FirstName,
			LastName:  res.LastName,
			Email:     res.Email,
			Phone:     res.MobileNum,
		},
	}, nil
}

func (c *UserServer) UpdateUser(ctx context.Context, body *pb.User) (*pb.UserRes, error) {
	log.Println("inside server update user", body)
	res, err := c.Repo.UpdateUser(ctx, domain.User{
		ID:        uint(body.ID),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		MobileNum: body.Phone,
	})
	if err != nil {
		return nil, err
	}
	log.Println("user details", res)
	return &pb.UserRes{
		Status: http.StatusOK,
		User: &pb.User{
			ID:        uint32(res.ID),
			FirstName: res.FirstName,
			LastName:  res.LastName,
			Email:     res.Email,
			Phone:     res.MobileNum,
		},
	}, nil
}

func (c *UserServer) DeleteUser(ctx context.Context, body *pb.UserReq) (*pb.DeleteUserRes, error) {
	log.Println("inside server delete user", body)
	err := c.Repo.DeleteUser(ctx, body.Userid)
	if err != nil {
		return nil, err
	}
	log.Println("deleted")
	return &pb.DeleteUserRes{
		Status: http.StatusOK,
	}, nil
}

func (c *UserServer) Users(ctx context.Context, body *pb.MethodReq) (*pb.MethodRes, error) {
	switch body.Method {
	case 1:
		c.Mutex.Lock()
		defer c.Mutex.Unlock()
		log.Println("time", body.WaitTime)
		res, err := c.Repo.Users()
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(body.WaitTime) * time.Second)
		return &pb.MethodRes{
			Status: http.StatusOK,
			User:   res,
		}, nil
	case 2:
		res, err := c.Repo.Users()
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(body.WaitTime) * time.Second)
		return &pb.MethodRes{
			Status: http.StatusOK,
			User:   res,
		}, nil
	default:
		return &pb.MethodRes{}, errors.New("wrong method number")
	}
}
