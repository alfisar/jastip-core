package http

import "github.com/gofiber/fiber/v2"

type TravelSchControllerContract interface {
	AddSchedule(ctx *fiber.Ctx) error
	GetListchedule(ctx *fiber.Ctx) error
	Details(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
