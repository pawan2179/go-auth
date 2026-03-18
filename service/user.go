package service

import (
	env "auth-go/config/env"
	db "auth-go/db/repositories"
	"auth-go/dto"
	"auth-go/models"
	utils "auth-go/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error)
	GetUserById() error
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
}

type UserServiceImpl struct {
	userRespository db.UserRespository
}

func NewUserService(_userRepository db.UserRespository) UserService {
	return &UserServiceImpl{
		userRespository: _userRepository,
	}
}

func (u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error) {
	fmt.Println("In User service -> Create User")
	password, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	user, err := u.userRespository.Create(
		payload.Username,
		payload.Email,
		password,
	)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("Fetching user in UserService")
	u.userRespository.GetById()
	return nil
}

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	email := payload.Email
	password := payload.Password
	user, err := u.userRespository.GetByEmail(email)

	if err != nil {
		fmt.Println("Failed to check if user exist: ", err)
		return "", nil
	}

	if user == nil {
		fmt.Println("No user found with given email: ", email)
		return "", nil
	}

	isPasswordValid := utils.CheckPasswordHash(password, user.Password)

	if !isPasswordValid {
		fmt.Println("Password does not match")
		return "", nil
	}
	fmt.Println("User logged in successfully, generate JWT here")
	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)
	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Failed to generate JWT token:", err)
		return "", nil
	}

	fmt.Println("Generated JWT successfully")
	return tokenString, nil
}
