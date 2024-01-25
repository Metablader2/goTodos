package service

import (
	"go-todo/pkg/model"
)

type TodoRepo interface {
	Save(todo model.Todo) error
	GetPage(page int) []model.Todo
}

type todoService struct {
	todoRepo TodoRepo
}

func NewTodoService(todoRepo TodoRepo) *todoService {
	myTodoService := todoService{
		todoRepo: todoRepo,
	}
	return &myTodoService
}

func (x *todoService) AddTodo(item model.Todo) error {
	x.todoRepo.Save(item)
	return nil
}

func (x *todoService) GetTodoPage(page int) []model.Todo {

	return x.todoRepo.GetPage(page)
}
