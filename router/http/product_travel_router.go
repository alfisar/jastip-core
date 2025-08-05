package router

import (
	controller "jastip-core/application/products_travel/controller/http"

	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/alfisar/jastip-import/helpers/middlewere"
	"github.com/gofiber/fiber/v2"
)

type productsTravelRouter struct {
	controller controller.ProductsTravelControllerContract
}

func NewProductTravelRouter(controller controller.ProductsTravelControllerContract) *productsTravelRouter {
	return &productsTravelRouter{
		controller: controller,
	}
}

func (r *productsTravelRouter) productTravelRouters(v1 fiber.Router) {
	middleweres := setMiddleware()
	v1.Post("/product-travel", middleweres.Authenticate, middlewere.Validation(handler.HandlerPostProductsTravel, nil), r.controller.Create)
	v1.Delete("/product-travel", middleweres.Authenticate, middlewere.Validation(handler.HandlerDeleteProductsTravel, nil), r.controller.Delete)
}
