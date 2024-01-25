package repo

import (
	"context"
	"go-todo/pkg/model"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type todoRepo struct {
	pool *pgxpool.Pool
}

func NewTodoRepo(pool *pgxpool.Pool) *todoRepo {
	myTodoRepo := todoRepo{
		pool: pool,
	}
	return &myTodoRepo
}

func (x *todoRepo) Save(item model.Todo) error {

	_, err := x.pool.Exec(context.Background(), "INSERT INTO todo (title, description, duedate) VALUES ($1, $2, $3)", item.Title, item.Description, item.DueDate)
	if err != nil {
		return err
	}

	return nil
}

func (x *todoRepo) GetPage(page int) []model.Todo {

	offset := (page * 10) - 10

	rows, err := x.pool.Query(context.Background(), "SELECT title, description, duedate FROM todo LIMIT 10 OFFSET $1", offset)
	if err != nil {
		log.Fatal(err)
	}

	var todos []model.Todo

	for rows.Next() {
		var currentTodo model.Todo
		err := rows.Scan(&currentTodo.Title, &currentTodo.Description, &currentTodo.DueDate)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, currentTodo)
	}

	return todos
}
