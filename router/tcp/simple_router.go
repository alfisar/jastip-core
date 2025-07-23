package router

import (
	simpleControll "jastip-core/application/simple/controller/tcp"

	corepb "github.com/alfisar/jastip-import/proto/core"

	"google.golang.org/grpc"
)

type simpleRouter struct {
	Controller simpleControll.SimpleGrpcController
}

func NewSimpleRouter(Controller simpleControll.SimpleGrpcController) *simpleRouter {
	return &simpleRouter{
		Controller: Controller,
	}
}

func (r *simpleRouter) simpleRouters(s *grpc.Server) {
	corepb.RegisterCheckHealthyServer(s, &r.Controller)
}
