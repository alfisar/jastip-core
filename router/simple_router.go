package router

import (
	"jastip-core/application/simple/controller"

	"github.com/gofiber/fiber/v2"
)

type simpleRouter struct {
	Controller controller.SimpleControllerContract
}

func NewSimpleRouter(Controller controller.SimpleControllerContract) *simpleRouter {
	return &simpleRouter{
		Controller: Controller,
	}
}

func (r *simpleRouter) simpleRouters(v1 fiber.Router) {
	v1.Get("", r.Controller.Healthy)
}
