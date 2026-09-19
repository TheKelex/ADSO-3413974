package http

import (
	"net/http"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

// partUsagePayload is one repeatable part row of the intervention form.
type partUsagePayload struct {
	PartName string `json:"partName"`
	Quantity int    `json:"quantity"`
}

// registerInterventionRequest is the payload of the intervention form.
type registerInterventionRequest struct {
	Description    string             `json:"description"`
	LaborHourCount float64            `json:"laborHourCount"`
	Part           []partUsagePayload `json:"part"`
}

// interventionResponse is one entry of the intervention list.
type interventionResponse struct {
	ID             string             `json:"id"`
	ServiceOrderID string             `json:"serviceOrderId"`
	TechnicianID   string             `json:"technicianId"`
	Description    string             `json:"description"`
	LaborHourCount float64            `json:"laborHourCount"`
	PerformedAt    string             `json:"performedAt"`
	Part           []partUsagePayload `json:"part"`
	WarrantyID     *string            `json:"warrantyId"`
	WarrantyKind   *string            `json:"kind"`
	WarrantyValid  *bool              `json:"valid"`
}

// InterventionHandler exposes the work executed on a vehicle.
type InterventionHandler struct {
	intervention usecase.InterventionUseCase
}

// NewInterventionHandler wires the intervention handler.
func NewInterventionHandler(intervention usecase.InterventionUseCase) InterventionHandler {
	return InterventionHandler{intervention: intervention}
}

// Register stores an intervention with the parts it consumed.
func (h InterventionHandler) Register(writer http.ResponseWriter, request *http.Request) {
	identity, err := callerFrom(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	var payload registerInterventionRequest
	if err := decode(writer, request, &payload); err != nil {
		failure(writer, err)
		return
	}
	part := make([]usecase.PartUsageInput, 0, len(payload.Part))
	for _, item := range payload.Part {
		part = append(part, usecase.PartUsageInput{PartName: item.PartName, Quantity: item.Quantity})
	}
	registered, err := h.intervention.Register(
		request.Context(), request.PathValue("serviceOrderId"),
		identity.UserID, payload.Description, payload.LaborHourCount, part,
	)
	if err != nil {
		failure(writer, err)
		return
	}
	respond(writer, http.StatusCreated, toInterventionResponse(usecase.InterventionView{Intervention: registered}))
}

// List returns the interventions recorded on an order.
func (h InterventionHandler) List(writer http.ResponseWriter, request *http.Request) {
	identity, err := callerFrom(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	orderID := request.PathValue("serviceOrderId")
	if err := h.intervention.AuthorizeRead(
		request.Context(), orderID, identity.UserID, identity.Role == domain.RoleAdministrator,
	); err != nil {
		failure(writer, err)
		return
	}
	listed, err := h.intervention.ListByServiceOrderView(request.Context(), orderID)
	if err != nil {
		failure(writer, err)
		return
	}
	payload := make([]interventionResponse, 0, len(listed))
	for _, item := range listed {
		payload = append(payload, toInterventionResponse(item))
	}
	respond(writer, http.StatusOK, payload)
}

func toInterventionResponse(view usecase.InterventionView) interventionResponse {
	intervention := view.Intervention
	part := make([]partUsagePayload, 0, len(intervention.Part))
	for _, item := range intervention.Part {
		part = append(part, partUsagePayload{PartName: item.PartName, Quantity: item.Quantity})
	}
	return interventionResponse{
		ID:             intervention.ID,
		ServiceOrderID: intervention.ServiceOrderID,
		TechnicianID:   intervention.TechnicianID,
		Description:    intervention.Description,
		LaborHourCount: intervention.LaborHourCount,
		PerformedAt:    formatTime(intervention.PerformedAt),
		Part:           part,
		WarrantyID:     view.WarrantyID,
		WarrantyKind:   warrantyKindValue(view.WarrantyKind),
		WarrantyValid:  view.WarrantyValid,
	}
}

func warrantyKindValue(kind *domain.WarrantyKind) *string {
	if kind == nil {
		return nil
	}
	value := string(*kind)
	return &value
}
