package handler

import (
	"encoding/json"
	"fmt"
	"go-todo/pkg/model"
	"net/http"
	"strconv"
)

type TodoService interface {
	AddTodo(item model.Todo) error
	GetTodoPage(page int) []model.Todo
}

type todoHandler struct {
	todoService TodoService
}

func NewTodoHandler(todoService TodoService) *todoHandler {
	myTodoHandler := todoHandler{
		todoService: todoService,
	}
	return &myTodoHandler
}

func (x *todoHandler) Todo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		x.postTodo(w, r)
	case "GET":
		x.getTodo(w, r)
	default:
		http.Error(w, "unhandled request type", http.StatusMethodNotAllowed)
	}

}

func (x *todoHandler) getTodo(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		fmt.Print(err, "error parseando")
	}


	myData := x.todoService.GetTodoPage(page)
	jsonData, err := json.Marshal(myData)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)

}

func (x *todoHandler) postTodo(w http.ResponseWriter, r *http.Request) {
	var item model.Todo

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = x.todoService.AddTodo(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
