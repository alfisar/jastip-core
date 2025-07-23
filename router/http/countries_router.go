package router

import (
	controller "jastip-core/application/countries/controller/http"

	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/alfisar/jastip-import/helpers/middlewere"
	"github.com/gofiber/fiber/v2"
)

type countriesRouter struct {
	controller controller.CountriesConrollerContract
}

func NewCountriesRouter(controller controller.CountriesConrollerContract) *countriesRouter {
	return &countriesRouter{
		controller: controller,
	}
}

func (r *countriesRouter) countriesRouters(v1 fiber.Router) {
	middleweres := setMiddleware()
	v1.Get("/countries", middleweres.Authenticate, middlewere.Validation(handler.HandlerParamCountries, nil), r.controller.GetList)

}
