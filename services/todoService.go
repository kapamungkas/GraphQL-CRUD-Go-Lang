package services

import (
	"graphql-todo/entities"
	"graphql-todo/repositories"
	"graphql-todo/requests"
)

type TodoService interface {
	GetByIDTodo(id int) (interface{}, error)
	GetAllTodo() (interface{}, error)

	CreateTodo(product requests.Product) (interface{}, error)
	UpdateTodo(id int, product entities.Product) (interface{}, error)
	DeleteTodo(id int) (interface{}, error)
}

type todoService struct {
	r repositories.TodoRepository
}

func NewTodoService(r repositories.TodoRepository) *todoService {
	return &todoService{r}
}

func (s todoService) GetByIDTodo(id int) (interface{}, error) {
	return s.r.GetByIDTodo(id)
}

func (s todoService) GetAllTodo() (interface{}, error) {
	return s.r.GetAllTodo()
}

func (s todoService) CreateTodo(product entities.Product) (interface{}, error) {
	return s.r.CreateTodo(product)
}

func (s todoService) UpdateTodo(id int, product entities.Product) (interface{}, error) {
	_, err := s.r.UpdateTodo(id, product)
	if err == nil {
		return s.r.GetByIDTodo(id)
	}
	return nil, err
}

func (s todoService) DeleteTodo(id int) (interface{}, error) {
	return s.r.DeleteTodo(id)
}
