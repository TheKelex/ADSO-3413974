package domain_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	"workshop/internal/domain"
)

func TestCustomerRejectsControlCharactersAndLength(t *testing.T) {
	now := time.Now()
	if _, err := domain.NewCustomer("customer-1", "Ana", "DOC-1", "555-0100", "ana@example.com", now); err != nil {
		t.Fatalf("a valid customer must be accepted: %v", err)
	}
	if _, err := domain.NewCustomer("customer-1", "Ana", "DOC-1", "555-0100\x07", "ana@example.com", now); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("a phone with a control character must be rejected, got %v", err)
	}
	if _, err := domain.NewCustomer("customer-1", strings.Repeat("a", 161), "DOC-1", "555-0100", "ana@example.com", now); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("a name longer than the schema column must be rejected, got %v", err)
	}
}

func TestCustomerRejectsMarkupInTheName(t *testing.T) {
	_, err := domain.NewCustomer("customer-1", "<script>alert(1)</script>", "DOC-1", "555-0100", "ana@example.com", time.Now())
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("markup in the full name must be rejected, got %v", err)
	}
}

func TestVehicleRejectsMarkupInStructuredFields(t *testing.T) {
	now := time.Now()
	if _, err := domain.NewVehicle("vehicle-1", "customer-1", "ABC123", "VIN1", "Ma<zda", "3", 2019, now); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("markup in the brand must be rejected, got %v", err)
	}
	if _, err := domain.NewVehicle("vehicle-1", "customer-1", "AB>12", "VIN1", "Mazda", "3", 2019, now); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("markup in the plate must be rejected, got %v", err)
	}
}

func TestServiceOrderRejectsControlCharactersInTheReportedFailure(t *testing.T) {
	_, err := domain.NewServiceOrder("order-1", "OS-0001", "vehicle-1", "Ruido\x00", time.Now())
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("a control character in the reported failure must be rejected, got %v", err)
	}
}

func TestDiagnosticRejectsControlCharacters(t *testing.T) {
	now := time.Now()
	if _, err := domain.NewDiagnostic("diagnostic-1", "order-1", "technician-1", "Fuga\x07", "Bomba", now); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("a control character in the finding must be rejected, got %v", err)
	}
	if _, err := domain.NewDiagnostic("diagnostic-1", "order-1", "technician-1", "Fuga", "Bomba\x07", now); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("a control character in the component must be rejected, got %v", err)
	}
}

func TestInterventionRejectsControlCharactersAndPartLength(t *testing.T) {
	now := time.Now()
	if _, err := domain.NewIntervention("intervention-1", "order-1", "technician-1", "Cambio\x07", 1.5, nil, now); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("a control character in the description must be rejected, got %v", err)
	}
	if _, err := domain.NewPartUsage("part-1", "intervention-1", strings.Repeat("p", 161), 2, now); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("a part name longer than the schema column must be rejected, got %v", err)
	}
}
