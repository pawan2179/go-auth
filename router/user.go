package router

import (
	"AuthInGo/controllers"
	"fmt"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		userController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	fmt.Println("In user router")
	r.Post("/signup", ur.userController.RegisterUser)
}
