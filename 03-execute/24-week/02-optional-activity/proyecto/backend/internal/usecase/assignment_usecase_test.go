package usecase_test

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

// fixedClock keeps the tests independent from the wall clock.
func fixedClock() func() time.Time {
	moment := time.Date(2026, time.March, 1, 9, 0, 0, 0, time.UTC)
	return func() time.Time { return moment }
}

// sequentialID hands out predictable identifiers so assertions can name them.
func sequentialID() func() string {
	counter := 0
	return func() string {
		counter++
		return "id-" + strconv.Itoa(counter)
	}
}

type fakeOrderRepository struct {
	order      map[string]domain.ServiceOrder
	transition []domain.StatusTransition
	updates    int
	holder     map[string]string
}

func newFakeOrderRepository(order ...domain.ServiceOrder) *fakeOrderRepository {
	stored := make(map[string]domain.ServiceOrder, len(order))
	for _, item := range order {
		stored[item.ID] = item
	}
	return &fakeOrderRepository{
		order:  stored,
		holder: make(map[string]string),
	}
}

func (f *fakeOrderRepository) heldBy(orderID, technicianID string) {
	f.holder[orderID] = technicianID
}

func (f *fakeOrderRepository) Save(_ context.Context, order domain.ServiceOrder) error {
	f.order[order.ID] = order
	return nil
}

func (f *fakeOrderRepository) FindByID(_ context.Context, id string) (domain.ServiceOrder, error) {
	found, ok := f.order[id]
	if !ok {
		return domain.ServiceOrder{}, domain.ErrNotFound
	}
	return found, nil
}

func (f *fakeOrderRepository) FindSummary(ctx context.Context, id string) (usecase.ServiceOrderSummary, error) {
	order, err := f.FindByID(ctx, id)
	if err != nil {
		return usecase.ServiceOrderSummary{}, err
	}
	return usecase.ServiceOrderSummary{Order: order}, nil
}

func (f *fakeOrderRepository) List(_ context.Context, _ string) ([]usecase.ServiceOrderSummary, error) {
	listing := make([]usecase.ServiceOrderSummary, 0, len(f.order))
	for _, order := range f.order {
		listing = append(listing, usecase.ServiceOrderSummary{Order: order})
	}
	return listing, nil
}

func (f *fakeOrderRepository) ListByTechnician(_ context.Context, technicianID, _ string) ([]usecase.ServiceOrderSummary, error) {
	listing := make([]usecase.ServiceOrderSummary, 0)
	for id, holder := range f.holder {
		if holder != technicianID {
			continue
		}
		order, ok := f.order[id]
		if ok {
			listing = append(listing, usecase.ServiceOrderSummary{Order: order})
		}
	}
	return listing, nil
}

func (f *fakeOrderRepository) ListByVehicle(_ context.Context, _ string) ([]domain.ServiceOrder, error) {
	return nil, nil
}

func (f *fakeOrderRepository) UpdateStatus(_ context.Context, order domain.ServiceOrder, transition domain.StatusTransition) error {
	f.order[order.ID] = order
	f.transition = append(f.transition, transition)
	f.updates++
	return nil
}

func (f *fakeOrderRepository) ListTransition(_ context.Context, serviceOrderID string) ([]domain.StatusTransition, error) {
	result := make([]domain.StatusTransition, 0)
	for _, item := range f.transition {
		if item.ServiceOrderID == serviceOrderID {
			result = append(result, item)
		}
	}
	return result, nil
}

func (f *fakeOrderRepository) CountByStatus(_ context.Context) (map[string]int, error) {
	counted := make(map[string]int)
	for _, item := range f.order {
		counted[string(item.Status)]++
	}
	return counted, nil
}

func (f *fakeOrderRepository) NextOrderNumber(_ context.Context) (string, error) {
	return "OS-0001", nil
}

type fakeAssignmentRepository struct {
	assignment []domain.Assignment
	released   int
}

func (f *fakeAssignmentRepository) Save(_ context.Context, assignment domain.Assignment) error {
	f.assignment = append(f.assignment, assignment)
	return nil
}

func (f *fakeAssignmentRepository) FindActiveByServiceOrder(_ context.Context, serviceOrderID string) (domain.Assignment, error) {
	for _, item := range f.assignment {
		if item.IsActive && item.ServiceOrderID == serviceOrderID {
			return item, nil
		}
	}
	return domain.Assignment{}, domain.ErrNotFound
}

func (f *fakeAssignmentRepository) FindActiveByTechnician(_ context.Context, technicianID string) (domain.Assignment, error) {
	for _, item := range f.assignment {
		if item.IsActive && item.TechnicianID == technicianID {
			return item, nil
		}
	}
	return domain.Assignment{}, domain.ErrNotFound
}

func (f *fakeAssignmentRepository) ReleaseByServiceOrder(_ context.Context, serviceOrderID string, releasedAt time.Time) error {
	for index := range f.assignment {
		if f.assignment[index].ServiceOrderID == serviceOrderID && f.assignment[index].IsActive {
			f.assignment[index].Release(releasedAt)
			f.released++
		}
	}
	return nil
}

type fakeTechnicianRepository struct {
	technician map[string]domain.Technician
	byUser     map[string]string
}

func newFakeTechnicianRepository(technician ...domain.Technician) *fakeTechnicianRepository {
	stored := make(map[string]domain.Technician, len(technician))
	byUser := make(map[string]string, len(technician))
	for _, item := range technician {
		stored[item.ID] = item
		byUser[item.UserID] = item.ID
	}
	return &fakeTechnicianRepository{technician: stored, byUser: byUser}
}

func (f *fakeTechnicianRepository) ListWorkload(_ context.Context) ([]domain.TechnicianWorkload, error) {
	workload := make([]domain.TechnicianWorkload, 0, len(f.technician))
	for _, item := range f.technician {
		workload = append(workload, domain.TechnicianWorkload{Technician: item})
	}
	return workload, nil
}

func (f *fakeTechnicianRepository) FindByID(_ context.Context, id string) (domain.Technician, error) {
	found, ok := f.technician[id]
	if !ok {
		return domain.Technician{}, domain.ErrNotFound
	}
	return found, nil
}

func (f *fakeTechnicianRepository) FindByUserID(_ context.Context, userID string) (domain.Technician, error) {
	id, ok := f.byUser[userID]
	if !ok {
		return domain.Technician{}, domain.ErrNotFound
	}
	return f.technician[id], nil
}

func buildOrder(t *testing.T, id, number string) domain.ServiceOrder {
	t.Helper()
	order, err := domain.NewServiceOrder(id, number, "vehicle-1", "Ruido en el motor", fixedClock()())
	if err != nil {
		t.Fatalf("building the order fixture failed: %v", err)
	}
	return order
}

func buildTechnician(t *testing.T, id, userID string) domain.Technician {
	t.Helper()
	technician, err := domain.NewTechnician(id, userID, "Motor", fixedClock()())
	if err != nil {
		t.Fatalf("building the technician fixture failed: %v", err)
	}
	return technician
}

func TestAssignRejectsATechnicianWhoAlreadyHoldsAnActiveOrder(t *testing.T) {
	orders := newFakeOrderRepository(buildOrder(t, "order-1", "OS-0001"), buildOrder(t, "order-2", "OS-0002"))
	technicians := newFakeTechnicianRepository(buildTechnician(t, "technician-1", "user-1"))
	assignments := &fakeAssignmentRepository{}
	useCase := usecase.NewAssignmentUseCase(assignments, orders, technicians, sequentialID(), fixedClock())

	if _, err := useCase.Assign(context.Background(), "order-1", "technician-1"); err != nil {
		t.Fatalf("the first assignment must be accepted: %v", err)
	}
	_, err := useCase.Assign(context.Background(), "order-2", "technician-1")
	if err == nil {
		t.Fatal("a technician holding an active order must not receive a second one")
	}
	if !errors.Is(err, domain.ErrConflict) {
		t.Fatalf("the rejection must be a conflict, got %v", err)
	}
	if len(assignments.assignment) != 1 {
		t.Fatalf("the rejected assignment must not be stored, got %d", len(assignments.assignment))
	}
}

func TestAssignAcceptsAFreeTechnicianAndMarksTheAssignmentActive(t *testing.T) {
	orders := newFakeOrderRepository(buildOrder(t, "order-1", "OS-0001"))
	technicians := newFakeTechnicianRepository(buildTechnician(t, "technician-1", "user-1"))
	assignments := &fakeAssignmentRepository{}
	useCase := usecase.NewAssignmentUseCase(assignments, orders, technicians, sequentialID(), fixedClock())

	assignment, err := useCase.Assign(context.Background(), "order-1", "technician-1")
	if err != nil {
		t.Fatalf("assigning a free technician must be accepted: %v", err)
	}
	if !assignment.IsActive || assignment.ActiveMarker == nil || *assignment.ActiveMarker != "technician-1" {
		t.Fatalf("the active marker must carry the technician identifier, got %+v", assignment)
	}
}

func TestAssignRejectsAnUnknownOrderAndAnUnknownTechnician(t *testing.T) {
	orders := newFakeOrderRepository(buildOrder(t, "order-1", "OS-0001"))
	technicians := newFakeTechnicianRepository(buildTechnician(t, "technician-1", "user-1"))
	useCase := usecase.NewAssignmentUseCase(&fakeAssignmentRepository{}, orders, technicians, sequentialID(), fixedClock())

	if _, err := useCase.Assign(context.Background(), "order-9", "technician-1"); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("an unknown order must be rejected as not found, got %v", err)
	}
	if _, err := useCase.Assign(context.Background(), "order-1", "technician-9"); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("an unknown technician must be rejected as not found, got %v", err)
	}
}
