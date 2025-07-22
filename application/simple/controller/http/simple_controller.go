package http

import (
	pb "jastip-core/proto"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type simpleController struct {
}

func NewSimpleController() *simpleController {
	return &simpleController{}
}

func (c *simpleController) Healthy(ctx *fiber.Ctx) error {

	_ = domain.DataPool
	ctx.Status(fasthttp.StatusOK).JSON(domain.ErrorData{
		Status:  "success",
		Code:    0,
		Message: "Welcome to API Core Justip.in version 1.0, enjoy and chersss :)",
	})
	return nil
}

func (c *simpleController) HealthyGRPC(ctx *fiber.Ctx) error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		ctx.Status(fasthttp.StatusInternalServerError).JSON(domain.ErrorData{
			Status:  "error",
			Code:    errorhandler.ErrCodeInternalServer,
			Message: "Cannot connect GRPC",
			Errors:  err.Error(),
		})
	}

	defer conn.Close()
	grpcClient := pb.NewCheckHealthyClient(conn)

	res, err := grpcClient.CheckRunning(ctx.Context(), &emptypb.Empty{})
	if err != nil {
		return ctx.Status(500).SendString("gRPC error: " + err.Error())
	}
	ctx.Status(fasthttp.StatusOK).JSON(domain.ErrorData{
		Status:  "success",
		Code:    0,
		Message: res.Message,
	})
	return nil

}
