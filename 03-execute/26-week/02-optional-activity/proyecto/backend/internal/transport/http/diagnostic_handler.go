package http

import (
	"net/http"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

// recordDiagnosticRequest is the payload of the diagnostic panel.
type recordDiagnosticRequest struct {
	Finding           string `json:"finding"`
	ComponentToRepair string `json:"componentToRepair"`
}

// diagnosticResponse is the diagnostic shown in the order detail.
type diagnosticResponse struct {
	ID                string `json:"id"`
	ServiceOrderID    string `json:"serviceOrderId"`
	TechnicianID      string `json:"technicianId"`
	Finding           string `json:"finding"`
	ComponentToRepair string `json:"componentToRepair"`
	CreatedAt         string `json:"createdAt"`
}

// DiagnosticHandler exposes the technical evaluation of an order.
type DiagnosticHandler struct {
	diagnostic usecase.DiagnosticUseCase
}

// NewDiagnosticHandler wires the diagnostic handler.
func NewDiagnosticHandler(diagnostic usecase.DiagnosticUseCase) DiagnosticHandler {
	return DiagnosticHandler{diagnostic: diagnostic}
}

// Record stores the diagnostic written by the assigned technician.
func (h DiagnosticHandler) Record(writer http.ResponseWriter, request *http.Request) {
	identity, err := callerFrom(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	var payload recordDiagnosticRequest
	if err := decode(writer, request, &payload); err != nil {
		failure(writer, err)
		return
	}
	recorded, err := h.diagnostic.Record(
		request.Context(), request.PathValue("serviceOrderId"),
		identity.UserID, payload.Finding, payload.ComponentToRepair,
	)
	if err != nil {
		failure(writer, err)
		return
	}
	respond(writer, http.StatusCreated, toDiagnosticResponse(recorded))
}

// Find returns the diagnostic of an order.
func (h DiagnosticHandler) Find(writer http.ResponseWriter, request *http.Request) {
	identity, err := callerFrom(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	orderID := request.PathValue("serviceOrderId")
	if err := h.diagnostic.AuthorizeRead(
		request.Context(), orderID, identity.UserID, identity.Role == domain.RoleAdministrator,
	); err != nil {
		failure(writer, err)
		return
	}
	found, err := h.diagnostic.FindByServiceOrder(request.Context(), orderID)
	if err != nil {
		failure(writer, err)
		return
	}
	respond(writer, http.StatusOK, toDiagnosticResponse(found))
}

func toDiagnosticResponse(diagnostic domain.Diagnostic) diagnosticResponse {
	return diagnosticResponse{
		ID:                diagnostic.ID,
		ServiceOrderID:    diagnostic.ServiceOrderID,
		TechnicianID:      diagnostic.TechnicianID,
		Finding:           diagnostic.Finding,
		ComponentToRepair: diagnostic.ComponentToRepair,
		CreatedAt:         formatTime(diagnostic.CreatedAt),
	}
}
