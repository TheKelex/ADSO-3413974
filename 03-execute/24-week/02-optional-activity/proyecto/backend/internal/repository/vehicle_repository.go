package repository

import (
	"context"
	"database/sql"
	"time"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

const vehicleSelect = "SELECT v.id, v.customer_id, v.plate, v.vin, v.brand, v.model, v.model_year, v.created_at, " +
	"c.full_name FROM vehicle v JOIN customer c ON c.id = v.customer_id"

// VehicleRepository persists and reads the vehicles of the workshop.
type VehicleRepository struct {
	database *sql.DB
	timeout  time.Duration
}

// NewVehicleRepository wires the vehicle adapter.
func NewVehicleRepository(database *sql.DB, timeout time.Duration) VehicleRepository {
	return VehicleRepository{database: database, timeout: timeout}
}

// Save stores a vehicle. A duplicate plate or VIN surfaces as a conflict.
func (r VehicleRepository) Save(ctx context.Context, vehicle domain.Vehicle) error {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.database.ExecContext(
		queryCtx,
		"INSERT INTO vehicle (id, customer_id, plate, vin, brand, model, model_year, created_at) "+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		vehicle.ID, vehicle.CustomerID, vehicle.Plate, vehicle.VIN,
		vehicle.Brand, vehicle.Model, vehicle.ModelYear, vehicle.CreatedAt,
	)
	return translate(err)
}

// List returns every vehicle with the name of its owner.
func (r VehicleRepository) List(ctx context.Context) ([]usecase.VehicleWithOwner, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.database.QueryContext(queryCtx, vehicleSelect+" ORDER BY v.created_at DESC")
	if err != nil {
		return nil, translate(err)
	}
	defer func() { _ = rows.Close() }()

	listed := make([]usecase.VehicleWithOwner, 0)
	for rows.Next() {
		item, scanErr := scanVehicle(rows)
		if scanErr != nil {
			return nil, translate(scanErr)
		}
		listed = append(listed, item)
	}
	return listed, translate(rows.Err())
}

// FindByID reads one vehicle with the name of its owner.
func (r VehicleRepository) FindByID(ctx context.Context, id string) (usecase.VehicleWithOwner, error) {
	queryCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	row := r.database.QueryRowContext(queryCtx, vehicleSelect+" WHERE v.id = ?", id)
	item, err := scanVehicle(row)
	if err != nil {
		return usecase.VehicleWithOwner{}, translate(err)
	}
	return item, nil
}

// scanner is what a *sql.Row and a *sql.Rows have in common.
type scanner interface {
	Scan(destination ...any) error
}

func scanVehicle(source scanner) (usecase.VehicleWithOwner, error) {
	var vehicle domain.Vehicle
	var ownerName string
	if err := source.Scan(
		&vehicle.ID, &vehicle.CustomerID, &vehicle.Plate, &vehicle.VIN,
		&vehicle.Brand, &vehicle.Model, &vehicle.ModelYear, &vehicle.CreatedAt, &ownerName,
	); err != nil {
		return usecase.VehicleWithOwner{}, err
	}
	return usecase.VehicleWithOwner{
		Vehicle:   vehicle,
		OwnerID:   vehicle.CustomerID,
		OwnerName: ownerName,
	}, nil
}
