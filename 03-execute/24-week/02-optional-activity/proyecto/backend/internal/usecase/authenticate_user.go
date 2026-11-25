package usecase

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"workshop/internal/domain"
)

// UserRepository is the narrow port the authentication use case needs.
type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (domain.User, error)
	FindByID(ctx context.Context, id string) (domain.User, error)
}

// TokenIssuer signs the session token the client sends back on every write.
type TokenIssuer interface {
	Issue(userID string, role domain.Role, issuedAt time.Time) (string, time.Time, error)
}

// Session is what a successful sign in returns to the caller.
type Session struct {
	Token     string
	ExpiresAt time.Time
	UserID    string
	Username  string
	FullName  string
	Role      domain.Role
}

// AuthenticateUser verifies credentials and issues a session token.
type AuthenticateUser struct {
	user  UserRepository
	token TokenIssuer
	now   func() time.Time
}

// NewAuthenticateUser wires the authentication use case.
func NewAuthenticateUser(user UserRepository, token TokenIssuer, now func() time.Time) AuthenticateUser {
	return AuthenticateUser{user: user, token: token, now: now}
}

// Execute returns a session when the credentials match the stored hash. A
// wrong username and a wrong password return the same error on purpose, so the
// response never says which half of the credential failed.
func (a AuthenticateUser) Execute(ctx context.Context, username, password string) (Session, error) {
	if username == "" || password == "" {
		return Session{}, fmt.Errorf("%w: username and password are required", domain.ErrUnauthorized)
	}
	found, err := a.user.FindByUsername(ctx, username)
	if err != nil {
		return Session{}, domain.ErrUnauthorized
	}
	if bcrypt.CompareHashAndPassword([]byte(found.PasswordHash), []byte(password)) != nil {
		return Session{}, domain.ErrUnauthorized
	}
	issuedAt := a.now()
	token, expiresAt, err := a.token.Issue(found.ID, found.Role, issuedAt)
	if err != nil {
		return Session{}, err
	}
	return Session{
		Token:     token,
		ExpiresAt: expiresAt,
		UserID:    found.ID,
		Username:  found.Username,
		FullName:  found.FullName,
		Role:      found.Role,
	}, nil
}
