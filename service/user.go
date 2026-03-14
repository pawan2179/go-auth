package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type UserService interface {
	CreateUser() error
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
	u.userRespository.Create()
	return nil
}
