package usecase_test

import (
	"context"
	"errors"
	"testing"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

type fakeDiagnosticRepository struct {
	byOrder map[string]domain.Diagnostic
	saved   int
}

func (f *fakeDiagnosticRepository) Save(_ context.Context, diagnostic domain.Diagnostic) error {
	f.byOrder[diagnostic.ServiceOrderID] = diagnostic
	f.saved++
	return nil
}

func (f *fakeDiagnosticRepository) FindByServiceOrder(_ context.Context, serviceOrderID string) (domain.Diagnostic, error) {
	found, ok := f.byOrder[serviceOrderID]
	if !ok {
		return domain.Diagnostic{}, domain.ErrNotFound
	}
	return found, nil
}

func (f *fakeDiagnosticRepository) ListByVehicle(_ context.Context, _ string) ([]domain.Diagnostic, error) {
	return nil, nil
}

func buildActiveAssignment(t *testing.T, id, orderID, technicianID string) domain.Assignment {
	t.Helper()
	assignment, err := domain.NewAssignment(id, orderID, technicianID, fixedClock()())
	if err != nil {
		t.Fatalf("building the assignment fixture failed: %v", err)
	}
	return assignment
}

func TestRecordIsForbiddenWhenTheOrderWasAlreadyDelivered(t *testing.T) {
	orders := newFakeOrderRepository(buildOrder(t, "order-1", "OS-0001"))
	delivered := orders.order["order-1"]
	delivered.Status = domain.StatusDelivered
	orders.order["order-1"] = delivered
	technicians := newFakeTechnicianRepository(buildTechnician(t, "technician-1", "user-1"))
	assignments := &fakeAssignmentRepository{}
	if err := assignments.Save(context.Background(), buildActiveAssignment(t, "assignment-1", "order-1", "technician-1")); err != nil {
		t.Fatalf("storing the assignment fixture failed: %v", err)
	}
	diagnosticStore := &fakeDiagnosticRepository{byOrder: map[string]domain.Diagnostic{}}
	useCase := usecase.NewDiagnosticUseCase(diagnosticStore, orders, assignments, technicians, sequentialID(), fixedClock())

	_, err := useCase.Record(context.Background(), "order-1", "user-1", "Ruido", "Bomba")
	if !errors.Is(err, domain.ErrForbidden) {
		t.Fatalf("a delivered order must refuse the diagnostic, got %v", err)
	}
	if diagnosticStore.saved != 0 {
		t.Fatalf("a refused diagnostic must not be stored, got %d saves", diagnosticStore.saved)
	}
	if orders.updates != 0 {
		t.Fatalf("a refused diagnostic must not move the order, got %d updates", orders.updates)
	}
}

func TestRegisterIsForbiddenWhenTheOrderWasAlreadyDelivered(t *testing.T) {
	orders := newFakeOrderRepository(buildOrder(t, "order-1", "OS-0001"))
	delivered := orders.order["order-1"]
	delivered.Status = domain.StatusDelivered
	orders.order["order-1"] = delivered
	technicians := newFakeTechnicianRepository(buildTechnician(t, "technician-1", "user-1"))
	assignments := &fakeAssignmentRepository{}
	if err := assignments.Save(context.Background(), buildActiveAssignment(t, "assignment-1", "order-1", "technician-1")); err != nil {
		t.Fatalf("storing the assignment fixture failed: %v", err)
	}
	interventionStore := &fakeInterventionRepository{}
	useCase := usecase.NewInterventionUseCase(interventionStore, orders, assignments, technicians, sequentialID(), fixedClock())

	_, err := useCase.Register(context.Background(), "order-1", "user-1", "Cambio de correa", 1.5, nil)
	if !errors.Is(err, domain.ErrForbidden) {
		t.Fatalf("a delivered order must refuse the intervention, got %v", err)
	}
	if len(interventionStore.intervention) != 0 {
		t.Fatalf("a refused intervention must not be stored, got %d stores", len(interventionStore.intervention))
	}
	if orders.updates != 0 {
		t.Fatalf("a refused intervention must not move the order, got %d updates", orders.updates)
	}
}
