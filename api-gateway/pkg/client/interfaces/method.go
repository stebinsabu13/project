package interfaces

import (
	"context"

	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/pb"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/utils"
)

type MethodCli interface {
	Methods(context.Context, utils.MethodReq) (*pb.MethodRes, error)
}
