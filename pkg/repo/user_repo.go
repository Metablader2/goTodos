package repo

import (
	"context"
	"go-todo/pkg/model"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *userRepo {
	myUserRepo := userRepo{
		pool: pool,
	}
	return &myUserRepo
}

func (u *userRepo) RegisterUser(item model.User) error {
	_, err := u.pool.Exec(context.Background(), "INSERT INTO users (email, password) VALUES ($1, $2)", item.Email, item.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) GetUser(currUser model.User) model.User {

	rows, err := u.pool.Query(context.Background(), "SELECT * FROM users WHERE email = $1 AND password = $2", currUser.Email, currUser.Password)

	if err != nil {
		log.Fatal(err)
	}

	var userFound model.User

	for rows.Next() {
		var currentUser model.User
		err := rows.Scan(&currentUser.Email, &currentUser.Password)
		if err != nil {
			log.Fatal(err)
		}

		userFound.Email = currentUser.Email
		userFound.Password = currentUser.Password
	}

	return userFound
}

func (u *userRepo) UserExists(username string, password string) bool {

	rows, err := u.pool.Query(context.Background(), "SELECT * FROM users WHERE email = $1 AND password = $2", username, password)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	count := 0

	for rows.Next() {
		count++
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if count == 1 {
		return true
	} else {
		return false
	}
}
