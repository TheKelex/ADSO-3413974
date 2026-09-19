package repository

import (
	"context"
	"database/sql"
	"time"

	"workshop/internal/domain"
)

const technicianColumn = "id, user_id, specialty, created_at"

// workloadSelect reads every technician with the order they currently hold.
// The left joins keep a free technician in the result with empty order data.
const workloadSelect = "SELECT t.id, t.user_id, t.specialty, t.created_at, u.full_name, " +
	"COALESCE(a.service_order_id, ''), COALESCE(so.order_number, ''), COALESCE(v.plate, '') " +
	"FROM technician t " +
	"JOIN `user` u ON u.id = t.user_id " +
	"LEFT JOIN assignment a ON a.technician_id = t.id AND a.is_active = 1 " +
	"LEFT JOIN service_order so ON so.id = a.service_order_id " +
	"LEFT JOIN vehicle v ON v.id = so.vehicle_id " +
	"ORDER BY u.full_name"

// TechnicianRepository reads the mechanic profiles and their workload.
type TechnicianRepository struct {
	database *sql.DB
	timeout  time.Duration
}

// NewTechnicianRepository wires the technician adapter.
func NewTechnicianRepository(database *sql.DB, timeout time.Duration) TechnicianRepository {
	return TechnicianRepository{database: database, timeout: timeout}
}

// ListWorkload returns every technician and the order they are working on.
func (r TechnicianRepository) ListWorkload(ctx context.Context) ([]domain.TechnicianWorkload, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(queryCtx, workloadSelect)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	listed := make([]domain.TechnicianWorkload, 0)
	for rows.Next() {
		var workload domain.TechnicianWorkload
		if err := rows.Scan(
			&workload.Technician.ID, &workload.Technician.UserID,
			&workload.Technician.Specialty, &workload.Technician.CreatedAt,
			&workload.FullName, &workload.ActiveOrderID,
			&workload.ActiveOrderNumber, &workload.ActiveVehiclePlate,
		); err != nil {
			return nil, translate(err)
		}
		workload.Busy = workload.ActiveOrderID != ""
		listed = append(listed, workload)
	}
	return listed, translate(rows.Err())
}

// FindByID reads one mechanic profile.
func (r TechnicianRepository) FindByID(ctx context.Context, id string) (domain.Technician, error) {
	return r.findBy(ctx, "SELECT "+technicianColumn+" FROM technician WHERE id = ?", id)
}

// FindByUserID resolves the mechanic profile of a signed in user.
func (r TechnicianRepository) FindByUserID(ctx context.Context, userID string) (domain.Technician, error) {
	return r.findBy(ctx, "SELECT "+technicianColumn+" FROM technician WHERE user_id = ?", userID)
}

func (r TechnicianRepository) findBy(ctx context.Context, query string, argument any) (domain.Technician, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var technician domain.Technician
	err := r.database.QueryRowContext(queryCtx, query, argument).Scan(
		&technician.ID, &technician.UserID, &technician.Specialty, &technician.CreatedAt,
	)
	if err != nil {
		return domain.Technician{}, translate(err)
	}
	return technician, nil
}
