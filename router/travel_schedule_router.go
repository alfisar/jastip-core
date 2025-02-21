package router

import (
	"jastip-core/application/travel_schedule/controller"

	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/alfisar/jastip-import/helpers/helper"
	"github.com/alfisar/jastip-import/helpers/middlewere"
	"github.com/gofiber/fiber/v2"
)

type travelScheduleRouter struct {
	controller controller.TravelSchControllerContract
}

func NewTravelSchRouter(controller controller.TravelSchControllerContract) *travelScheduleRouter {
	return &travelScheduleRouter{
		controller: controller,
	}
}

func (r *travelScheduleRouter) travelSchRouters(v1 fiber.Router) {
	middleweres := setMiddleware()
	v1.Post("/schedule", middleweres.Authenticate, middlewere.Validation(handler.HandlerPostSchedule, helper.ValidationPostSchedule), r.controller.AddSchedule)
	v1.Get("/schedule", middleweres.Authenticate, middlewere.Validation(handler.HandlerParamSch, nil), r.controller.GetListchedule)
}
