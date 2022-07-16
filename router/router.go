package router

import (
	"database/sql"
	"devcode_challenge/controller/activity_controller"
	"devcode_challenge/controller/todo_controller"
	"devcode_challenge/repository/activity_repository"
	"devcode_challenge/repository/todo_repository"
	"devcode_challenge/service/activity_service"
	"devcode_challenge/service/todo_service"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func Router(fiber *fiber.App, db *sql.DB) {
	todoRoute(fiber, db)
	activityRoute(fiber, db)
}

func activityRoute(fiber *fiber.App, db *sql.DB) {
	var doOnce sync.Once
	var controller activity_controller.ActivityController

	doOnce.Do(func() {
		repo := activity_repository.NewActivityRepository(db)
		cache := cache.New(2*time.Minute, 10*time.Minute)
		service := activity_service.NewActivityService(repo, cache)
		controller = activity_controller.NewActivityController(service)
	})

	fiber.Get("/activity-groups", controller.GetAll)
	fiber.Post("/activity-groups", controller.CreateActivity)
	fiber.Get("/activity-groups/:id", controller.DetailActivity)
	fiber.Patch("/activity-groups/:id", controller.UpdateActivity)
	fiber.Delete("/activity-groups/:id", controller.DeleteActivity)
}

func todoRoute(fiber *fiber.App, db *sql.DB) {
	var doOnce sync.Once
	var controller todo_controller.TodoController

	doOnce.Do(func() {
		repo := todo_repository.NewTodoRepository(db)
		cache := cache.New(2*time.Minute, 10*time.Minute)
		service := todo_service.NewTodoService(repo, cache)
		controller = todo_controller.NewTodoController(service)
	})

	fiber.Get("/todo-items", controller.GetAll)
	fiber.Post("/todo-items", controller.CreateTodo)
	fiber.Get("/todo-items/:id", controller.DetailTodo)
	fiber.Patch("/todo-items/:id", controller.UpdateTodo)
	fiber.Delete("/todo-items/:id", controller.DeleteTodo)
}
