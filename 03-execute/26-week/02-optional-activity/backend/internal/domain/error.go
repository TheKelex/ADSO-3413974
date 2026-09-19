package domain

import "errors"

// Domain errors. The transport layer maps each one to a status code; no layer
// below transport knows anything about HTTP.
var (
	// ErrNotFound is returned when a requested resource does not exist.
	ErrNotFound = errors.New("resource not found")
	// ErrConflict is returned when a uniqueness rule rejects the operation.
	ErrConflict = errors.New("resource conflict")
	// ErrInvalidInput is returned when the payload violates a domain rule.
	ErrInvalidInput = errors.New("invalid input")
	// ErrUnauthorized is returned when credentials or the session token fail.
	ErrUnauthorized = errors.New("invalid credentials")
	// ErrForbidden is returned when the caller may not act on the resource.
	ErrForbidden = errors.New("operation not allowed for this user")
	// ErrInvalidTransition is returned when a status move leaves the lifecycle.
	ErrInvalidTransition = errors.New("service order status transition not allowed")
)
