package usecase

import (
	"context"
	"time"

	"workshop/internal/domain"
)

// ServiceOrderSummary is the read model of the order list: the order plus the
// plate of its vehicle and the name of the technician holding it.
type ServiceOrderSummary struct {
	Order          domain.ServiceOrder
	VehiclePlate   string
	TechnicianName string
}

// ServiceOrderRepository is the narrow port the service order use case needs.
// UpdateStatus writes the order and its transition record in one transaction.
type ServiceOrderRepository interface {
	Save(ctx context.Context, order domain.ServiceOrder) error
	FindByID(ctx context.Context, id string) (domain.ServiceOrder, error)
	FindSummary(ctx context.Context, id string) (ServiceOrderSummary, error)
	List(ctx context.Context, status string) ([]ServiceOrderSummary, error)
	ListByTechnician(ctx context.Context, technicianID, status string) ([]ServiceOrderSummary, error)
	ListByVehicle(ctx context.Context, vehicleID string) ([]domain.ServiceOrder, error)
	UpdateStatus(ctx context.Context, order domain.ServiceOrder, transition domain.StatusTransition) error
	ListTransition(ctx context.Context, serviceOrderID string) ([]domain.StatusTransition, error)
	CountByStatus(ctx context.Context) (map[string]int, error)
	NextOrderNumber(ctx context.Context) (string, error)
}

// ServiceOrderUseCase opens orders at check-in and advances their lifecycle.
type ServiceOrderUseCase struct {
	order      ServiceOrderRepository
	vehicle    VehicleRepository
	assignment AssignmentRepository
	technician TechnicianRepository
	newID      func() string
	now        func() time.Time
}

// NewServiceOrderUseCase wires the service order use case.
func NewServiceOrderUseCase(
	order ServiceOrderRepository,
	vehicle VehicleRepository,
	assignment AssignmentRepository,
	technician TechnicianRepository,
	newID func() string,
	now func() time.Time,
) ServiceOrderUseCase {
	return ServiceOrderUseCase{
		order:      order,
		vehicle:    vehicle,
		assignment: assignment,
		technician: technician,
		newID:      newID,
		now:        now,
	}
}

// Open registers the check-in of a vehicle and returns the created order.
func (s ServiceOrderUseCase) Open(ctx context.Context, vehicleID, reportedFailure string) (domain.ServiceOrder, error) {
	if _, err := s.vehicle.FindByID(ctx, vehicleID); err != nil {
		return domain.ServiceOrder{}, err
	}
	orderNumber, err := s.order.NextOrderNumber(ctx)
	if err != nil {
		return domain.ServiceOrder{}, err
	}
	order, err := domain.NewServiceOrder(s.newID(), orderNumber, vehicleID, reportedFailure, s.now())
	if err != nil {
		return domain.ServiceOrder{}, err
	}
	if err := s.order.Save(ctx, order); err != nil {
		return domain.ServiceOrder{}, err
	}
	return order, nil
}

// AuthorizeRead refuses the acting user when they may not read an order. The
// administrator may read every order; a technician only the ones assigned.
func (s ServiceOrderUseCase) AuthorizeRead(ctx context.Context, orderID, actorUserID string, isAdministrator bool) error {
	return OrderReadAllowed(ctx, s.assignment, s.technician, orderID, actorUserID, isAdministrator)
}

// List returns the orders the acting user may see, optionally filtered by a
// lifecycle status. Administrators see every order; a technician only the ones
// the technician currently holds.
func (s ServiceOrderUseCase) List(ctx context.Context, status, actorUserID string, isAdministrator bool) ([]ServiceOrderSummary, error) {
	if isAdministrator {
		return s.order.List(ctx, status)
	}
	profile, err := s.technician.FindByUserID(ctx, actorUserID)
	if err != nil {
		return nil, domain.ErrForbidden
	}
	return s.order.ListByTechnician(ctx, profile.ID, status)
}

// Find returns one order by its identifier.
func (s ServiceOrderUseCase) Find(ctx context.Context, orderID string) (domain.ServiceOrder, error) {
	return s.order.FindByID(ctx, orderID)
}

// FindSummary returns an order with the vehicle plate and active technician.
func (s ServiceOrderUseCase) FindSummary(ctx context.Context, orderID string) (ServiceOrderSummary, error) {
	return s.order.FindSummary(ctx, orderID)
}

// ListTransition returns the status history of an order.
func (s ServiceOrderUseCase) ListTransition(ctx context.Context, orderID string) ([]domain.StatusTransition, error) {
	return s.order.ListTransition(ctx, orderID)
}

// Advance moves an order to the next status. Only the administrator or the
// technician holding the order can move it; the domain rejects a move outside
// the lifecycle before anything is written, so the stored order is untouched.
// Reaching DELIVERED releases the technician who held the order.
func (s ServiceOrderUseCase) Advance(ctx context.Context, orderID string, next domain.ServiceOrderStatus, actorUserID string, isAdministrator bool) (domain.ServiceOrder, error) {
	if err := OrderReadAllowed(ctx, s.assignment, s.technician, orderID, actorUserID, isAdministrator); err != nil {
		return domain.ServiceOrder{}, err
	}
	order, err := s.order.FindByID(ctx, orderID)
	if err != nil {
		return domain.ServiceOrder{}, err
	}
	changedAt := s.now()
	transition, err := order.MoveTo(next, s.newID(), actorUserID, changedAt)
	if err != nil {
		return domain.ServiceOrder{}, err
	}
	if err := s.order.UpdateStatus(ctx, order, transition); err != nil {
		return domain.ServiceOrder{}, err
	}
	if next == domain.StatusDelivered {
		if err := s.assignment.ReleaseByServiceOrder(ctx, order.ID, changedAt); err != nil {
			return domain.ServiceOrder{}, err
		}
	}
	return order, nil
}
