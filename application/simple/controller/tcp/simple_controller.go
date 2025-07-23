package tcp

import (
	"context"

	corepb "github.com/alfisar/jastip-import/proto/core"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SimpleGrpcController struct {
	corepb.UnimplementedCheckHealthyServer
}

func NewSimpleController() *SimpleGrpcController {
	return &SimpleGrpcController{}
}

func (h *SimpleGrpcController) CheckRunning(ctx context.Context, _ *emptypb.Empty) (*corepb.Healthy, error) {
	return &corepb.Healthy{
		Message: "Welcome to gRPC Core Jastip.in version 1.0, enjoy and chersss :)",
	}, nil
}
