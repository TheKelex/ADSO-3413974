package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

type fakeUserRepository struct {
	user map[string]domain.User
}

func newFakeUserRepository(t *testing.T, username, password string) *fakeUserRepository {
	t.Helper()
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		t.Fatalf("hashing the fixture password failed: %v", err)
	}
	user, err := domain.NewUser("user-1", username, string(hash), "Administrador del taller", domain.RoleAdministrator, time.Now())
	if err != nil {
		t.Fatalf("building the user fixture failed: %v", err)
	}
	return &fakeUserRepository{user: map[string]domain.User{username: user}}
}

func (f *fakeUserRepository) FindByUsername(_ context.Context, username string) (domain.User, error) {
	found, ok := f.user[username]
	if !ok {
		return domain.User{}, domain.ErrNotFound
	}
	return found, nil
}

func (f *fakeUserRepository) FindByID(_ context.Context, id string) (domain.User, error) {
	for _, item := range f.user {
		if item.ID == id {
			return item, nil
		}
	}
	return domain.User{}, domain.ErrNotFound
}

type fakeTokenIssuer struct {
	issued int
}

func (f *fakeTokenIssuer) Issue(userID string, _ domain.Role, issuedAt time.Time) (string, time.Time, error) {
	f.issued++
	return "token-for-" + userID, issuedAt.Add(time.Hour), nil
}

func TestAuthenticateIssuesASessionForValidCredentials(t *testing.T) {
	users := newFakeUserRepository(t, "admin", "Admin2026")
	tokens := &fakeTokenIssuer{}
	useCase := usecase.NewAuthenticateUser(users, tokens, fixedClock())

	session, err := useCase.Execute(context.Background(), "admin", "Admin2026")
	if err != nil {
		t.Fatalf("valid credentials must be accepted: %v", err)
	}
	if session.Token != "token-for-user-1" {
		t.Fatalf("the session must carry the issued token, got %q", session.Token)
	}
	if session.Role != domain.RoleAdministrator {
		t.Fatalf("the session must carry the role of the user, got %s", session.Role)
	}
}

func TestAuthenticateRejectsAWrongPasswordWithoutIssuingAToken(t *testing.T) {
	users := newFakeUserRepository(t, "admin", "Admin2026")
	tokens := &fakeTokenIssuer{}
	useCase := usecase.NewAuthenticateUser(users, tokens, fixedClock())

	_, err := useCase.Execute(context.Background(), "admin", "wrong-password")
	if !errors.Is(err, domain.ErrUnauthorized) {
		t.Fatalf("a wrong password must be rejected as unauthorized, got %v", err)
	}
	if tokens.issued != 0 {
		t.Fatalf("no token may be issued for a failed sign in, got %d", tokens.issued)
	}
}

func TestAuthenticateRejectsAnUnknownUserWithTheSameError(t *testing.T) {
	users := newFakeUserRepository(t, "admin", "Admin2026")
	useCase := usecase.NewAuthenticateUser(users, &fakeTokenIssuer{}, fixedClock())

	_, unknownErr := useCase.Execute(context.Background(), "ghost", "Admin2026")
	_, wrongPasswordErr := useCase.Execute(context.Background(), "admin", "nope")
	if !errors.Is(unknownErr, domain.ErrUnauthorized) || !errors.Is(wrongPasswordErr, domain.ErrUnauthorized) {
		t.Fatalf("both failures must be unauthorized, got %v and %v", unknownErr, wrongPasswordErr)
	}
	if unknownErr.Error() != wrongPasswordErr.Error() {
		t.Fatal("an unknown user and a wrong password must fail identically, so the response reveals nothing")
	}
}

func TestAuthenticateRejectsEmptyCredentials(t *testing.T) {
	useCase := usecase.NewAuthenticateUser(newFakeUserRepository(t, "admin", "Admin2026"), &fakeTokenIssuer{}, fixedClock())
	if _, err := useCase.Execute(context.Background(), "", ""); !errors.Is(err, domain.ErrUnauthorized) {
		t.Fatalf("empty credentials must be rejected as unauthorized, got %v", err)
	}
}
