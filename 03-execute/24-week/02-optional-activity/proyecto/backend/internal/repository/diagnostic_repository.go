package repository

import (
	"context"
	"database/sql"
	"time"

	"workshop/internal/domain"
)

const diagnosticColumn = "id, service_order_id, technician_id, finding, component_to_repair, created_at"

// DiagnosticRepository persists the technical evaluation of a service order.
type DiagnosticRepository struct {
	database *sql.DB
	timeout  time.Duration
}

// NewDiagnosticRepository wires the diagnostic adapter.
func NewDiagnosticRepository(database *sql.DB, timeout time.Duration) DiagnosticRepository {
	return DiagnosticRepository{database: database, timeout: timeout}
}

// Save stores a diagnostic.
func (r DiagnosticRepository) Save(ctx context.Context, diagnostic domain.Diagnostic) error {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.database.ExecContext(
		queryCtx,
		"INSERT INTO diagnostic ("+diagnosticColumn+") VALUES (?, ?, ?, ?, ?, ?)",
		diagnostic.ID, diagnostic.ServiceOrderID, diagnostic.TechnicianID,
		diagnostic.Finding, diagnostic.ComponentToRepair, diagnostic.CreatedAt,
	)
	return translate(err)
}

// FindByServiceOrder reads the most recent diagnostic of an order.
func (r DiagnosticRepository) FindByServiceOrder(ctx context.Context, serviceOrderID string) (domain.Diagnostic, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var diagnostic domain.Diagnostic
	err := r.database.QueryRowContext(
		queryCtx,
		"SELECT "+diagnosticColumn+" FROM diagnostic WHERE service_order_id = ? ORDER BY created_at DESC LIMIT 1",
		serviceOrderID,
	).Scan(
		&diagnostic.ID, &diagnostic.ServiceOrderID, &diagnostic.TechnicianID,
		&diagnostic.Finding, &diagnostic.ComponentToRepair, &diagnostic.CreatedAt,
	)
	if err != nil {
		return domain.Diagnostic{}, translate(err)
	}
	return diagnostic, nil
}

// ListByVehicle reads every diagnostic recorded on a vehicle, oldest first,
// for the clinical timeline.
func (r DiagnosticRepository) ListByVehicle(ctx context.Context, vehicleID string) ([]domain.Diagnostic, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(
		queryCtx,
		"SELECT d.id, d.service_order_id, d.technician_id, d.finding, d.component_to_repair, d.created_at "+
			"FROM diagnostic d JOIN service_order so ON so.id = d.service_order_id "+
			"WHERE so.vehicle_id = ? ORDER BY d.created_at",
		vehicleID,
	)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	listed := make([]domain.Diagnostic, 0)
	for rows.Next() {
		var diagnostic domain.Diagnostic
		if err := rows.Scan(
			&diagnostic.ID, &diagnostic.ServiceOrderID, &diagnostic.TechnicianID,
			&diagnostic.Finding, &diagnostic.ComponentToRepair, &diagnostic.CreatedAt,
		); err != nil {
			return nil, translate(err)
		}
		listed = append(listed, diagnostic)
	}
	return listed, translate(rows.Err())
}
