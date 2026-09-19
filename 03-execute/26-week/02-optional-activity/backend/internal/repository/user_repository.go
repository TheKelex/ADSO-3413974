package repository

import (
	"context"
	"database/sql"
	"time"

	"workshop/internal/domain"
)

const userColumn = "id, username, password_hash, role, full_name, created_at"

// UserRepository reads user accounts for authentication.
type UserRepository struct {
	database *sql.DB
	timeout  time.Duration
}

// NewUserRepository wires the user adapter.
func NewUserRepository(database *sql.DB, timeout time.Duration) UserRepository {
	return UserRepository{database: database, timeout: timeout}
}

// FindByUsername reads the account that signs in with this username.
func (r UserRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	return r.findBy(ctx, "SELECT "+userColumn+" FROM `user` WHERE username = ?", username)
}

// FindByID reads the account behind a session token.
func (r UserRepository) FindByID(ctx context.Context, id string) (domain.User, error) {
	return r.findBy(ctx, "SELECT "+userColumn+" FROM `user` WHERE id = ?", id)
}

func (r UserRepository) findBy(ctx context.Context, query string, argument any) (domain.User, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var user domain.User
	var role string
	err := r.database.QueryRowContext(queryCtx, query, argument).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &role, &user.FullName, &user.CreatedAt,
	)
	if err != nil {
		return domain.User{}, translate(err)
	}
	user.Role = domain.Role(role)
	return user, nil
}
