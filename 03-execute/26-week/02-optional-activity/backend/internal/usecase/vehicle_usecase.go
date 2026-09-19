package usecase

import (
	"context"
	"time"

	"workshop/internal/domain"
)

// VehicleWithOwner is the read model the vehicle list and the timeline header
// consume: the vehicle plus the name of the customer who owns it.
type VehicleWithOwner struct {
	Vehicle   domain.Vehicle
	OwnerID   string
	OwnerName string
}

// VehicleRepository is the narrow port the vehicle use case needs.
type VehicleRepository interface {
	Save(ctx context.Context, vehicle domain.Vehicle) error
	List(ctx context.Context) ([]VehicleWithOwner, error)
	FindByID(ctx context.Context, id string) (VehicleWithOwner, error)
}

// VehicleUseCase registers and lists the vehicles of the workshop.
type VehicleUseCase struct {
	vehicle  VehicleRepository
	customer CustomerRepository
	newID    func() string
	now      func() time.Time
}

// NewVehicleUseCase wires the vehicle use case.
func NewVehicleUseCase(vehicle VehicleRepository, customer CustomerRepository, newID func() string, now func() time.Time) VehicleUseCase {
	return VehicleUseCase{vehicle: vehicle, customer: customer, newID: newID, now: now}
}

// Register stores a vehicle for an existing customer. A duplicate plate or VIN
// is rejected by the storage layer and surfaces as a conflict.
func (v VehicleUseCase) Register(ctx context.Context, customerID, plate, vin, brand, model string, modelYear int) (domain.Vehicle, error) {
	if _, err := v.customer.FindByID(ctx, customerID); err != nil {
		return domain.Vehicle{}, err
	}
	vehicle, err := domain.NewVehicle(v.newID(), customerID, plate, vin, brand, model, modelYear, v.now())
	if err != nil {
		return domain.Vehicle{}, err
	}
	if err := v.vehicle.Save(ctx, vehicle); err != nil {
		return domain.Vehicle{}, err
	}
	return vehicle, nil
}

// List returns every vehicle with the name of its owner.
func (v VehicleUseCase) List(ctx context.Context) ([]VehicleWithOwner, error) {
	return v.vehicle.List(ctx)
}

// Find returns one vehicle with the name of its owner.
func (v VehicleUseCase) Find(ctx context.Context, vehicleID string) (VehicleWithOwner, error) {
	return v.vehicle.FindByID(ctx, vehicleID)
}
