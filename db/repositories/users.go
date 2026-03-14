package db

import "fmt"

// import "database/sql"

type UserRespository interface {
	Create() error
}

type UserRespositoryImpl struct {
	// db *sql.DB
}

func NewUserRepository() UserRespository {
	return &UserRespositoryImpl{
		// db: db,
	}
}

func (u *UserRespositoryImpl) Create() error {
	fmt.Println("in user repo for create user")
	return nil
}
