package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"workshop/internal/domain"
)

// AssignmentRepository is the narrow port the allocation use case needs.
type AssignmentRepository interface {
	Save(ctx context.Context, assignment domain.Assignment) error
	FindActiveByServiceOrder(ctx context.Context, serviceOrderID string) (domain.Assignment, error)
	FindActiveByTechnician(ctx context.Context, technicianID string) (domain.Assignment, error)
	ReleaseByServiceOrder(ctx context.Context, serviceOrderID string, releasedAt time.Time) error
}

// AssignmentUseCase binds a technician to a service order.
type AssignmentUseCase struct {
	assignment AssignmentRepository
	order      ServiceOrderRepository
	technician TechnicianRepository
	newID      func() string
	now        func() time.Time
}

// NewAssignmentUseCase wires the allocation use case.
func NewAssignmentUseCase(
	assignment AssignmentRepository,
	order ServiceOrderRepository,
	technician TechnicianRepository,
	newID func() string,
	now func() time.Time,
) AssignmentUseCase {
	return AssignmentUseCase{assignment: assignment, order: order, technician: technician, newID: newID, now: now}
}

// Assign gives a service order to a technician. It is rejected when that
// technician still holds an order that has not been delivered. The check here
// gives the user a clear message; the unique active marker in storage is what
// makes the rule hold under a concurrent retry.
func (a AssignmentUseCase) Assign(ctx context.Context, serviceOrderID, technicianID string) (domain.Assignment, error) {
	if _, err := a.order.FindByID(ctx, serviceOrderID); err != nil {
		return domain.Assignment{}, err
	}
	if _, err := a.technician.FindByID(ctx, technicianID); err != nil {
		return domain.Assignment{}, err
	}
	if _, err := a.assignment.FindActiveByTechnician(ctx, technicianID); err == nil {
		return domain.Assignment{}, fmt.Errorf("%w: the technician already holds an active service order", domain.ErrConflict)
	} else if !errors.Is(err, domain.ErrNotFound) {
		return domain.Assignment{}, err
	}
	existing, err := a.assignment.FindActiveByServiceOrder(ctx, serviceOrderID)
	if err == nil {
		if existing.TechnicianID == technicianID {
			return existing, nil
		}
		return domain.Assignment{}, fmt.Errorf("%w: the service order already has an active technician", domain.ErrConflict)
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return domain.Assignment{}, err
	}
	assignment, err := domain.NewAssignment(a.newID(), serviceOrderID, technicianID, a.now())
	if err != nil {
		return domain.Assignment{}, err
	}
	if err := a.assignment.Save(ctx, assignment); err != nil {
		return domain.Assignment{}, err
	}
	return assignment, nil
}

// AuthorizeRead refuses the acting user when they may not read an order.
func (a AssignmentUseCase) AuthorizeRead(ctx context.Context, orderID, actorUserID string, isAdministrator bool) error {
	return OrderReadAllowed(ctx, a.assignment, a.technician, orderID, actorUserID, isAdministrator)
}

// FindActive returns the technician currently holding a service order.
func (a AssignmentUseCase) FindActive(ctx context.Context, serviceOrderID string) (domain.Assignment, error) {
	return a.assignment.FindActiveByServiceOrder(ctx, serviceOrderID)
}

// requireAssignedTechnician resolves the mechanic profile of the acting user
// and refuses when that technician does not hold the order. It is the single
// place where the ownership rule of a write on a service order lives.
func requireAssignedTechnician(
	ctx context.Context,
	assignment AssignmentRepository,
	technician TechnicianRepository,
	serviceOrderID, actorUserID string,
) (domain.Technician, error) {
	profile, err := technician.FindByUserID(ctx, actorUserID)
	if err != nil {
		return domain.Technician{}, fmt.Errorf("%w: only a technician can write on a service order", domain.ErrForbidden)
	}
	active, err := assignment.FindActiveByServiceOrder(ctx, serviceOrderID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return domain.Technician{}, fmt.Errorf("%w: the service order has no assigned technician", domain.ErrForbidden)
		}
		return domain.Technician{}, err
	}
	if active.TechnicianID != profile.ID {
		return domain.Technician{}, fmt.Errorf("%w: only the assigned technician can write on this service order", domain.ErrForbidden)
	}
	return profile, nil
}

// OrderReadAllowed refuses the acting user when they may not access a service
// order. The administrator may access every order; a technician only the ones
// assigned to them. It is the shared rule behind the order reads and the
// status advance.
func OrderReadAllowed(
	ctx context.Context,
	assignment AssignmentRepository,
	technician TechnicianRepository,
	serviceOrderID, actorUserID string,
	isAdministrator bool,
) error {
	if isAdministrator {
		return nil
	}
	_, err := requireAssignedTechnician(ctx, assignment, technician, serviceOrderID, actorUserID)
	return err
}
