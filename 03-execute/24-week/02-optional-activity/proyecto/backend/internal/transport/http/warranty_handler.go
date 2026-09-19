package http

import (
	"net/http"
	"time"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

// issueWarrantyRequest is the payload of the warranty form.
type issueWarrantyRequest struct {
	InterventionID     string `json:"interventionId"`
	Kind               string `json:"kind"`
	CoverageMonthCount int    `json:"coverageMonthCount"`
}

// warrantyResponse is one row of the warranty table, with the badge that says
// whether it is valid at the consulted date.
type warrantyResponse struct {
	ID                 string `json:"id"`
	InterventionID     string `json:"interventionId"`
	OrderNumber        string `json:"orderNumber"`
	VehiclePlate       string `json:"vehiclePlate"`
	Kind               string `json:"kind"`
	CoverageMonthCount int    `json:"coverageMonthCount"`
	IssuedAt           string `json:"issuedAt"`
	ExpirationDate     string `json:"expirationDate"`
	Valid              bool   `json:"valid"`
}

// WarrantyHandler exposes the coverage issued over an intervention.
type WarrantyHandler struct {
	warranty usecase.WarrantyUseCase
}

// NewWarrantyHandler wires the warranty handler.
func NewWarrantyHandler(warranty usecase.WarrantyUseCase) WarrantyHandler {
	return WarrantyHandler{warranty: warranty}
}

// Issue creates a warranty over an existing intervention.
func (h WarrantyHandler) Issue(writer http.ResponseWriter, request *http.Request) {
	if _, err := requireAdministrator(request.Context()); err != nil {
		failure(writer, err)
		return
	}
	var payload issueWarrantyRequest
	if err := decode(writer, request, &payload); err != nil {
		failure(writer, err)
		return
	}
	issued, err := h.warranty.Issue(
		request.Context(), payload.InterventionID,
		domain.WarrantyKind(payload.Kind), payload.CoverageMonthCount,
	)
	if err != nil {
		failure(writer, err)
		return
	}
	respond(writer, http.StatusCreated, warrantyResponse{
		ID:                 issued.ID,
		InterventionID:     issued.InterventionID,
		Kind:               string(issued.Kind),
		CoverageMonthCount: issued.CoverageMonthCount,
		IssuedAt:           formatTime(issued.IssuedAt),
		ExpirationDate:     formatTime(issued.ExpirationDate),
		Valid:              true,
	})
}

// List returns every warranty with its validity at the consulted date. The
// date arrives as the query parameter consultedAt in RFC3339; without it the
// use case evaluates validity now.
func (h WarrantyHandler) List(writer http.ResponseWriter, request *http.Request) {
	consultedAt := time.Time{}
	if raw := request.URL.Query().Get("consultedAt"); raw != "" {
		parsed, err := time.Parse(time.RFC3339, raw)
		if err != nil {
			parsed, err = time.Parse("2006-01-02", raw)
			if err != nil {
				failure(writer, domain.ErrInvalidInput)
				return
			}
		}
		consultedAt = parsed
	}
	if _, err := requireAdministrator(request.Context()); err != nil {
		failure(writer, err)
		return
	}
	listed, err := h.warranty.List(request.Context(), consultedAt)
	if err != nil {
		failure(writer, err)
		return
	}
	payload := make([]warrantyResponse, 0, len(listed))
	for _, item := range listed {
		payload = append(payload, warrantyResponse{
			ID:                 item.Warranty.ID,
			InterventionID:     item.Warranty.InterventionID,
			OrderNumber:        item.OrderNumber,
			VehiclePlate:       item.VehiclePlate,
			Kind:               string(item.Warranty.Kind),
			CoverageMonthCount: item.Warranty.CoverageMonthCount,
			IssuedAt:           formatTime(item.Warranty.IssuedAt),
			ExpirationDate:     formatTime(item.Warranty.ExpirationDate),
			Valid:              item.Valid,
		})
	}
	respond(writer, http.StatusOK, payload)
}
