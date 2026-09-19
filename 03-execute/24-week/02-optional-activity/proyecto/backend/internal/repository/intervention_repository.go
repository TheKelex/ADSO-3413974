package repository

import (
	"context"
	"database/sql"
	"time"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

const interventionColumn = "id, service_order_id, technician_id, description, labor_hour_count, performed_at, created_at"

// InterventionRepository persists the work executed on a vehicle together with
// the parts it consumed.
type InterventionRepository struct {
	database *sql.DB
	timeout  time.Duration
}

// NewInterventionRepository wires the intervention adapter.
func NewInterventionRepository(database *sql.DB, timeout time.Duration) InterventionRepository {
	return InterventionRepository{database: database, timeout: timeout}
}

// Save writes the intervention and its parts in one transaction, so an
// intervention is never stored without the parts it declared.
func (r InterventionRepository) Save(ctx context.Context, intervention domain.Intervention) error {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	transaction, err := r.database.BeginTx(queryCtx, nil)
	if err != nil {
		return translate(err)
	}
	defer func() { _ = transaction.Rollback() }()

	if _, err := transaction.ExecContext(
		queryCtx,
		"INSERT INTO intervention ("+interventionColumn+") VALUES (?, ?, ?, ?, ?, ?, ?)",
		intervention.ID, intervention.ServiceOrderID, intervention.TechnicianID,
		intervention.Description, intervention.LaborHourCount,
		intervention.PerformedAt, intervention.CreatedAt,
	); err != nil {
		return translate(err)
	}
	for _, part := range intervention.Part {
		if _, err := transaction.ExecContext(
			queryCtx,
			"INSERT INTO part_usage (id, intervention_id, part_name, quantity, created_at) VALUES (?, ?, ?, ?, ?)",
			part.ID, intervention.ID, part.PartName, part.Quantity, part.CreatedAt,
		); err != nil {
			return translate(err)
		}
	}
	return translate(transaction.Commit())
}

// FindByID reads one intervention with its parts.
func (r InterventionRepository) FindByID(ctx context.Context, id string) (domain.Intervention, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var intervention domain.Intervention
	err := r.database.QueryRowContext(
		queryCtx, "SELECT "+interventionColumn+" FROM intervention WHERE id = ?", id,
	).Scan(
		&intervention.ID, &intervention.ServiceOrderID, &intervention.TechnicianID,
		&intervention.Description, &intervention.LaborHourCount,
		&intervention.PerformedAt, &intervention.CreatedAt,
	)
	if err != nil {
		return domain.Intervention{}, translate(err)
	}
	part, err := r.listPart(queryCtx, intervention.ID)
	if err != nil {
		return domain.Intervention{}, err
	}
	intervention.Part = part
	return intervention, nil
}

// ListByServiceOrder reads the interventions of an order, oldest first.
func (r InterventionRepository) ListByServiceOrder(ctx context.Context, serviceOrderID string) ([]domain.Intervention, error) {
	return r.list(
		ctx,
		"SELECT "+interventionColumn+" FROM intervention WHERE service_order_id = ? ORDER BY performed_at",
		serviceOrderID,
	)
}

// ListByServiceOrderView reads interventions with their latest warranty, if
// one exists, while preserving interventions that have no warranty.
func (r InterventionRepository) ListByServiceOrderView(ctx context.Context, serviceOrderID string) ([]usecase.InterventionView, error) {
	query := "SELECT i.id, i.service_order_id, i.technician_id, i.description, i.labor_hour_count, " +
		"i.performed_at, i.created_at, w.id, w.warranty_kind, w.expiration_date " +
		"FROM intervention i " +
		"LEFT JOIN warranty w ON w.intervention_id = i.id " +
		"LEFT JOIN warranty newer ON newer.intervention_id = w.intervention_id AND " +
		"(newer.issued_at > w.issued_at OR (newer.issued_at = w.issued_at AND newer.id > w.id)) " +
		"WHERE i.service_order_id = ? AND newer.id IS NULL ORDER BY i.performed_at"
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(queryCtx, query, serviceOrderID)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	listed := make([]usecase.InterventionView, 0)
	for rows.Next() {
		var view usecase.InterventionView
		var warrantyID, warrantyKind sql.NullString
		var warrantyExpiration sql.NullTime
		if err := rows.Scan(
			&view.Intervention.ID, &view.Intervention.ServiceOrderID, &view.Intervention.TechnicianID,
			&view.Intervention.Description, &view.Intervention.LaborHourCount,
			&view.Intervention.PerformedAt, &view.Intervention.CreatedAt,
			&warrantyID, &warrantyKind, &warrantyExpiration,
		); err != nil {
			return nil, translate(err)
		}
		if warrantyID.Valid {
			view.WarrantyID = &warrantyID.String
			kind := domain.WarrantyKind(warrantyKind.String)
			view.WarrantyKind = &kind
			if warrantyExpiration.Valid {
				expiration := warrantyExpiration.Time
				view.WarrantyExpiration = &expiration
			}
		}
		listed = append(listed, view)
	}
	if err := rows.Err(); err != nil {
		return nil, translate(err)
	}
	for index := range listed {
		part, err := r.listPart(queryCtx, listed[index].Intervention.ID)
		if err != nil {
			return nil, err
		}
		listed[index].Intervention.Part = part
	}
	return listed, nil
}

// ListByVehicle reads every intervention of a vehicle for the timeline.
func (r InterventionRepository) ListByVehicle(ctx context.Context, vehicleID string) ([]domain.Intervention, error) {
	return r.list(
		ctx,
		"SELECT i.id, i.service_order_id, i.technician_id, i.description, i.labor_hour_count, "+
			"i.performed_at, i.created_at FROM intervention i "+
			"JOIN service_order so ON so.id = i.service_order_id "+
			"WHERE so.vehicle_id = ? ORDER BY i.performed_at",
		vehicleID,
	)
}

func (r InterventionRepository) list(ctx context.Context, query, argument string) ([]domain.Intervention, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(queryCtx, query, argument)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	listed := make([]domain.Intervention, 0)
	for rows.Next() {
		var intervention domain.Intervention
		if err := rows.Scan(
			&intervention.ID, &intervention.ServiceOrderID, &intervention.TechnicianID,
			&intervention.Description, &intervention.LaborHourCount,
			&intervention.PerformedAt, &intervention.CreatedAt,
		); err != nil {
			return nil, translate(err)
		}
		listed = append(listed, intervention)
	}
	if err := rows.Err(); err != nil {
		return nil, translate(err)
	}
	for index := range listed {
		part, err := r.listPart(queryCtx, listed[index].ID)
		if err != nil {
			return nil, err
		}
		listed[index].Part = part
	}
	return listed, nil
}

func (r InterventionRepository) listPart(ctx context.Context, interventionID string) ([]domain.PartUsage, error) {
	rows, err := r.database.QueryContext(
		ctx,
		"SELECT id, intervention_id, part_name, quantity, created_at FROM part_usage "+
			"WHERE intervention_id = ? ORDER BY created_at",
		interventionID,
	)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	listed := make([]domain.PartUsage, 0)
	for rows.Next() {
		var part domain.PartUsage
		if err := rows.Scan(
			&part.ID, &part.InterventionID, &part.PartName, &part.Quantity, &part.CreatedAt,
		); err != nil {
			return nil, translate(err)
		}
		listed = append(listed, part)
	}
	return listed, translate(rows.Err())
}
