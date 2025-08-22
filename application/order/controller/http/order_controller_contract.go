package http

import "github.com/gofiber/fiber/v2"

type OrderControllerContract interface {
	Create(ctx *fiber.Ctx) error
	GetList(ctx *fiber.Ctx) error
	Details(ctx *fiber.Ctx) error
}
