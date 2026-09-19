package usecase

import (
	"context"

	"workshop/internal/domain"
)

// StatusCount is the number of service orders sitting in one lifecycle status.
type StatusCount struct {
	Status domain.ServiceOrderStatus
	Count  int
}

// Dashboard is the operational summary the workshop manager reads.
type Dashboard struct {
	OpenOrderCount int
	StatusCount    []StatusCount
	BusyTechnician []domain.TechnicianWorkload
}

// DashboardUseCase reports the workload of the workshop.
type DashboardUseCase struct {
	order      ServiceOrderRepository
	technician TechnicianRepository
}

// NewDashboardUseCase wires the dashboard use case.
func NewDashboardUseCase(order ServiceOrderRepository, technician TechnicianRepository) DashboardUseCase {
	return DashboardUseCase{order: order, technician: technician}
}

// Build returns the open order count, the count per status in lifecycle order
// and the technicians currently holding an active order.
func (d DashboardUseCase) Build(ctx context.Context) (Dashboard, error) {
	counted, err := d.order.CountByStatus(ctx)
	if err != nil {
		return Dashboard{}, err
	}
	statusOrder := []domain.ServiceOrderStatus{
		domain.StatusReceived,
		domain.StatusInDiagnosis,
		domain.StatusInRepair,
		domain.StatusReady,
		domain.StatusDelivered,
	}
	summary := make([]StatusCount, 0, len(statusOrder))
	openCount := 0
	for _, status := range statusOrder {
		count := counted[string(status)]
		summary = append(summary, StatusCount{Status: status, Count: count})
		if status.IsOpen() {
			openCount += count
		}
	}
	workload, err := d.technician.ListWorkload(ctx)
	if err != nil {
		return Dashboard{}, err
	}
	busy := make([]domain.TechnicianWorkload, 0, len(workload))
	for _, item := range workload {
		if item.Busy {
			busy = append(busy, item)
		}
	}
	return Dashboard{OpenOrderCount: openCount, StatusCount: summary, BusyTechnician: busy}, nil
}
