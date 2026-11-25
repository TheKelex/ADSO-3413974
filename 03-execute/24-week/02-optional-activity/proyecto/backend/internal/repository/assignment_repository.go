package repository

import (
	"context"
	"database/sql"
	"time"

	"workshop/internal/domain"
)

const assignmentColumn = "id, service_order_id, technician_id, is_active, active_marker, assigned_at, released_at"

// AssignmentRepository persists which technician holds which service order.
type AssignmentRepository struct {
	database *sql.DB
	timeout  time.Duration
}

// NewAssignmentRepository wires the assignment adapter.
func NewAssignmentRepository(database *sql.DB, timeout time.Duration) AssignmentRepository {
	return AssignmentRepository{database: database, timeout: timeout}
}

// Save stores an assignment. The unique active marker turns a second active
// assignment for the same technician into a conflict, even under a race.
func (r AssignmentRepository) Save(ctx context.Context, assignment domain.Assignment) error {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.database.ExecContext(
		queryCtx,
		"INSERT INTO assignment ("+assignmentColumn+") VALUES (?, ?, ?, ?, ?, ?, ?)",
		assignment.ID, assignment.ServiceOrderID, assignment.TechnicianID,
		assignment.IsActive, assignment.ActiveMarker, assignment.AssignedAt, assignment.ReleasedAt,
	)
	return translate(err)
}

// FindActiveByServiceOrder reads the technician currently holding an order.
func (r AssignmentRepository) FindActiveByServiceOrder(ctx context.Context, serviceOrderID string) (domain.Assignment, error) {
	return r.findActive(ctx, "service_order_id", serviceOrderID)
}

// FindActiveByTechnician reads the order a technician is currently holding.
func (r AssignmentRepository) FindActiveByTechnician(ctx context.Context, technicianID string) (domain.Assignment, error) {
	return r.findActive(ctx, "technician_id", technicianID)
}

// ReleaseByServiceOrder frees the technician when the vehicle is delivered.
func (r AssignmentRepository) ReleaseByServiceOrder(ctx context.Context, serviceOrderID string, releasedAt time.Time) error {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.database.ExecContext(
		queryCtx,
		"UPDATE assignment SET is_active = 0, active_marker = NULL, released_at = ? "+
			"WHERE service_order_id = ? AND is_active = 1",
		releasedAt, serviceOrderID,
	)
	return translate(err)
}

func (r AssignmentRepository) findActive(ctx context.Context, column, value string) (domain.Assignment, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var assignment domain.Assignment
	var marker sql.NullString
	var releasedAt sql.NullTime
	err := r.database.QueryRowContext(
		queryCtx,
		"SELECT "+assignmentColumn+" FROM assignment WHERE "+column+" = ? AND is_active = 1",
		value,
	).Scan(
		&assignment.ID, &assignment.ServiceOrderID, &assignment.TechnicianID,
		&assignment.IsActive, &marker, &assignment.AssignedAt, &releasedAt,
	)
	if err != nil {
		return domain.Assignment{}, translate(err)
	}
	if marker.Valid {
		value := marker.String
		assignment.ActiveMarker = &value
	}
	if releasedAt.Valid {
		moment := releasedAt.Time
		assignment.ReleasedAt = &moment
	}
	return assignment, nil
}
