package controller

import "github.com/gofiber/fiber/v2"

type SimpleControllerContract interface {
	Healthy(ctx *fiber.Ctx) error
}
