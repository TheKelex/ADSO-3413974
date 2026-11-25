package usecase

import (
	"context"
	"time"

	"workshop/internal/domain"
)

// CustomerRepository is the narrow port the customer use case needs.
type CustomerRepository interface {
	Save(ctx context.Context, customer domain.Customer) error
	List(ctx context.Context) ([]domain.Customer, error)
	FindByID(ctx context.Context, id string) (domain.Customer, error)
}

// CustomerUseCase registers and lists the owners of the serviced vehicles.
type CustomerUseCase struct {
	customer CustomerRepository
	newID    func() string
	now      func() time.Time
}

// NewCustomerUseCase wires the customer use case.
func NewCustomerUseCase(customer CustomerRepository, newID func() string, now func() time.Time) CustomerUseCase {
	return CustomerUseCase{customer: customer, newID: newID, now: now}
}

// Create registers a customer with its contact data.
func (c CustomerUseCase) Create(ctx context.Context, fullName, documentNumber, phone, email string) (domain.Customer, error) {
	customer, err := domain.NewCustomer(c.newID(), fullName, documentNumber, phone, email, c.now())
	if err != nil {
		return domain.Customer{}, err
	}
	if err := c.customer.Save(ctx, customer); err != nil {
		return domain.Customer{}, err
	}
	return customer, nil
}

// List returns every registered customer.
func (c CustomerUseCase) List(ctx context.Context) ([]domain.Customer, error) {
	return c.customer.List(ctx)
}
