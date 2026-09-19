package http

import (
	"net/http"

	"workshop/internal/usecase"
)

// technicianResponse is one row of the technician table, with the badge that
// says whether the mechanic is available.
type technicianResponse struct {
	ID                 string `json:"id"`
	UserID             string `json:"userId"`
	FullName           string `json:"fullName"`
	Specialty          string `json:"specialty"`
	Busy               bool   `json:"busy"`
	ActiveOrderID      string `json:"activeOrderId"`
	ActiveOrderNumber  string `json:"activeOrderNumber"`
	ActiveVehiclePlate string `json:"activeVehiclePlate"`
}

// TechnicianHandler exposes the allocation panel data.
type TechnicianHandler struct {
	technician usecase.TechnicianUseCase
}

// NewTechnicianHandler wires the technician handler.
func NewTechnicianHandler(technician usecase.TechnicianUseCase) TechnicianHandler {
	return TechnicianHandler{technician: technician}
}

// List returns every technician with the order they currently hold.
func (h TechnicianHandler) List(writer http.ResponseWriter, request *http.Request) {
	if _, err := requireAdministrator(request.Context()); err != nil {
		failure(writer, err)
		return
	}
	listed, err := h.technician.ListWorkload(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	payload := make([]technicianResponse, 0, len(listed))
	for _, item := range listed {
		payload = append(payload, technicianResponse{
			ID:                 item.Technician.ID,
			UserID:             item.Technician.UserID,
			FullName:           item.FullName,
			Specialty:          item.Technician.Specialty,
			Busy:               item.Busy,
			ActiveOrderID:      item.ActiveOrderID,
			ActiveOrderNumber:  item.ActiveOrderNumber,
			ActiveVehiclePlate: item.ActiveVehiclePlate,
		})
	}
	respond(writer, http.StatusOK, payload)
}
