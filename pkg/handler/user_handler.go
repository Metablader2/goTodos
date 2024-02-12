package handler

import (
	"go-todo/pkg/model"
	"net/http"
)

type UserService interface {
	GetUserInfo(user model.User) model.User
	AddUser(user model.User) error
	CheckUserExists(username string, password string) bool
}

type userHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *userHandler {
	myUserHandler := userHandler{
		userService: userService,
	}
	return &myUserHandler
}

func (u *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		u.postUserToDatabase(w, r)
	case "GET":
		http.Error(w, "a GET request was sent", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "unhandled request type", http.StatusMethodNotAllowed)
	}
}

func (u *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		u.findExistingUser(w, r)
	default:
		http.Error(w, "unhandled request type", http.StatusMethodNotAllowed)
	}
}

func (u *userHandler) findExistingUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	var myUser model.User

	myUser.Email = username
	myUser.Password = password

	exists := u.userService.CheckUserExists(username, password)
	w.WriteHeader(http.StatusOK)
	if exists {
		w.Write([]byte("User found"))
	} else {
		w.Write([]byte("User NOT found"))
	}

}

func (u *userHandler) postUserToDatabase(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	var myUser model.User

	myUser.Email = username
	myUser.Password = password

	err = u.userService.AddUser(myUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registered successfully"))
	// component := templates.HelloMessage()
	// component.Render(r.Context(), w)

}
