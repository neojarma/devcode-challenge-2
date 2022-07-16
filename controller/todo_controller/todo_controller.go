package todo_controller

import "github.com/gofiber/fiber/v2"

type TodoController interface {
	GetAll(ctx *fiber.Ctx) error
	DetailTodo(ctx *fiber.Ctx) error
	CreateTodo(ctx *fiber.Ctx) error
	UpdateTodo(ctx *fiber.Ctx) error
	DeleteTodo(ctx *fiber.Ctx) error
}
