package router

import (
	controller "jastip-core/application/order/controller/http"

	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/alfisar/jastip-import/helpers/middlewere"
	"github.com/gofiber/fiber/v2"
)

type orderRouter struct {
	controller controller.OrderControllerContract
}

func NewOrderRouter(controller controller.OrderControllerContract) *orderRouter {
	return &orderRouter{
		controller: controller,
	}
}

func (r *orderRouter) orderRouters(v1 fiber.Router) {
	middleweres := setMiddleware()
	v1.Post("/order", middleweres.Authenticate, middlewere.Validation(handler.HandlerPostOrder, nil), r.controller.Create)
	v1.Get("/order", middleweres.Authenticate, middlewere.Validation(handler.HandlerParamOrders, nil), r.controller.GetList)
	v1.Get("/order/:id", middleweres.Authenticate, middlewere.ValidationPath(handler.HandlerPathID), r.controller.Details)
}
