package domain_test

import (
	"errors"
	"testing"
	"time"

	"workshop/internal/domain"
)

func TestVehicleNormalizesPlateAndVIN(t *testing.T) {
	vehicle, err := domain.NewVehicle("vehicle-1", "customer-1", " abc123 ", " vin000001 ", "Mazda", "3", 2019, time.Now())
	if err != nil {
		t.Fatalf("registering a valid vehicle failed: %v", err)
	}
	if vehicle.Plate != "ABC123" {
		t.Fatalf("the plate must be stored uppercase and trimmed, got %q", vehicle.Plate)
	}
	if vehicle.VIN != "VIN000001" {
		t.Fatalf("the VIN must be stored uppercase and trimmed, got %q", vehicle.VIN)
	}
}

func TestVehicleRejectsMissingIdentifyingData(t *testing.T) {
	now := time.Now()
	if _, err := domain.NewVehicle("vehicle-1", "customer-1", "  ", "VIN1", "Mazda", "3", 2019, now); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("an empty plate must be rejected, got %v", err)
	}
	if _, err := domain.NewVehicle("vehicle-1", "customer-1", "ABC123", "  ", "Mazda", "3", 2019, now); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("an empty VIN must be rejected, got %v", err)
	}
	if _, err := domain.NewVehicle("vehicle-1", "", "ABC123", "VIN1", "Mazda", "3", 2019, now); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("a vehicle without an owner must be rejected, got %v", err)
	}
}

func TestVehicleRejectsAnImpossibleModelYear(t *testing.T) {
	if _, err := domain.NewVehicle("vehicle-1", "customer-1", "ABC123", "VIN1", "Mazda", "3", 1800, time.Now()); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("a model year before the first automobile must be rejected, got %v", err)
	}
}
