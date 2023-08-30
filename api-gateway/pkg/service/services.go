package service

import (
	"log"

	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/config"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	UserCli   pb.UserServiceClient
	MethodCli pb.MethodServiceClient
}

func InitClient(c *config.Config) (Clients, error) {
	usercc, user_err := grpc.Dial(c.UserService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if user_err != nil {
		log.Println(user_err, "inside init client")
		return Clients{}, user_err
	}
	methodcc, methoderr := grpc.Dial(c.MethodService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if methoderr != nil {
		log.Println(methoderr, "inside init client")
		return Clients{}, methoderr
	}
	userclient := pb.NewUserServiceClient(usercc)
	methodclient := pb.NewMethodServiceClient(methodcc)
	log.Println("initialized")
	return Clients{
		UserCli:   userclient,
		MethodCli: methodclient,
	}, nil
}
