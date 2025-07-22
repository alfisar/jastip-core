package router

import (
	simpleControll "jastip-core/application/simple/controller/http"

	"github.com/gofiber/fiber/v2"
)

type simpleRouter struct {
	Controller simpleControll.SimpleControllerContract
}

func NewSimpleRouter(Controller simpleControll.SimpleControllerContract) *simpleRouter {
	return &simpleRouter{
		Controller: Controller,
	}
}

func (r *simpleRouter) simpleRouters(v1 fiber.Router) {
	v1.Get("", r.Controller.Healthy)
	v1.Get("/grpc", r.Controller.HealthyGRPC)
}
