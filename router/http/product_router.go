package router

import (
	controller "jastip-core/application/products/controller/http"

	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/alfisar/jastip-import/helpers/helper"
	"github.com/alfisar/jastip-import/helpers/middlewere"
	"github.com/gofiber/fiber/v2"
)

type productsRouter struct {
	controller controller.ProductContollerContract
}

func NewProductRouter(controller controller.ProductContollerContract) *productsRouter {
	return &productsRouter{
		controller: controller,
	}
}

func (r *productsRouter) productRouters(v1 fiber.Router) {
	middleweres := setMiddleware()
	v1.Post("/product", middleweres.Authenticate, middlewere.Validation(handler.HandlerPostProducts, helper.ValidationPostProducts), r.controller.Create)
	v1.Get("/product", middleweres.Authenticate, middlewere.Validation(handler.HandlerParamProducts, nil), r.controller.GetList)
	v1.Get("/product/travel/:id", middleweres.Authenticate, middlewere.Validation(handler.HandlerParamProducts, nil), middlewere.ValidationPath(handler.HandlerPathID), r.controller.GetListProductTravel)
	v1.Patch("/product/:id", middleweres.Authenticate, middlewere.Validation(handler.HandlerUpdateProducts, helper.ValidationUpdateProduct), middlewere.ValidationPath(handler.HandlerPathID), r.controller.Update)
	v1.Delete("/product/:id", middleweres.Authenticate, middlewere.ValidationPath(handler.HandlerPathID), r.controller.Delete)
}
