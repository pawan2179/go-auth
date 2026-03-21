-- +goose Up
CREATE TABLE IF NOT EXISTS role_permissions (
  id SERIAL PRIMARY KEY,
  role_id BIGINT UNSIGNED NOT NULL,
  permission_id BIGINT UNSIGNED NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
  FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);

-- INSERT INTO role_permissions (role_id, permission_id)
-- SELECT 1, id FROM permissions; -- admin has all permission

-- INSERT INTO role_permissions (role_id, permission_id) -- all users have read permission to other users
-- SELECT 2, id FROM permissions where name IN ('user:read');

-- +goose Down
SELECT 'down SQL query';
