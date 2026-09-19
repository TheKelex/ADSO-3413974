package http

import (
	"net/http"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

// createServiceOrderRequest is the payload of the check-in form.
type createServiceOrderRequest struct {
	VehicleID       string `json:"vehicleId"`
	ReportedFailure string `json:"reportedFailure"`
}

// advanceStatusRequest is the payload of the status advance buttons.
type advanceStatusRequest struct {
	Status string `json:"status"`
}

// serviceOrderResponse is one row of the order table and the order detail.
type serviceOrderResponse struct {
	ID              string `json:"id"`
	OrderNumber     string `json:"orderNumber"`
	VehicleID       string `json:"vehicleId"`
	VehiclePlate    string `json:"vehiclePlate"`
	TechnicianName  string `json:"technicianName"`
	ReportedFailure string `json:"reportedFailure"`
	Status          string `json:"status"`
	ReceivedAt      string `json:"receivedAt"`
	UpdatedAt       string `json:"updatedAt"`
}

// statusTransitionResponse is one row of the status history panel.
type statusTransitionResponse struct {
	ID              string `json:"id"`
	FromStatus      string `json:"fromStatus"`
	ToStatus        string `json:"toStatus"`
	ChangedByUserID string `json:"changedByUserId"`
	ChangedAt       string `json:"changedAt"`
}

// ServiceOrderHandler exposes the check-in and the lifecycle of an order.
type ServiceOrderHandler struct {
	order usecase.ServiceOrderUseCase
}

// NewServiceOrderHandler wires the service order handler.
func NewServiceOrderHandler(order usecase.ServiceOrderUseCase) ServiceOrderHandler {
	return ServiceOrderHandler{order: order}
}

// Create opens a service order at check-in.
func (h ServiceOrderHandler) Create(writer http.ResponseWriter, request *http.Request) {
	if _, err := requireAdministrator(request.Context()); err != nil {
		failure(writer, err)
		return
	}
	var payload createServiceOrderRequest
	if err := decode(writer, request, &payload); err != nil {
		failure(writer, err)
		return
	}
	created, err := h.order.Open(request.Context(), payload.VehicleID, payload.ReportedFailure)
	if err != nil {
		failure(writer, err)
		return
	}
	respond(writer, http.StatusCreated, toServiceOrderResponse(created, "", ""))
}

// List returns the orders the acting user may see, optionally filtered by
// status. Administrators see every order; a technician only the ones held.
func (h ServiceOrderHandler) List(writer http.ResponseWriter, request *http.Request) {
	identity, err := callerFrom(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	listed, err := h.order.List(
		request.Context(), request.URL.Query().Get("status"), identity.UserID, identity.Role == domain.RoleAdministrator,
	)
	if err != nil {
		failure(writer, err)
		return
	}
	payload := make([]serviceOrderResponse, 0, len(listed))
	for _, item := range listed {
		payload = append(payload, toServiceOrderResponse(item.Order, item.VehiclePlate, item.TechnicianName))
	}
	respond(writer, http.StatusOK, payload)
}

// Find returns one order by its identifier.
func (h ServiceOrderHandler) Find(writer http.ResponseWriter, request *http.Request) {
	identity, err := callerFrom(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	orderID := request.PathValue("serviceOrderId")
	if err := h.order.AuthorizeRead(
		request.Context(), orderID, identity.UserID, identity.Role == domain.RoleAdministrator,
	); err != nil {
		failure(writer, err)
		return
	}
	summary, err := h.order.FindSummary(request.Context(), orderID)
	if err != nil {
		failure(writer, err)
		return
	}
	respond(writer, http.StatusOK, toServiceOrderResponse(summary.Order, summary.VehiclePlate, summary.TechnicianName))
}

// ListTransition returns the status history of an order.
func (h ServiceOrderHandler) ListTransition(writer http.ResponseWriter, request *http.Request) {
	identity, err := callerFrom(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	orderID := request.PathValue("serviceOrderId")
	if err := h.order.AuthorizeRead(
		request.Context(), orderID, identity.UserID, identity.Role == domain.RoleAdministrator,
	); err != nil {
		failure(writer, err)
		return
	}
	history, err := h.order.ListTransition(request.Context(), orderID)
	if err != nil {
		failure(writer, err)
		return
	}
	payload := make([]statusTransitionResponse, 0, len(history))
	for _, item := range history {
		payload = append(payload, statusTransitionResponse{
			ID:              item.ID,
			FromStatus:      string(item.FromStatus),
			ToStatus:        string(item.ToStatus),
			ChangedByUserID: item.ChangedByUserID,
			ChangedAt:       formatTime(item.ChangedAt),
		})
	}
	respond(writer, http.StatusOK, payload)
}

// Advance moves an order to the next lifecycle status.
func (h ServiceOrderHandler) Advance(writer http.ResponseWriter, request *http.Request) {
	identity, err := callerFrom(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	var payload advanceStatusRequest
	if err := decode(writer, request, &payload); err != nil {
		failure(writer, err)
		return
	}
	order, err := h.order.Advance(
		request.Context(),
		request.PathValue("serviceOrderId"),
		domain.ServiceOrderStatus(payload.Status),
		identity.UserID,
		identity.Role == domain.RoleAdministrator,
	)
	if err != nil {
		failure(writer, err)
		return
	}
	respond(writer, http.StatusOK, toServiceOrderResponse(order, "", ""))
}

func toServiceOrderResponse(order domain.ServiceOrder, plate, technicianName string) serviceOrderResponse {
	return serviceOrderResponse{
		ID:              order.ID,
		OrderNumber:     order.OrderNumber,
		VehicleID:       order.VehicleID,
		VehiclePlate:    plate,
		TechnicianName:  technicianName,
		ReportedFailure: order.ReportedFailure,
		Status:          string(order.Status),
		ReceivedAt:      formatTime(order.ReceivedAt),
		UpdatedAt:       formatTime(order.UpdatedAt),
	}
}
