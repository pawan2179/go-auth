package service

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	utils "AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	CreateUser() error
	GetUserById() error
	LoginUser() (string, error)
}

type UserServiceImpl struct {
	userRespository db.UserRespository
}

func NewUserService(_userRepository db.UserRespository) UserService {
	return &UserServiceImpl{
		userRespository: _userRepository,
	}
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("In User service -> Create User")
	password, err := utils.HashPassword("user_password")
	if err != nil {
		return err
	}

	u.userRespository.Create(
		"username_example",
		"user_email",
		password,
	)
	return nil
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("Fetching user in UserService")
	u.userRespository.GetById()
	return nil
}

func (u *UserServiceImpl) LoginUser() (string, error) {
	email := "user_email"
	password := "user_password"
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
	payload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Failed to generate JWT token:", err)
		return "", nil
	}

	fmt.Println("Generated JWT successfully")
	return tokenString, nil
}
