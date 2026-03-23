package db

import (
	"auth-go/models"
	"database/sql"
)

type PermissionRepository interface {
	GetPermissionsById(id int64) (*models.Permission, error)
	GetPermissionsByName(name string) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)
	CreatePermissions(name string, description string, resource string, action string) (*models.Permission, error)
	DeletePermissionsById(id int64) error
	UpdatePermissions(id int64, name string, description string, resource string, action string) (*models.Permission, error)
}

type PermissionRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionRepository(_db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: _db,
	}
}

func (r *PermissionRepositoryImpl) GetPermissionsById(id int64) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE id = ?"
	row := r.db.QueryRow(query, id)

	permission := &models.Permission{}
	if err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt); err != nil {
		return nil, err
	}
	return permission, nil
}

func (r *PermissionRepositoryImpl) GetPermissionsByName(name string) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE name = ?"
	row := r.db.QueryRow(query, name)

	permission := &models.Permission{}
	if err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt); err != nil {
		return nil, err
	}
	return permission, nil
}

func (r *PermissionRepositoryImpl) GetAllPermissions() ([]*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions"
	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer row.Close()
	var permissions []*models.Permission
	for row.Next() {
		permission := &models.Permission{}
		if err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	if err := row.Err(); err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *PermissionRepositoryImpl) CreatePermissions(name string, description string, resource string, action string) (*models.Permission, error) {
	query := "INSERT INTO permissions (name, description, resource, action, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())"
	result, err := r.db.Exec(query, name, description, resource, action)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &models.Permission{
		Id:          id,
		Name:        name,
		Description: description,
		Resource:    resource,
		Action:      action,
		CreatedAt:   "",
		UpdatedAt:   "",
	}, nil
}

func (r *PermissionRepositoryImpl) DeletePermissionsById(id int64) error {
	query := "DELETE FROM permissions WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *PermissionRepositoryImpl) UpdatePermissions(id int64, name string, description string, resource string, action string) (*models.Permission, error) {
	query := "UPDATE permissions SET name = ?, description = ?, resource = ?, action = ?, updated_at = NOW() WHERE id = ?"
	_, err := r.db.Exec(query, name, description, resource, action, id)
	if err != nil {
		return nil, err
	}

	return &models.Permission{
		Id:          id,
		Name:        name,
		Description: description,
		Resource:    resource,
		Action:      action,
		CreatedAt:   "",
		UpdatedAt:   "",
	}, nil
}
