package http

import (
	"net/http"

	"workshop/internal/usecase"
)

// signInRequest is the credential payload the login screen sends.
type signInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// sessionResponse is what a successful sign in returns. It never carries the
// password hash.
type sessionResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
	UserID    string `json:"userId"`
	Username  string `json:"username"`
	FullName  string `json:"fullName"`
	Role      string `json:"role"`
}

// AuthHandler exposes the sign in operation.
type AuthHandler struct {
	authenticate usecase.AuthenticateUser
}

// NewAuthHandler wires the authentication handler.
func NewAuthHandler(authenticate usecase.AuthenticateUser) AuthHandler {
	return AuthHandler{authenticate: authenticate}
}

// SignIn verifies the credentials and returns a session token.
func (h AuthHandler) SignIn(writer http.ResponseWriter, request *http.Request) {
	var payload signInRequest
	if err := decode(writer, request, &payload); err != nil {
		failure(writer, err)
		return
	}
	session, err := h.authenticate.Execute(request.Context(), payload.Username, payload.Password)
	if err != nil {
		failure(writer, err)
		return
	}
	respond(writer, http.StatusOK, sessionResponse{
		Token:     session.Token,
		ExpiresAt: formatTime(session.ExpiresAt),
		UserID:    session.UserID,
		Username:  session.Username,
		FullName:  session.FullName,
		Role:      string(session.Role),
	})
}
