package router

import (
	"auth-go/controllers"
	"auth-go/middlewares"
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
	r.With(middlewares.CreateUserRequestValidator).Post("/signup", ur.userController.RegisterUser)
	r.With(middlewares.JWTAuthMiddleware).Get("/profile", ur.userController.GetById)
	r.With(middlewares.LoginUserRequestValidator).Post("/login", ur.userController.LoginUser)
}
