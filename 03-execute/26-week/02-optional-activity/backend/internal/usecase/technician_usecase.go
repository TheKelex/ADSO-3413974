package usecase

import (
	"context"

	"workshop/internal/domain"
)

// TechnicianRepository is the narrow port the allocation panel needs.
type TechnicianRepository interface {
	ListWorkload(ctx context.Context) ([]domain.TechnicianWorkload, error)
	FindByID(ctx context.Context, id string) (domain.Technician, error)
	FindByUserID(ctx context.Context, userID string) (domain.Technician, error)
}

// TechnicianUseCase reports who is available to receive a service order.
type TechnicianUseCase struct {
	technician TechnicianRepository
}

// NewTechnicianUseCase wires the technician use case.
func NewTechnicianUseCase(technician TechnicianRepository) TechnicianUseCase {
	return TechnicianUseCase{technician: technician}
}

// ListWorkload returns every technician with the order they currently hold.
func (t TechnicianUseCase) ListWorkload(ctx context.Context) ([]domain.TechnicianWorkload, error) {
	return t.technician.ListWorkload(ctx)
}

// FindByUserID resolves the mechanic profile of a signed in user.
func (t TechnicianUseCase) FindByUserID(ctx context.Context, userID string) (domain.Technician, error) {
	return t.technician.FindByUserID(ctx, userID)
}
