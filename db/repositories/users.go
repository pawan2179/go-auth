package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRespository interface {
	Create() error
	GetById() (*models.User, error)
}

type UserRespositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRespository {
	return &UserRespositoryImpl{
		db: db,
	}
}

func (u *UserRespositoryImpl) Create() error {
	fmt.Println("in user repo for create user")
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, err := u.db.Exec(query, "testuser", "test@test.com", "password123")

	if err != nil {
		fmt.Println("Failed to create user: ", err)
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()
	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not created")
		return nil
	}

	fmt.Println("User created successfully, rows affected: ", rowsAffected)
	return nil
}

func (u *UserRespositoryImpl) GetById() (*models.User, error) {
	fmt.Println("In user repository: Get By Id")

	query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?"
	row := u.db.QueryRow(query, 1)
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with given id:", err)
			return nil, err
		} else {
			fmt.Println("Error in finding user:", err)
			return nil, err
		}
	}

	fmt.Println("User fetched successfully :", user)
	return user, nil
}
