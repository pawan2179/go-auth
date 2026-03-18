package db

import (
	"auth-go/models"
	"database/sql"
	"fmt"
)

type UserRespository interface {
	Create(username string, email string, hashedPassword string) (*models.User, error)
	GetById() (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
}

type UserRespositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRespository {
	return &UserRespositoryImpl{
		db: db,
	}
}

func (u *UserRespositoryImpl) Create(username string, email string, hashedPassword string) (*models.User, error) {
	fmt.Println("in user repo for create user")
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, err := u.db.Exec(query, username, email, hashedPassword)

	if err != nil {
		fmt.Println("Failed to create user: ", err)
		return nil, err
	}

	// Retrieve the inserted record ID so we can return the created user
	insertedID, idErr := result.LastInsertId()
	if idErr != nil {
		fmt.Println("Failed to get last insert id:", idErr)
		return nil, idErr
	}
	user := &models.User{}
	selectQuery := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"
	row := u.db.QueryRow(selectQuery, insertedID)
	if scanErr := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); scanErr != nil {
		fmt.Println("Failed to query created user:", scanErr)
		return nil, scanErr
	}

	fmt.Println("User created successfully:", user)
	return user, nil
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

func (u *UserRespositoryImpl) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, email, password FROM users where email = ?"
	row := u.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with given email")
			return nil, err
		} else {
			fmt.Println("Error in scanning for user:", err)
			return nil, err
		}
	}
	fmt.Println("Fetched user from email: ", user)
	return user, nil
}

func (u *UserRespositoryImpl) GetAll() ([]*models.User, error) {
	query := "SELECT id, username, email, created_at, updated_at FROM users"
	rows, err := u.db.Query(query)

	if err != nil {
		fmt.Println("Error fetching users: ", err)
		return nil, err
	}

	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			fmt.Println("Error in scanning user", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return users, nil
}

func (u *UserRespositoryImpl) DeleteById(id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	rows, err := u.db.Exec(query, id)
	if err != nil {
		fmt.Println("Failed to delete user:", err)
		return err
	}

	rowsAffected, rowErr := rows.RowsAffected()

	if rowErr != nil {
		fmt.Println("Failed to get row:", rowErr)
		return rowErr
	}
	if rowsAffected == 0 {
		fmt.Println("Failed to get affected rows:", rowErr)
		return rowErr
	}
	fmt.Println("Successfully deleted user:", rowsAffected)
	return nil
}
