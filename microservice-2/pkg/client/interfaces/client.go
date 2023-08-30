package interfaces

import (
	"context"

	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/pb"
)

type UserClient interface {
	ListUsers(context.Context, *pb.MethodReq) (*pb.MethodRes, error)
}
