package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

const serviceOrderColumn = "id, order_number, vehicle_id, reported_failure, status, received_at, created_at, updated_at"

const serviceOrderSummarySelect = "SELECT so.id, so.order_number, so.vehicle_id, so.reported_failure, so.status, " +
	"so.received_at, so.created_at, so.updated_at, v.plate, COALESCE(u.full_name, '') " +
	"FROM service_order so " +
	"JOIN vehicle v ON v.id = so.vehicle_id " +
	"LEFT JOIN assignment a ON a.service_order_id = so.id AND a.is_active = 1 " +
	"LEFT JOIN technician t ON t.id = a.technician_id " +
	"LEFT JOIN `user` u ON u.id = t.user_id"

// ServiceOrderRepository persists service orders and their status history.
type ServiceOrderRepository struct {
	database *sql.DB
	timeout  time.Duration
}

// NewServiceOrderRepository wires the service order adapter.
func NewServiceOrderRepository(database *sql.DB, timeout time.Duration) ServiceOrderRepository {
	return ServiceOrderRepository{database: database, timeout: timeout}
}

// Save stores a service order opened at check-in.
func (r ServiceOrderRepository) Save(ctx context.Context, order domain.ServiceOrder) error {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.database.ExecContext(
		queryCtx,
		"INSERT INTO service_order ("+serviceOrderColumn+") VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		order.ID, order.OrderNumber, order.VehicleID, order.ReportedFailure,
		string(order.Status), order.ReceivedAt, order.CreatedAt, order.UpdatedAt,
	)
	return translate(err)
}

// FindByID reads one service order.
func (r ServiceOrderRepository) FindByID(ctx context.Context, id string) (domain.ServiceOrder, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var order domain.ServiceOrder
	var status string
	err := r.database.QueryRowContext(
		queryCtx, "SELECT "+serviceOrderColumn+" FROM service_order WHERE id = ?", id,
	).Scan(
		&order.ID, &order.OrderNumber, &order.VehicleID, &order.ReportedFailure,
		&status, &order.ReceivedAt, &order.CreatedAt, &order.UpdatedAt,
	)
	if err != nil {
		return domain.ServiceOrder{}, translate(err)
	}
	order.Status = domain.ServiceOrderStatus(status)
	return order, nil
}

// FindSummary reads one order with its vehicle plate and active technician.
func (r ServiceOrderRepository) FindSummary(ctx context.Context, id string) (usecase.ServiceOrderSummary, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(queryCtx, serviceOrderSummarySelect+" WHERE so.id = ?", id)
	if err != nil {
		return usecase.ServiceOrderSummary{}, translate(err)
	}
	defer func() { _ = rows.Close() }()
	listed, err := scanSummaries(rows)
	if err != nil {
		return usecase.ServiceOrderSummary{}, err
	}
	if len(listed) == 0 {
		return usecase.ServiceOrderSummary{}, domain.ErrNotFound
	}
	return listed[0], nil
}

// List returns the orders, optionally filtered by a lifecycle status.
func (r ServiceOrderRepository) List(ctx context.Context, status string) ([]usecase.ServiceOrderSummary, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := serviceOrderSummarySelect + " ORDER BY so.received_at DESC"
	argument := make([]any, 0, 1)
	if status != "" {
		query = serviceOrderSummarySelect + " WHERE so.status = ? ORDER BY so.received_at DESC"
		argument = append(argument, status)
	}
	rows, err := r.database.QueryContext(queryCtx, query, argument...)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()
	return scanSummaries(rows)
}

// ListByTechnician returns the orders currently held by one technician,
// optionally filtered by a lifecycle status.
func (r ServiceOrderRepository) ListByTechnician(ctx context.Context, technicianID, status string) ([]usecase.ServiceOrderSummary, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := serviceOrderSummarySelect + " WHERE a.technician_id = ? ORDER BY so.received_at DESC"
	argument := []any{technicianID}
	if status != "" {
		query = serviceOrderSummarySelect + " WHERE a.technician_id = ? AND so.status = ? ORDER BY so.received_at DESC"
		argument = append(argument, status)
	}
	rows, err := r.database.QueryContext(queryCtx, query, argument...)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()
	return scanSummaries(rows)
}

func scanSummaries(rows *sql.Rows) ([]usecase.ServiceOrderSummary, error) {
	listed := make([]usecase.ServiceOrderSummary, 0)
	for rows.Next() {
		var summary usecase.ServiceOrderSummary
		var statusValue string
		if err := rows.Scan(
			&summary.Order.ID, &summary.Order.OrderNumber, &summary.Order.VehicleID,
			&summary.Order.ReportedFailure, &statusValue, &summary.Order.ReceivedAt,
			&summary.Order.CreatedAt, &summary.Order.UpdatedAt,
			&summary.VehiclePlate, &summary.TechnicianName,
		); err != nil {
			return nil, translate(err)
		}
		summary.Order.Status = domain.ServiceOrderStatus(statusValue)
		listed = append(listed, summary)
	}
	return listed, translate(rows.Err())
}

// ListByVehicle returns every order of a vehicle, oldest first, for the
// clinical timeline.
func (r ServiceOrderRepository) ListByVehicle(ctx context.Context, vehicleID string) ([]domain.ServiceOrder, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(
		queryCtx,
		"SELECT "+serviceOrderColumn+" FROM service_order WHERE vehicle_id = ? ORDER BY received_at",
		vehicleID,
	)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	listed := make([]domain.ServiceOrder, 0)
	for rows.Next() {
		var order domain.ServiceOrder
		var status string
		if err := rows.Scan(
			&order.ID, &order.OrderNumber, &order.VehicleID, &order.ReportedFailure,
			&status, &order.ReceivedAt, &order.CreatedAt, &order.UpdatedAt,
		); err != nil {
			return nil, translate(err)
		}
		order.Status = domain.ServiceOrderStatus(status)
		listed = append(listed, order)
	}
	return listed, translate(rows.Err())
}

// UpdateStatus writes the new status and its transition record inside one
// transaction, so the history can never disagree with the order.
func (r ServiceOrderRepository) UpdateStatus(ctx context.Context, order domain.ServiceOrder, transition domain.StatusTransition) error {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	transaction, err := r.database.BeginTx(queryCtx, nil)
	if err != nil {
		return translate(err)
	}
	defer func() { _ = transaction.Rollback() }()

	if _, err := transaction.ExecContext(
		queryCtx,
		"UPDATE service_order SET status = ?, updated_at = ? WHERE id = ?",
		string(order.Status), order.UpdatedAt, order.ID,
	); err != nil {
		return translate(err)
	}
	if _, err := transaction.ExecContext(
		queryCtx,
		"INSERT INTO status_transition (id, service_order_id, from_status, to_status, changed_by_user_id, changed_at) "+
			"VALUES (?, ?, ?, ?, ?, ?)",
		transition.ID, transition.ServiceOrderID, string(transition.FromStatus),
		string(transition.ToStatus), transition.ChangedByUserID, transition.ChangedAt,
	); err != nil {
		return translate(err)
	}
	return translate(transaction.Commit())
}

// ListTransition returns the status history of an order, oldest first.
func (r ServiceOrderRepository) ListTransition(ctx context.Context, serviceOrderID string) ([]domain.StatusTransition, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(
		queryCtx,
		"SELECT id, service_order_id, from_status, to_status, changed_by_user_id, changed_at "+
			"FROM status_transition WHERE service_order_id = ? "+
			"ORDER BY changed_at, FIELD(from_status, 'RECEIVED', 'IN_DIAGNOSIS', 'IN_REPAIR', 'READY', 'DELIVERED'), id",
		serviceOrderID,
	)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	listed := make([]domain.StatusTransition, 0)
	for rows.Next() {
		var transition domain.StatusTransition
		var from, to string
		if err := rows.Scan(
			&transition.ID, &transition.ServiceOrderID, &from, &to,
			&transition.ChangedByUserID, &transition.ChangedAt,
		); err != nil {
			return nil, translate(err)
		}
		transition.FromStatus = domain.ServiceOrderStatus(from)
		transition.ToStatus = domain.ServiceOrderStatus(to)
		listed = append(listed, transition)
	}
	return listed, translate(rows.Err())
}

// CountByStatus returns how many orders sit in each lifecycle status.
func (r ServiceOrderRepository) CountByStatus(ctx context.Context) (map[string]int, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(
		queryCtx, "SELECT status, COUNT(*) FROM service_order GROUP BY status",
	)
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	counted := make(map[string]int)
	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return nil, translate(err)
		}
		counted[status] = count
	}
	return counted, translate(rows.Err())
}

// NextOrderNumber builds the next sequential workshop order number.
func (r ServiceOrderRepository) NextOrderNumber(ctx context.Context) (string, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var count int
	if err := r.database.QueryRowContext(
		queryCtx, "SELECT COUNT(*) FROM service_order",
	).Scan(&count); err != nil {
		return "", translate(err)
	}
	return fmt.Sprintf("OS-%04d", count+1), nil
}
