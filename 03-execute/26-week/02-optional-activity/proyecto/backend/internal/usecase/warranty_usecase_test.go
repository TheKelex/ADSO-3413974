package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

type fakeInterventionRepository struct {
	intervention map[string]domain.Intervention
}

func newFakeInterventionRepository(t *testing.T, id string) *fakeInterventionRepository {
	t.Helper()
	intervention, err := domain.NewIntervention(id, "order-1", "technician-1", "Cambio de bujias", 1.5, nil, fixedClock()())
	if err != nil {
		t.Fatalf("building the intervention fixture failed: %v", err)
	}
	return &fakeInterventionRepository{intervention: map[string]domain.Intervention{id: intervention}}
}

func (f *fakeInterventionRepository) Save(_ context.Context, intervention domain.Intervention) error {
	f.intervention[intervention.ID] = intervention
	return nil
}

func (f *fakeInterventionRepository) FindByID(_ context.Context, id string) (domain.Intervention, error) {
	found, ok := f.intervention[id]
	if !ok {
		return domain.Intervention{}, domain.ErrNotFound
	}
	return found, nil
}

func (f *fakeInterventionRepository) ListByServiceOrder(_ context.Context, _ string) ([]domain.Intervention, error) {
	return nil, nil
}

func (f *fakeInterventionRepository) ListByVehicle(_ context.Context, _ string) ([]domain.Intervention, error) {
	return nil, nil
}

type fakeWarrantyRepository struct {
	warranty map[string]domain.Warranty
}

func newFakeWarrantyRepository() *fakeWarrantyRepository {
	return &fakeWarrantyRepository{warranty: make(map[string]domain.Warranty)}
}

func (f *fakeWarrantyRepository) Save(_ context.Context, warranty domain.Warranty) error {
	f.warranty[warranty.ID] = warranty
	return nil
}

func (f *fakeWarrantyRepository) FindByID(_ context.Context, id string) (domain.Warranty, error) {
	found, ok := f.warranty[id]
	if !ok {
		return domain.Warranty{}, domain.ErrNotFound
	}
	return found, nil
}

func (f *fakeWarrantyRepository) List(_ context.Context) ([]usecase.WarrantyView, error) {
	view := make([]usecase.WarrantyView, 0, len(f.warranty))
	for _, item := range f.warranty {
		view = append(view, usecase.WarrantyView{Warranty: item, OrderNumber: "OS-0001", VehiclePlate: "ABC123"})
	}
	return view, nil
}

func (f *fakeWarrantyRepository) ListByVehicle(_ context.Context, _ string) ([]domain.Warranty, error) {
	return nil, nil
}

func TestIssueStoresTheComputedExpiration(t *testing.T) {
	warranties := newFakeWarrantyRepository()
	useCase := usecase.NewWarrantyUseCase(
		warranties, newFakeInterventionRepository(t, "intervention-1"), sequentialID(), fixedClock(),
	)

	warranty, err := useCase.Issue(context.Background(), "intervention-1", domain.WarrantyLabor, 12)
	if err != nil {
		t.Fatalf("issuing a warranty over a known intervention must be accepted: %v", err)
	}
	expected := fixedClock()().AddDate(0, 12, 0)
	if !warranty.ExpirationDate.Equal(expected) {
		t.Fatalf("expected the expiration at %s, got %s", expected, warranty.ExpirationDate)
	}
	if _, ok := warranties.warranty[warranty.ID]; !ok {
		t.Fatal("the issued warranty must be stored")
	}
}

func TestIssueRejectsAnUnknownIntervention(t *testing.T) {
	useCase := usecase.NewWarrantyUseCase(
		newFakeWarrantyRepository(), newFakeInterventionRepository(t, "intervention-1"), sequentialID(), fixedClock(),
	)
	if _, err := useCase.Issue(context.Background(), "intervention-9", domain.WarrantyPart, 6); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("an unknown intervention must be rejected as not found, got %v", err)
	}
}

func TestValidityAtFlipsAroundTheExpirationDate(t *testing.T) {
	warranties := newFakeWarrantyRepository()
	useCase := usecase.NewWarrantyUseCase(
		warranties, newFakeInterventionRepository(t, "intervention-1"), sequentialID(), fixedClock(),
	)
	issued, err := useCase.Issue(context.Background(), "intervention-1", domain.WarrantyLabor, 6)
	if err != nil {
		t.Fatalf("issuing the warranty fixture failed: %v", err)
	}

	_, validBefore, err := useCase.ValidityAt(context.Background(), issued.ID, fixedClock()().AddDate(0, 3, 0))
	if err != nil || !validBefore {
		t.Fatalf("the warranty must be valid three months in, got valid=%v err=%v", validBefore, err)
	}
	_, validAfter, err := useCase.ValidityAt(context.Background(), issued.ID, fixedClock()().AddDate(0, 9, 0))
	if err != nil || validAfter {
		t.Fatalf("the warranty must be expired nine months in, got valid=%v err=%v", validAfter, err)
	}
}

func TestListEvaluatesValidityAtTheConsultedDate(t *testing.T) {
	warranties := newFakeWarrantyRepository()
	useCase := usecase.NewWarrantyUseCase(
		warranties, newFakeInterventionRepository(t, "intervention-1"), sequentialID(), fixedClock(),
	)
	if _, err := useCase.Issue(context.Background(), "intervention-1", domain.WarrantyPart, 6); err != nil {
		t.Fatalf("issuing the warranty fixture failed: %v", err)
	}

	view, err := useCase.List(context.Background(), fixedClock()().AddDate(0, 1, 0))
	if err != nil {
		t.Fatalf("listing warranties failed: %v", err)
	}
	if len(view) != 1 || !view[0].Valid {
		t.Fatalf("the warranty must read as valid one month in, got %+v", view)
	}
	expiredView, err := useCase.List(context.Background(), fixedClock()().AddDate(1, 0, 0))
	if err != nil {
		t.Fatalf("listing warranties failed: %v", err)
	}
	if len(expiredView) != 1 || expiredView[0].Valid {
		t.Fatalf("the warranty must read as expired one year in, got %+v", expiredView)
	}
}

func TestListWithoutADateUsesTheCurrentMoment(t *testing.T) {
	warranties := newFakeWarrantyRepository()
	useCase := usecase.NewWarrantyUseCase(
		warranties, newFakeInterventionRepository(t, "intervention-1"), sequentialID(), fixedClock(),
	)
	if _, err := useCase.Issue(context.Background(), "intervention-1", domain.WarrantyLabor, 12); err != nil {
		t.Fatalf("issuing the warranty fixture failed: %v", err)
	}
	view, err := useCase.List(context.Background(), time.Time{})
	if err != nil {
		t.Fatalf("listing warranties failed: %v", err)
	}
	if len(view) != 1 || !view[0].Valid {
		t.Fatalf("a freshly issued warranty must read as valid now, got %+v", view)
	}
}
