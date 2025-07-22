package http

import "github.com/gofiber/fiber/v2"

type SimpleControllerContract interface {
	Healthy(ctx *fiber.Ctx) error
	HealthyGRPC(ctx *fiber.Ctx) error
}
