package usecase

import (
	"context"
	"fmt"
	"time"

	"workshop/internal/domain"
)

// PartUsageInput is one part row submitted with an intervention.
type PartUsageInput struct {
	PartName string
	Quantity int
}

// InterventionView is the intervention list read model, including the most
// recent warranty issued for the intervention when one exists.
type InterventionView struct {
	Intervention       domain.Intervention
	WarrantyID         *string
	WarrantyKind       *domain.WarrantyKind
	WarrantyExpiration *time.Time
	WarrantyValid      *bool
}

type interventionViewRepository interface {
	ListByServiceOrderView(ctx context.Context, serviceOrderID string) ([]InterventionView, error)
}

// InterventionRepository is the narrow port the intervention use case needs.
// Save writes the intervention and its parts in one transaction.
type InterventionRepository interface {
	Save(ctx context.Context, intervention domain.Intervention) error
	FindByID(ctx context.Context, id string) (domain.Intervention, error)
	ListByServiceOrder(ctx context.Context, serviceOrderID string) ([]domain.Intervention, error)
	ListByVehicle(ctx context.Context, vehicleID string) ([]domain.Intervention, error)
}

// InterventionUseCase registers the physical work executed on a vehicle.
type InterventionUseCase struct {
	intervention InterventionRepository
	order        ServiceOrderRepository
	assignment   AssignmentRepository
	technician   TechnicianRepository
	newID        func() string
	now          func() time.Time
}

// NewInterventionUseCase wires the intervention use case.
func NewInterventionUseCase(
	intervention InterventionRepository,
	order ServiceOrderRepository,
	assignment AssignmentRepository,
	technician TechnicianRepository,
	newID func() string,
	now func() time.Time,
) InterventionUseCase {
	return InterventionUseCase{
		intervention: intervention,
		order:        order,
		assignment:   assignment,
		technician:   technician,
		newID:        newID,
		now:          now,
	}
}

// Register stores an intervention with its parts and moves the order to
// IN_REPAIR the first time work is recorded. A technician who does not hold
// the order or an order that was already delivered is refused.
func (i InterventionUseCase) Register(
	ctx context.Context,
	serviceOrderID, actorUserID, description string,
	laborHourCount float64,
	part []PartUsageInput,
) (domain.Intervention, error) {
	profile, err := requireAssignedTechnician(ctx, i.assignment, i.technician, serviceOrderID, actorUserID)
	if err != nil {
		return domain.Intervention{}, err
	}
	order, err := i.order.FindByID(ctx, serviceOrderID)
	if err != nil {
		return domain.Intervention{}, err
	}
	if order.Status == domain.StatusDelivered {
		return domain.Intervention{}, fmt.Errorf("%w: the order was already delivered", domain.ErrForbidden)
	}
	performedAt := i.now()
	interventionID := i.newID()
	usage := make([]domain.PartUsage, 0, len(part))
	for _, item := range part {
		built, partErr := domain.NewPartUsage(i.newID(), interventionID, item.PartName, item.Quantity, performedAt)
		if partErr != nil {
			return domain.Intervention{}, partErr
		}
		usage = append(usage, built)
	}
	intervention, err := domain.NewIntervention(
		interventionID, serviceOrderID, profile.ID, description, laborHourCount, usage, performedAt,
	)
	if err != nil {
		return domain.Intervention{}, err
	}
	if order.Status == domain.StatusInDiagnosis {
		transition, moveErr := order.MoveTo(domain.StatusInRepair, i.newID(), actorUserID, performedAt)
		if moveErr != nil {
			return domain.Intervention{}, moveErr
		}
		if err := i.order.UpdateStatus(ctx, order, transition); err != nil {
			return domain.Intervention{}, err
		}
	}
	if err := i.intervention.Save(ctx, intervention); err != nil {
		return domain.Intervention{}, err
	}
	return intervention, nil
}

// AuthorizeRead refuses the acting user when they may not read an order.
func (i InterventionUseCase) AuthorizeRead(ctx context.Context, orderID, actorUserID string, isAdministrator bool) error {
	return OrderReadAllowed(ctx, i.assignment, i.technician, orderID, actorUserID, isAdministrator)
}

// ListByServiceOrder returns the interventions recorded on an order.
func (i InterventionUseCase) ListByServiceOrder(ctx context.Context, serviceOrderID string) ([]domain.Intervention, error) {
	return i.intervention.ListByServiceOrder(ctx, serviceOrderID)
}

// ListByServiceOrderView returns interventions with their latest warranty
// state. The fallback preserves compatibility with repositories that only
// implement the original intervention list port.
func (i InterventionUseCase) ListByServiceOrderView(ctx context.Context, serviceOrderID string) ([]InterventionView, error) {
	viewRepository, ok := i.intervention.(interventionViewRepository)
	if !ok {
		listed, err := i.intervention.ListByServiceOrder(ctx, serviceOrderID)
		if err != nil {
			return nil, err
		}
		views := make([]InterventionView, 0, len(listed))
		for _, item := range listed {
			views = append(views, InterventionView{Intervention: item})
		}
		return views, nil
	}
	views, err := viewRepository.ListByServiceOrderView(ctx, serviceOrderID)
	if err != nil {
		return nil, err
	}
	consultedAt := i.now()
	for index := range views {
		if views[index].WarrantyExpiration == nil {
			continue
		}
		valid := consultedAt.Before(*views[index].WarrantyExpiration)
		views[index].WarrantyValid = &valid
	}
	return views, nil
}
