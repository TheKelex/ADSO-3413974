package usecase

import (
	"context"
	"sort"
	"time"

	"workshop/internal/domain"
)

// TimelineKind labels what an entry of the clinical timeline describes.
type TimelineKind string

const (
	// TimelineOrder is the check-in of a service order.
	TimelineOrder TimelineKind = "ORDER"
	// TimelineDiagnostic is a recorded technical evaluation.
	TimelineDiagnostic TimelineKind = "DIAGNOSTIC"
	// TimelineIntervention is a physical action executed on the vehicle.
	TimelineIntervention TimelineKind = "INTERVENTION"
	// TimelineWarranty is a coverage issued over an intervention.
	TimelineWarranty TimelineKind = "WARRANTY"
)

// TimelineEntry is one dated fact in the clinical history of a vehicle.
type TimelineEntry struct {
	Kind        TimelineKind
	OccurredAt  time.Time
	Title       string
	Description string
	Reference   string
}

// Timeline is the clinical history of one vehicle.
type Timeline struct {
	Vehicle VehicleWithOwner
	Entry   []TimelineEntry
}

// TimelineUseCase assembles the chronological history of a vehicle.
type TimelineUseCase struct {
	vehicle      VehicleRepository
	order        ServiceOrderRepository
	diagnostic   DiagnosticRepository
	intervention InterventionRepository
	warranty     WarrantyRepository
}

// NewTimelineUseCase wires the clinical timeline use case.
func NewTimelineUseCase(
	vehicle VehicleRepository,
	order ServiceOrderRepository,
	diagnostic DiagnosticRepository,
	intervention InterventionRepository,
	warranty WarrantyRepository,
) TimelineUseCase {
	return TimelineUseCase{
		vehicle:      vehicle,
		order:        order,
		diagnostic:   diagnostic,
		intervention: intervention,
		warranty:     warranty,
	}
}

// Build returns every fact recorded for a vehicle, oldest first.
func (t TimelineUseCase) Build(ctx context.Context, vehicleID string) (Timeline, error) {
	vehicle, err := t.vehicle.FindByID(ctx, vehicleID)
	if err != nil {
		return Timeline{}, err
	}
	entry := make([]TimelineEntry, 0)

	order, err := t.order.ListByVehicle(ctx, vehicleID)
	if err != nil {
		return Timeline{}, err
	}
	for _, item := range order {
		entry = append(entry, TimelineEntry{
			Kind:        TimelineOrder,
			OccurredAt:  item.ReceivedAt,
			Title:       item.OrderNumber,
			Description: item.ReportedFailure,
			Reference:   item.ID,
		})
	}

	diagnostic, err := t.diagnostic.ListByVehicle(ctx, vehicleID)
	if err != nil {
		return Timeline{}, err
	}
	for _, item := range diagnostic {
		entry = append(entry, TimelineEntry{
			Kind:        TimelineDiagnostic,
			OccurredAt:  item.CreatedAt,
			Title:       item.Finding,
			Description: item.ComponentToRepair,
			Reference:   item.ServiceOrderID,
		})
	}

	intervention, err := t.intervention.ListByVehicle(ctx, vehicleID)
	if err != nil {
		return Timeline{}, err
	}
	for _, item := range intervention {
		entry = append(entry, TimelineEntry{
			Kind:        TimelineIntervention,
			OccurredAt:  item.PerformedAt,
			Title:       item.Description,
			Description: partSummary(item.Part),
			Reference:   item.ServiceOrderID,
		})
	}

	warranty, err := t.warranty.ListByVehicle(ctx, vehicleID)
	if err != nil {
		return Timeline{}, err
	}
	for _, item := range warranty {
		entry = append(entry, TimelineEntry{
			Kind:        TimelineWarranty,
			OccurredAt:  item.IssuedAt,
			Title:       string(item.Kind),
			Description: item.ExpirationDate.Format(time.RFC3339),
			Reference:   item.InterventionID,
		})
	}

	sort.SliceStable(entry, func(left, right int) bool {
		return entry[left].OccurredAt.Before(entry[right].OccurredAt)
	})
	return Timeline{Vehicle: vehicle, Entry: entry}, nil
}

func partSummary(part []domain.PartUsage) string {
	if len(part) == 0 {
		return ""
	}
	summary := ""
	for index, item := range part {
		if index > 0 {
			summary += ", "
		}
		summary += item.PartName
	}
	return summary
}
