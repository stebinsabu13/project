package api

import (
	"fmt"
	"net"

	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/config"
	"github.com/stebin13/x-tentioncrew/microservice-2/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	gs   *grpc.Server
	Lis  net.Listener
	Port string
}

func NewgrpcServe(c *config.Config, service pb.MethodServiceServer) (*Server, error) {
	grpcserver := grpc.NewServer()
	pb.RegisterMethodServiceServer(grpcserver, service)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		return nil, err
	}
	return &Server{
		gs:   grpcserver,
		Lis:  lis,
		Port: c.Port,
	}, nil
}

func (s *Server) Start() error {
	fmt.Println("Method service on:", s.Port)
	return s.gs.Serve(s.Lis)
}
