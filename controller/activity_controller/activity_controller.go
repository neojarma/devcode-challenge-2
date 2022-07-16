package activity_controller

import "github.com/gofiber/fiber/v2"

type ActivityController interface {
	GetAll(ctx *fiber.Ctx) error
	DetailActivity(ctx *fiber.Ctx) error
	CreateActivity(ctx *fiber.Ctx) error
	UpdateActivity(ctx *fiber.Ctx) error
	DeleteActivity(ctx *fiber.Ctx) error
}
