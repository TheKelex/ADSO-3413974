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

type fakeTechnicianStore struct {
	technician map[string]domain.Technician
	byUser     map[string]string
}

func newFakeTechnicianStore(t *testing.T, id, userID string) *fakeTechnicianStore {
	t.Helper()
	technician, err := domain.NewTechnician(id, userID, "Motor", testMoment)
	if err != nil {
		t.Fatalf("building the technician fixture failed: %v", err)
	}
	return &fakeTechnicianStore{
		technician: map[string]domain.Technician{id: technician},
		byUser:     map[string]string{userID: id},
	}
}

func (f *fakeTechnicianStore) ListWorkload(_ context.Context) ([]domain.TechnicianWorkload, error) {
	return nil, nil
}

func (f *fakeTechnicianStore) FindByID(_ context.Context, id string) (domain.Technician, error) {
	found, ok := f.technician[id]
	if !ok {
		return domain.Technician{}, domain.ErrNotFound
	}
	return found, nil
}

func (f *fakeTechnicianStore) FindByUserID(_ context.Context, userID string) (domain.Technician, error) {
	id, ok := f.byUser[userID]
	if !ok {
		return domain.Technician{}, domain.ErrNotFound
	}
	return f.technician[id], nil
}

func newAssignmentHandler(t *testing.T, orders *fakeOrderStore, assignments *fakeAssignmentStore) AssignmentHandler {
	t.Helper()
	return NewAssignmentHandler(usecase.NewAssignmentUseCase(
		assignments, orders, newFakeTechnicianStore(t, "technician-1", "user-2"),
		func() string { return "generated-id" }, testClock(),
	))
}

func assignRequest(orderID, technicianID string, role domain.Role) *http.Request {
	request := asCaller(httptest.NewRequest(
		http.MethodPost, "/api/service-order/"+orderID+"/assignment",
		strings.NewReader(`{"technicianId":"`+technicianID+`"}`),
	), role)
	request.SetPathValue("serviceOrderId", orderID)
	return request
}

func TestAssignAnswersCreatedForAFreeTechnician(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	handler := newAssignmentHandler(t, orders, &fakeAssignmentStore{})
	recorder := httptest.NewRecorder()

	handler.Assign(recorder, assignRequest("order-1", "technician-1", domain.RoleAdministrator))

	if recorder.Code != http.StatusCreated {
		t.Fatalf("assigning a free technician must answer 201, got %d body %s", recorder.Code, recorder.Body.String())
	}
	var payload assignmentResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("the response must be an assignment payload: %v", err)
	}
	if !payload.IsActive || payload.TechnicianID != "technician-1" {
		t.Fatalf("the assignment must be active for the technician, got %+v", payload)
	}
}

func TestAssignABusyTechnicianAnswersConflictInSpanish(t *testing.T) {
	order := orderFixture(t, "order-1", domain.StatusReceived)
	second := orderFixture(t, "order-2", domain.StatusReceived)
	second.OrderNumber = "OS-0002"
	orders := newFakeOrderStore(order, second)
	assignments := &fakeAssignmentStore{}
	handler := newAssignmentHandler(t, orders, assignments)

	first := httptest.NewRecorder()
	handler.Assign(first, assignRequest("order-1", "technician-1", domain.RoleAdministrator))
	if first.Code != http.StatusCreated {
		t.Fatalf("the first assignment must be accepted, got %d", first.Code)
	}

	recorder := httptest.NewRecorder()
	handler.Assign(recorder, assignRequest("order-2", "technician-1", domain.RoleAdministrator))

	if recorder.Code != http.StatusConflict {
		t.Fatalf("a busy technician must answer 409, got %d body %s", recorder.Code, recorder.Body.String())
	}
	var payload errorPayload
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("the response must be an error payload: %v", err)
	}
	if payload.Message != "El tecnico ya tiene una orden activa." {
		t.Fatalf("the message shown to the user must be in Spanish, got %q", payload.Message)
	}
	if len(assignments.assignment) != 1 {
		t.Fatalf("the rejected assignment must not be stored, got %d", len(assignments.assignment))
	}
}

func TestAssignIsForbiddenForATechnician(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	handler := newAssignmentHandler(t, orders, &fakeAssignmentStore{})
	recorder := httptest.NewRecorder()

	handler.Assign(recorder, assignRequest("order-1", "technician-1", domain.RoleTechnician))

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("a technician must not assign work, got %d", recorder.Code)
	}
}

func TestFindAssignmentAnswersNotFoundWhenNobodyHoldsTheOrder(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	handler := newAssignmentHandler(t, orders, &fakeAssignmentStore{})
	request := asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order/order-1/assignment", nil), domain.RoleAdministrator)
	request.SetPathValue("serviceOrderId", "order-1")
	recorder := httptest.NewRecorder()

	handler.Find(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("an unassigned order must answer 404, got %d", recorder.Code)
	}
}
