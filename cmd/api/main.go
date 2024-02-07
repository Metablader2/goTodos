package main

import (
	"context"
	"fmt"
	"go-todo/pkg/handler"
	"go-todo/pkg/repo"
	"go-todo/pkg/service"
	"go-todo/templates"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	component := templates.Hello("Juan")
	component.Render(r.Context(),w)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file juan")
	}

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	todoRepo := repo.NewTodoRepo(pool)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	http.HandleFunc("/", getIndex)
	http.HandleFunc("/todo", todoHandler.Todo)
	err = http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}
}
