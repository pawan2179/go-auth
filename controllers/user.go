package controllers

import (
	services "AuthInGo/service"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in user controller for /register")
	uc.UserService.CreateUser()
}

func (uc *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In get user by id controller")
	uc.UserService.GetUserById()
	w.Write([]byte("User fetching endpoint"))
}
