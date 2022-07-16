package todo_service

import (
	"devcode_challenge/model/request"
	"devcode_challenge/model/response"
)

type TodoService interface {
	GetAll(activityGroupId *string) ([]*response.TodoResponse, error)
	DetailTodo(todoId int) (*response.TodoResponse, error)
	UpdateTodo(req *request.TodoRequest) (*response.TodoResponse, error)
	CreateTodo(req *request.TodoRequest) (*response.TodoResponsePost, error)
	DeleteTodo(todoId int) error
}
