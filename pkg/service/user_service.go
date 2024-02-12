package service

import "go-todo/pkg/model"

type UserRepo interface {
	RegisterUser(user model.User) error
	GetUser(currentUser model.User) model.User
	UserExists(username string, password string) bool
}

type userService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *userService {
	myUserService := userService{
		userRepo: userRepo,
	}
	return &myUserService
}

func (u *userService) AddUser(user model.User) error {
	u.userRepo.RegisterUser(user)
	return nil
}

func (u *userService) GetUserInfo(user model.User) model.User {
	return u.userRepo.GetUser(user)
}

func (u *userService) CheckUserExists(username string, password string) bool {
	return u.userRepo.UserExists(username, password)
}
