package controller

import "github.com/gofiber/fiber/v2"

type TravelSchControllerContract interface {
	AddSchedule(ctx *fiber.Ctx) error
}
