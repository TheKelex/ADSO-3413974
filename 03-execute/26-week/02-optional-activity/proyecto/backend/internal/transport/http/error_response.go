package http

import (
	"errors"
	"net/http"
	"strings"

	"workshop/internal/domain"
)

// errorPayload is the only error shape the API returns. The message is written
// in Spanish because the end user reads it; it never carries a driver message,
// a stack trace or an internal identifier.
type errorPayload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// failure maps a domain error to a status code and a sanitized Spanish
// message. An error the domain does not declare becomes a generic 500, so an
// unexpected internal failure never reaches the client as text.
func failure(writer http.ResponseWriter, err error) {
	status, code, message := classify(err)
	respond(writer, status, errorPayload{Code: code, Message: message})
}

// tooManyAttempts answers the login rate limiter to an identity that exhausted
// its failure budget.
func tooManyAttempts(writer http.ResponseWriter) {
	respond(writer, http.StatusTooManyRequests, errorPayload{
		Code:    "too_many_attempts",
		Message: "Demasiados intentos fallidos. Espere unos minutos antes de reintentar.",
	})
}

func classify(err error) (int, string, string) {
	switch {
	case errors.Is(err, domain.ErrUnauthorized):
		return http.StatusUnauthorized, "unauthorized", "Usuario o contrasena incorrectos."
	case errors.Is(err, domain.ErrForbidden):
		if strings.Contains(err.Error(), "already delivered") {
			return http.StatusForbidden, "forbidden", "La orden ya fue entregada y no puede modificarse."
		}
		return http.StatusForbidden, "forbidden", "No tiene permiso para realizar esta accion."
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound, "not_found", "El registro solicitado no existe."
	case errors.Is(err, domain.ErrConflict):
		return http.StatusConflict, "conflict", conflictMessage(err)
	case errors.Is(err, domain.ErrInvalidTransition):
		return http.StatusUnprocessableEntity, "invalid_transition", "Transicion de estado no permitida."
	case errors.Is(err, domain.ErrInvalidInput):
		return http.StatusBadRequest, "invalid_input", invalidInputMessage(err)
	default:
		return http.StatusInternalServerError, "internal_error", "Ocurrio un error inesperado. Intente de nuevo."
	}
}

// conflictMessage turns the few conflicts the user can act on into a precise
// Spanish sentence, and keeps a neutral one for the rest.
func conflictMessage(err error) string {
	text := err.Error()
	switch {
	case strings.Contains(text, "technician already holds"):
		return "El tecnico ya tiene una orden activa."
	case strings.Contains(text, "already has an active technician"):
		return "La orden ya tiene un tecnico asignado."
	default:
		return "El registro ya existe o entra en conflicto con otro."
	}
}

func invalidInputMessage(err error) string {
	text := err.Error()
	switch {
	case strings.Contains(text, "plate"):
		return "La placa no es valida."
	case strings.Contains(text, "VIN"):
		return "El VIN no es valido."
	case strings.Contains(text, "email"):
		return "El correo no es valido."
	case strings.Contains(text, "labor hour"):
		return "Las horas de trabajo deben ser mayores que cero."
	case strings.Contains(text, "part quantity"):
		return "La cantidad del repuesto debe ser mayor que cero."
	case strings.Contains(text, "coverage in months"):
		return "La cobertura en meses debe ser mayor que cero."
	case strings.Contains(text, "control characters"):
		return "El texto contiene caracteres no permitidos."
	case strings.Contains(text, "markup characters"):
		return "El campo contiene caracteres de marcado no permitidos."
	case strings.Contains(text, "maximum length"):
		return "El texto supera la longitud maxima."
	case strings.Contains(text, "cuerpo de la peticion"):
		return "El cuerpo de la peticion no es valido."
	default:
		return "Los datos enviados no son validos."
	}
}
