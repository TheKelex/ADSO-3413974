package http

import (
	"net/http"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

// assignTechnicianRequest is the payload of the allocation panel.
type assignTechnicianRequest struct {
	TechnicianID string `json:"technicianId"`
}

// assignmentResponse describes who holds a service order.
type assignmentResponse struct {
	ID             string `json:"id"`
	ServiceOrderID string `json:"serviceOrderId"`
	TechnicianID   string `json:"technicianId"`
	IsActive       bool   `json:"isActive"`
	AssignedAt     string `json:"assignedAt"`
}

// AssignmentHandler exposes the allocation of a technician to an order.
type AssignmentHandler struct {
	assignment usecase.AssignmentUseCase
}

// NewAssignmentHandler wires the assignment handler.
func NewAssignmentHandler(assignment usecase.AssignmentUseCase) AssignmentHandler {
	return AssignmentHandler{assignment: assignment}
}

// Assign gives the order to a technician.
func (h AssignmentHandler) Assign(writer http.ResponseWriter, request *http.Request) {
	if _, err := requireAdministrator(request.Context()); err != nil {
		failure(writer, err)
		return
	}
	var payload assignTechnicianRequest
	if err := decode(writer, request, &payload); err != nil {
		failure(writer, err)
		return
	}
	assignment, err := h.assignment.Assign(
		request.Context(), request.PathValue("serviceOrderId"), payload.TechnicianID,
	)
	if err != nil {
		failure(writer, err)
		return
	}
	respond(writer, http.StatusCreated, toAssignmentResponse(assignment))
}

// Find returns the technician currently holding the order.
func (h AssignmentHandler) Find(writer http.ResponseWriter, request *http.Request) {
	identity, err := callerFrom(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	orderID := request.PathValue("serviceOrderId")
	if err := h.assignment.AuthorizeRead(
		request.Context(), orderID, identity.UserID, identity.Role == domain.RoleAdministrator,
	); err != nil {
		failure(writer, err)
		return
	}
	assignment, err := h.assignment.FindActive(request.Context(), orderID)
	if err != nil {
		failure(writer, err)
		return
	}
	respond(writer, http.StatusOK, toAssignmentResponse(assignment))
}

func toAssignmentResponse(assignment domain.Assignment) assignmentResponse {
	return assignmentResponse{
		ID:             assignment.ID,
		ServiceOrderID: assignment.ServiceOrderID,
		TechnicianID:   assignment.TechnicianID,
		IsActive:       assignment.IsActive,
		AssignedAt:     formatTime(assignment.AssignedAt),
	}
}
