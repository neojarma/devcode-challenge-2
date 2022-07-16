package todo_controller

import (
	"devcode_challenge/model/request"
	"devcode_challenge/model/response"
	todoservice "devcode_challenge/service/todo_service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type TodoControllerImpl struct {
	Service todoservice.TodoService
}

func NewTodoController(service todoservice.TodoService) TodoController {
	var doOnce sync.Once
	controller := new(TodoControllerImpl)

	doOnce.Do(func() {
		controller = &TodoControllerImpl{
			Service: service,
		}
	})

	return controller
}

func (controller *TodoControllerImpl) GetAll(ctx *fiber.Ctx) error {
	var serviceRes []*response.TodoResponse
	var err error

	query := ctx.Query("activity_group_id")
	if query == "" {
		serviceRes, err = controller.Service.GetAll(nil)
	} else {
		serviceRes, err = controller.Service.GetAll(&query)
	}

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.NewResponse("failed", "failed", serviceRes))
	}

	return ctx.Status(http.StatusOK).JSON(response.NewResponse("Success", "Success", serviceRes))
}

func (controller *TodoControllerImpl) DetailTodo(ctx *fiber.Ctx) error {
	todoId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.NewEmptyDataResponse("failed", "failed"))
	}

	serviceRes, err := controller.Service.DetailTodo(todoId)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(response.NewEmptyDataResponse("Not Found", fmt.Sprintf("Todo with ID %v Not Found", todoId)))
	}

	return ctx.Status(http.StatusOK).JSON(response.NewResponse("Success", "Success", serviceRes))
}

func (controller *TodoControllerImpl) CreateTodo(ctx *fiber.Ctx) error {
	req := new(request.TodoRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.NewEmptyDataResponse("failed", "failed"))
	}

	serviceRes, err := controller.Service.CreateTodo(req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.NewEmptyDataResponse("Bad Request", err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(response.NewResponse("Success", "Success", serviceRes))
}

func (controller *TodoControllerImpl) UpdateTodo(ctx *fiber.Ctx) error {
	req := new(request.TodoRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		log.Println(err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(response.NewEmptyDataResponse("failed", "failed"))
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.NewEmptyDataResponse("failed", "failed"))
	}

	req.Id = int64(id)
	serviceRes, err := controller.Service.UpdateTodo(req)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(response.NewEmptyDataResponse("Not Found", fmt.Sprintf("Todo with ID %v Not Found", req.Id)))
	}

	return ctx.Status(http.StatusOK).JSON(response.NewResponse("Success", "Success", serviceRes))
}

func (controller *TodoControllerImpl) DeleteTodo(ctx *fiber.Ctx) error {
	activityId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(response.NewEmptyDataResponse("failed", "failed"))
		}
	}

	err = controller.Service.DeleteTodo(activityId)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(response.NewEmptyDataResponse("Not Found", fmt.Sprintf("Todo with ID %v Not Found", activityId)))
	}

	return ctx.Status(http.StatusOK).JSON(response.NewEmptyDataResponse("Success", "Success"))
}
