package repository

import (
	"context"
	"database/sql"
	"time"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

const warrantyColumn = "id, intervention_id, warranty_kind, coverage_month_count, issued_at, expiration_date, created_at"

// WarrantyRepository persists the coverage issued over an intervention.
type WarrantyRepository struct {
	database *sql.DB
	timeout  time.Duration
}

// NewWarrantyRepository wires the warranty adapter.
func NewWarrantyRepository(database *sql.DB, timeout time.Duration) WarrantyRepository {
	return WarrantyRepository{database: database, timeout: timeout}
}

// Save stores a warranty with its computed expiration.
func (r WarrantyRepository) Save(ctx context.Context, warranty domain.Warranty) error {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.database.ExecContext(
		queryCtx,
		"INSERT INTO warranty ("+warrantyColumn+") VALUES (?, ?, ?, ?, ?, ?, ?)",
		warranty.ID, warranty.InterventionID, string(warranty.Kind),
		warranty.CoverageMonthCount, warranty.IssuedAt, warranty.ExpirationDate, warranty.IssuedAt,
	)
	return translate(err)
}

// FindByID reads one warranty.
func (r WarrantyRepository) FindByID(ctx context.Context, id string) (domain.Warranty, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var warranty domain.Warranty
	var kind string
	var createdAt time.Time
	err := r.database.QueryRowContext(
		queryCtx, "SELECT "+warrantyColumn+" FROM warranty WHERE id = ?", id,
	).Scan(
		&warranty.ID, &warranty.InterventionID, &kind, &warranty.CoverageMonthCount,
		&warranty.IssuedAt, &warranty.ExpirationDate, &createdAt,
	)
	if err != nil {
		return domain.Warranty{}, translate(err)
	}
	warranty.Kind = domain.WarrantyKind(kind)
	return warranty, nil
}

// List returns every warranty with the order and the plate it belongs to. The
// validity flag is filled by the use case against the consulted date.
func (r WarrantyRepository) List(ctx context.Context) ([]usecase.WarrantyView, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(
		queryCtx,
		"SELECT w.id, w.intervention_id, w.warranty_kind, w.coverage_month_count, w.issued_at, "+
			"w.expiration_date, so.order_number, v.plate FROM warranty w "+
			"JOIN intervention i ON i.id = w.intervention_id "+
			"JOIN service_order so ON so.id = i.service_order_id "+
			"JOIN vehicle v ON v.id = so.vehicle_id ORDER BY w.issued_at DESC",
	)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	listed := make([]usecase.WarrantyView, 0)
	for rows.Next() {
		var view usecase.WarrantyView
		var kind string
		if err := rows.Scan(
			&view.Warranty.ID, &view.Warranty.InterventionID, &kind,
			&view.Warranty.CoverageMonthCount, &view.Warranty.IssuedAt,
			&view.Warranty.ExpirationDate, &view.OrderNumber, &view.VehiclePlate,
		); err != nil {
			return nil, translate(err)
		}
		view.Warranty.Kind = domain.WarrantyKind(kind)
		listed = append(listed, view)
	}
	return listed, translate(rows.Err())
}

// ListByVehicle reads every warranty issued on a vehicle for the timeline.
func (r WarrantyRepository) ListByVehicle(ctx context.Context, vehicleID string) ([]domain.Warranty, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(
		queryCtx,
		"SELECT w.id, w.intervention_id, w.warranty_kind, w.coverage_month_count, w.issued_at, "+
			"w.expiration_date FROM warranty w "+
			"JOIN intervention i ON i.id = w.intervention_id "+
			"JOIN service_order so ON so.id = i.service_order_id "+
			"WHERE so.vehicle_id = ? ORDER BY w.issued_at",
		vehicleID,
	)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	listed := make([]domain.Warranty, 0)
	for rows.Next() {
		var warranty domain.Warranty
		var kind string
		if err := rows.Scan(
			&warranty.ID, &warranty.InterventionID, &kind,
			&warranty.CoverageMonthCount, &warranty.IssuedAt, &warranty.ExpirationDate,
		); err != nil {
			return nil, translate(err)
		}
		warranty.Kind = domain.WarrantyKind(kind)
		listed = append(listed, warranty)
	}
	return listed, translate(rows.Err())
}
