package controllers

import (
	"auth-go/dto"
	services "auth-go/service"
	"auth-go/utils"
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

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In user login controller")

	var payload dto.LoginUserRequestDTO

	if err := utils.ReadJsonBody(r, &payload); err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went wrong while logging in", err)
		return
	}

	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", validationErr)
	}

	jwtToken, tokenErr := uc.UserService.LoginUser(&payload)
	if tokenErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Faield to login user", tokenErr)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}
