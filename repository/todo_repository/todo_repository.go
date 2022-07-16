package todo_repository

import "devcode_challenge/model/domain"

type TodoRepository interface {
	GetAll(activityGroupId *string) ([]*domain.TodoDomain, error)
	DetailTodo(todoId int) (*domain.TodoDomain, error)
	CreateTodo(req *domain.TodoDomain) (*domain.TodoDomain, error)
	UpdateTodo(req *domain.TodoDomain) (*domain.TodoDomain, error)
	DeleteTodo(todoId int) error
}
