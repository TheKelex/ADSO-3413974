package http

import (
	"net/http"

	"workshop/internal/usecase"
)

// statusCountResponse is one status card of the dashboard.
type statusCountResponse struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

// dashboardResponse is the operational summary the manager reads.
type dashboardResponse struct {
	OpenOrderCount int                   `json:"openOrderCount"`
	StatusCount    []statusCountResponse `json:"statusCount"`
	BusyTechnician []technicianResponse  `json:"busyTechnician"`
}

// DashboardHandler exposes the workload summary of the workshop.
type DashboardHandler struct {
	dashboard usecase.DashboardUseCase
}

// NewDashboardHandler wires the dashboard handler.
func NewDashboardHandler(dashboard usecase.DashboardUseCase) DashboardHandler {
	return DashboardHandler{dashboard: dashboard}
}

// Build returns the counts per status and the occupied technicians.
func (h DashboardHandler) Build(writer http.ResponseWriter, request *http.Request) {
	summary, err := h.dashboard.Build(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	status := make([]statusCountResponse, 0, len(summary.StatusCount))
	for _, item := range summary.StatusCount {
		status = append(status, statusCountResponse{Status: string(item.Status), Count: item.Count})
	}
	busy := make([]technicianResponse, 0, len(summary.BusyTechnician))
	for _, item := range summary.BusyTechnician {
		busy = append(busy, technicianResponse{
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
	respond(writer, http.StatusOK, dashboardResponse{
		OpenOrderCount: summary.OpenOrderCount,
		StatusCount:    status,
		BusyTechnician: busy,
	})
}
