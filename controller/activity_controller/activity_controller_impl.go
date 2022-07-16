package activity_controller

import (
	"devcode_challenge/model/request"
	"devcode_challenge/model/response"
	activity_service "devcode_challenge/service/activity_service"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type ActivityControllerImpl struct {
	Service activity_service.ActivityService
}

func NewActivityController(service activity_service.ActivityService) ActivityController {
	var doOnce sync.Once
	controller := new(ActivityControllerImpl)

	doOnce.Do(func() {
		controller = &ActivityControllerImpl{
			Service: service,
		}
	})

	return controller
}

func (controller *ActivityControllerImpl) GetAll(ctx *fiber.Ctx) error {
	serviceRes, err := controller.Service.GetAll()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.NewResponse("failed", "failed", serviceRes))
	}
	return ctx.Status(http.StatusOK).JSON(response.NewResponse("Success", "Success", serviceRes))
}

func (controller *ActivityControllerImpl) DetailActivity(ctx *fiber.Ctx) error {
	activityId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.NewEmptyDataResponse("failed", "failed"))
	}

	serviceRes, err := controller.Service.DetailActivity(activityId)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(response.NewEmptyDataResponse("Not Found", fmt.Sprintf("Activity with ID %v Not Found", activityId)))
	}

	return ctx.Status(http.StatusOK).JSON(response.NewResponse("Success", "Success", serviceRes))
}

func (controller *ActivityControllerImpl) CreateActivity(ctx *fiber.Ctx) error {
	req := new(request.ActivityRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.NewEmptyDataResponse("failed", "failed"))
	}

	serviceRes, err := controller.Service.CreateActivity(req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.NewEmptyDataResponse("Bad Request", "title cannot be null"))
	}

	return ctx.Status(http.StatusCreated).JSON(response.NewResponse("Success", "Success", serviceRes))
}

func (controller *ActivityControllerImpl) UpdateActivity(ctx *fiber.Ctx) error {
	req := new(request.ActivityRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.NewEmptyDataResponse("failed", "failed"))
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.NewEmptyDataResponse("failed", "failed"))
	}

	req.Id = int64(id)
	serviceRes, err := controller.Service.UpdateActivity(req)
	if err != nil {
		if err.Error() == "not found" {
			return ctx.Status(http.StatusNotFound).JSON(response.NewEmptyDataResponse("Not Found", fmt.Sprintf("Activity with ID %v Not Found", req.Id)))
		}

		return ctx.Status(http.StatusBadRequest).JSON(response.NewEmptyDataResponse("Bad Request", "title cannot be null"))
	}

	return ctx.Status(http.StatusOK).JSON(response.NewResponse("Success", "Success", serviceRes))
}

func (controller *ActivityControllerImpl) DeleteActivity(ctx *fiber.Ctx) error {
	activityId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(response.NewEmptyDataResponse("failed", "failed"))
		}
	}

	err = controller.Service.DeleteActivity(activityId)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(response.NewEmptyDataResponse("Not Found", fmt.Sprintf("Activity with ID %v Not Found", activityId)))
	}

	return ctx.Status(http.StatusOK).JSON(response.NewEmptyDataResponse("Success", "Success"))
}
