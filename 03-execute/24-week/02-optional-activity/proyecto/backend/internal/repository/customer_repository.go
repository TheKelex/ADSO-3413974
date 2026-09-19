package repository

import (
	"context"
	"database/sql"
	"time"

	"workshop/internal/domain"
)

const customerColumn = "id, full_name, document_number, phone, email, created_at"

// CustomerRepository persists and reads the owners of the serviced vehicles.
type CustomerRepository struct {
	database *sql.DB
	timeout  time.Duration
}

// NewCustomerRepository wires the customer adapter.
func NewCustomerRepository(database *sql.DB, timeout time.Duration) CustomerRepository {
	return CustomerRepository{database: database, timeout: timeout}
}

// Save stores a customer.
func (r CustomerRepository) Save(ctx context.Context, customer domain.Customer) error {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.database.ExecContext(
		queryCtx,
		"INSERT INTO customer ("+customerColumn+") VALUES (?, ?, ?, ?, ?, ?)",
		customer.ID, customer.FullName, customer.DocumentNumber, customer.Phone, customer.Email, customer.CreatedAt,
	)
	return translate(err)
}

// List returns every customer, most recent first.
func (r CustomerRepository) List(ctx context.Context) ([]domain.Customer, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(
		queryCtx, "SELECT "+customerColumn+" FROM customer ORDER BY created_at DESC",
	)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	listed := make([]domain.Customer, 0)
	for rows.Next() {
		var customer domain.Customer
		if err := rows.Scan(
			&customer.ID, &customer.FullName, &customer.DocumentNumber,
			&customer.Phone, &customer.Email, &customer.CreatedAt,
		); err != nil {
			return nil, translate(err)
		}
		listed = append(listed, customer)
	}
	return listed, translate(rows.Err())
}

// FindByID reads one customer.
func (r CustomerRepository) FindByID(ctx context.Context, id string) (domain.Customer, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var customer domain.Customer
	err := r.database.QueryRowContext(
		queryCtx, "SELECT "+customerColumn+" FROM customer WHERE id = ?", id,
	).Scan(
		&customer.ID, &customer.FullName, &customer.DocumentNumber,
		&customer.Phone, &customer.Email, &customer.CreatedAt,
	)
	if err != nil {
		return domain.Customer{}, translate(err)
	}
	return customer, nil
}
