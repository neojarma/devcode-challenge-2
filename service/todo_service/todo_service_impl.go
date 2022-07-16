package todo_service

import (
	"devcode_challenge/model/domain"
	"devcode_challenge/model/request"
	"devcode_challenge/model/response"
	"devcode_challenge/repository/todo_repository"
	"errors"
	"strconv"
	"sync"
	"time"
)

type TodoServiceImpl struct {
	Repo todo_repository.TodoRepository
}

func NewTodoService(repo todo_repository.TodoRepository) TodoService {
	var doOnce sync.Once
	service := new(TodoServiceImpl)

	doOnce.Do(func() {
		service = &TodoServiceImpl{
			Repo: repo,
		}
	})

	return service
}

func (service *TodoServiceImpl) GetAll(activityGroupId *string) ([]*response.TodoResponse, error) {
	domainRes, err := service.Repo.GetAll(activityGroupId)
	result := make([]*response.TodoResponse, 0)
	if err != nil {
		return result, err
	}

	for _, each := range domainRes {
		todo := &response.TodoResponse{
			Id:              each.Id,
			Title:           each.Title,
			IsActive:        each.IsActive,
			Priority:        each.Priority,
			CreatedAt:       each.CreatedAt,
			UpdatedAt:       each.UpdatedAt,
			DeletedAt:       each.DeletedAt,
			ActivityGroupId: strconv.Itoa(int(each.ActivityGroupId)),
		}

		result = append(result, todo)
	}

	return result, nil
}

func (service *TodoServiceImpl) DetailTodo(todoId int) (*response.TodoResponse, error) {
	domainRes, err := service.Repo.DetailTodo(todoId)
	if err != nil {
		return nil, err
	}

	return &response.TodoResponse{
		Id:              domainRes.Id,
		Title:           domainRes.Title,
		IsActive:        domainRes.IsActive,
		Priority:        domainRes.Priority,
		CreatedAt:       domainRes.CreatedAt,
		UpdatedAt:       domainRes.UpdatedAt,
		DeletedAt:       domainRes.DeletedAt,
		ActivityGroupId: strconv.Itoa(int(domainRes.ActivityGroupId)),
	}, nil
}

func (service *TodoServiceImpl) UpdateTodo(req *request.TodoRequest) (*response.TodoResponse, error) {
	isActive := "1"
	if !req.IsActive {
		isActive = "10"
	}
	domainReq := &domain.TodoDomain{
		Title:     req.Title,
		Id:        req.Id,
		IsActive:  isActive,
		UpdatedAt: time.Now().String(),
	}
	_, err := service.Repo.UpdateTodo(domainReq)
	if err != nil {
		return nil, err
	}

	updatedData, err := service.DetailTodo(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &response.TodoResponse{
		Id:              updatedData.Id,
		Title:           updatedData.Title,
		IsActive:        updatedData.IsActive,
		Priority:        updatedData.Priority,
		CreatedAt:       updatedData.CreatedAt,
		UpdatedAt:       updatedData.UpdatedAt,
		DeletedAt:       updatedData.DeletedAt,
		ActivityGroupId: updatedData.ActivityGroupId,
	}, nil
}

func (service *TodoServiceImpl) CreateTodo(req *request.TodoRequest) (*response.TodoResponsePost, error) {
	if req.Title == "" {
		return nil, errors.New("title cannot be null")
	}

	if req.ActivityGroupId == nil {
		return nil, errors.New("activity_group_id cannot be null")
	}

	dateTime := time.Now().String()
	domainReq := &domain.TodoDomain{
		Title:           req.Title,
		ActivityGroupId: *req.ActivityGroupId,
		CreatedAt:       dateTime,
		UpdatedAt:       dateTime,
	}

	domainRes, err := service.Repo.CreateTodo(domainReq)
	if err != nil {
		return nil, err
	}

	return &response.TodoResponsePost{
		Id:              domainRes.Id,
		Title:           domainRes.Title,
		IsActive:        true,
		Priority:        "very-high",
		CreatedAt:       domainRes.CreatedAt,
		UpdatedAt:       domainRes.UpdatedAt,
		ActivityGroupId: domainReq.ActivityGroupId,
	}, nil
}

func (service *TodoServiceImpl) DeleteTodo(todoId int) error {
	err := service.Repo.DeleteTodo(todoId)
	if err != nil {
		return err
	}

	return nil
}
