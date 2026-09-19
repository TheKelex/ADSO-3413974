package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

type fakeInterventionStore struct {
	intervention map[string]domain.Intervention
}

func newFakeInterventionStore(t *testing.T, id string) *fakeInterventionStore {
	t.Helper()
	intervention, err := domain.NewIntervention(id, "order-1", "technician-1", "Cambio de bujias", 1.5, nil, testMoment)
	if err != nil {
		t.Fatalf("building the intervention fixture failed: %v", err)
	}
	return &fakeInterventionStore{intervention: map[string]domain.Intervention{id: intervention}}
}

func (f *fakeInterventionStore) Save(_ context.Context, intervention domain.Intervention) error {
	f.intervention[intervention.ID] = intervention
	return nil
}

func (f *fakeInterventionStore) FindByID(_ context.Context, id string) (domain.Intervention, error) {
	found, ok := f.intervention[id]
	if !ok {
		return domain.Intervention{}, domain.ErrNotFound
	}
	return found, nil
}

func (f *fakeInterventionStore) ListByServiceOrder(_ context.Context, _ string) ([]domain.Intervention, error) {
	return nil, nil
}

func (f *fakeInterventionStore) ListByVehicle(_ context.Context, _ string) ([]domain.Intervention, error) {
	return nil, nil
}

type fakeWarrantyStore struct {
	warranty map[string]domain.Warranty
}

func newFakeWarrantyStore() *fakeWarrantyStore {
	return &fakeWarrantyStore{warranty: make(map[string]domain.Warranty)}
}

func (f *fakeWarrantyStore) Save(_ context.Context, warranty domain.Warranty) error {
	f.warranty[warranty.ID] = warranty
	return nil
}

func (f *fakeWarrantyStore) FindByID(_ context.Context, id string) (domain.Warranty, error) {
	found, ok := f.warranty[id]
	if !ok {
		return domain.Warranty{}, domain.ErrNotFound
	}
	return found, nil
}

func (f *fakeWarrantyStore) List(_ context.Context) ([]usecase.WarrantyView, error) {
	view := make([]usecase.WarrantyView, 0, len(f.warranty))
	for _, item := range f.warranty {
		view = append(view, usecase.WarrantyView{Warranty: item, OrderNumber: "OS-0001", VehiclePlate: "ABC123"})
	}
	return view, nil
}

func (f *fakeWarrantyStore) ListByVehicle(_ context.Context, _ string) ([]domain.Warranty, error) {
	return nil, nil
}

func newWarrantyHandler(t *testing.T, warranties *fakeWarrantyStore) WarrantyHandler {
	t.Helper()
	return NewWarrantyHandler(usecase.NewWarrantyUseCase(
		warranties, newFakeInterventionStore(t, "intervention-1"),
		func() string { return "warranty-1" }, testClock(),
	))
}

func issueWarranty(t *testing.T, handler WarrantyHandler, body string, role domain.Role) *httptest.ResponseRecorder {
	t.Helper()
	request := asCaller(httptest.NewRequest(http.MethodPost, "/api/warranty", strings.NewReader(body)), role)
	recorder := httptest.NewRecorder()
	handler.Issue(recorder, request)
	return recorder
}

func TestIssueWarrantyAnswersTheComputedExpiration(t *testing.T) {
	handler := newWarrantyHandler(t, newFakeWarrantyStore())
	recorder := issueWarranty(t,
		handler, `{"interventionId":"intervention-1","kind":"LABOR","coverageMonthCount":12}`, domain.RoleAdministrator,
	)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("issuing a warranty must answer 201, got %d body %s", recorder.Code, recorder.Body.String())
	}
	var payload warrantyResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("the response must be a warranty payload: %v", err)
	}
	expected := testMoment.AddDate(0, 12, 0).UTC().Format("2006-01-02T15:04:05Z")
	if payload.ExpirationDate != expected {
		t.Fatalf("expected the expiration at %s, got %s", expected, payload.ExpirationDate)
	}
}

func TestIssueWarrantyRejectsAnUnknownKindInSpanish(t *testing.T) {
	handler := newWarrantyHandler(t, newFakeWarrantyStore())
	recorder := issueWarranty(t,
		handler, `{"interventionId":"intervention-1","kind":"OTHER","coverageMonthCount":12}`, domain.RoleAdministrator,
	)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("an unknown warranty kind must answer 400, got %d", recorder.Code)
	}
	var payload errorPayload
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("the response must be an error payload: %v", err)
	}
	if payload.Message == "" || strings.Contains(payload.Message, "unknown") {
		t.Fatalf("the message must be sanitized and in Spanish, got %q", payload.Message)
	}
}

func TestIssueWarrantyIsForbiddenForATechnician(t *testing.T) {
	handler := newWarrantyHandler(t, newFakeWarrantyStore())
	recorder := issueWarranty(t,
		handler, `{"interventionId":"intervention-1","kind":"LABOR","coverageMonthCount":12}`, domain.RoleTechnician,
	)

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("a technician must not issue a warranty, got %d", recorder.Code)
	}
}

func TestListWarrantyReportsValidityAtTheConsultedDate(t *testing.T) {
	warranties := newFakeWarrantyStore()
	handler := newWarrantyHandler(t, warranties)
	if code := issueWarranty(t,
		handler, `{"interventionId":"intervention-1","kind":"PART","coverageMonthCount":6}`, domain.RoleAdministrator,
	).Code; code != http.StatusCreated {
		t.Fatalf("issuing the warranty fixture failed with %d", code)
	}

	valid := listWarrantyAt(t, handler, testMoment.AddDate(0, 1, 0).Format("2006-01-02"))
	if len(valid) != 1 || !valid[0].Valid {
		t.Fatalf("the warranty must read as valid one month in, got %+v", valid)
	}
	expired := listWarrantyAt(t, handler, testMoment.AddDate(1, 0, 0).Format("2006-01-02"))
	if len(expired) != 1 || expired[0].Valid {
		t.Fatalf("the warranty must read as expired one year in, got %+v", expired)
	}
}

func listWarrantyAt(t *testing.T, handler WarrantyHandler, consultedAt string) []warrantyResponse {
	t.Helper()
	request := asCaller(
		httptest.NewRequest(http.MethodGet, "/api/warranty?consultedAt="+consultedAt, nil), domain.RoleAdministrator,
	)
	recorder := httptest.NewRecorder()
	handler.List(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("listing warranties must answer 200, got %d body %s", recorder.Code, recorder.Body.String())
	}
	var payload []warrantyResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("the response must be a warranty list: %v", err)
	}
	return payload
}
