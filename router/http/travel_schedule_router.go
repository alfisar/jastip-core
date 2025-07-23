package router

import (
	controller "jastip-core/application/travel_schedule/controller/http"

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
	v1.Get("/schedule/:id", middleweres.Authenticate, middlewere.ValidationPath(handler.HandlerPathID), r.controller.Details)
	v1.Patch("/schedule/:id", middleweres.Authenticate, middlewere.Validation(handler.HandlerUpdate, helper.ValidationUpdateTravelSch), middlewere.ValidationPath(handler.HandlerPathID), r.controller.Update)
	v1.Delete("/schedule/:id", middleweres.Authenticate, middlewere.ValidationPath(handler.HandlerPathID), r.controller.Delete)
}
