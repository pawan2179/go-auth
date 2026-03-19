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
	payload := r.Context().Value("payload").(dto.CreateUserRequestDTO)

	if err := utils.ReadJsonBody(r, &payload); err != nil {
		fmt.Println("Failed to read create user request payload:", err)
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Could not read request body for create user", err)
		return
	}
	fmt.Println("Calling service to create user")
	user, err := uc.UserService.CreateUser(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User added successfully", user)
}

func (uc *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In get user by id controller")
	userId := r.URL.Query().Get("id")
	if userId == "" {
		userId = r.Context().Value("userId").(string)
	}

	if userId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User Id is required", fmt.Errorf("missing user id in request"))
		return
	}

	user, err := uc.UserService.GetUserById(userId)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
		return
	}
	if user == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "User not found", fmt.Errorf("user with id:"+userId+" is not present"))
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
	fmt.Println("User fetched successfully: ", user)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("payload").(dto.LoginUserRequestDTO)
	fmt.Println("In user login controller")

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
