package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

func orderHeldBy(t *testing.T, orderID, technicianID string, assignments *fakeAssignmentStore) {
	t.Helper()
	assignment, err := domain.NewAssignment("assignment-1", orderID, technicianID, testMoment)
	if err != nil {
		t.Fatalf("building the assignment fixture failed: %v", err)
	}
	if err := assignments.Save(context.Background(), assignment); err != nil {
		t.Fatalf("storing the assignment fixture failed: %v", err)
	}
}

func newAssignmentHandlerWith(technicians *fakeTechnicianStore, orders *fakeOrderStore, assignments *fakeAssignmentStore) AssignmentHandler {
	return NewAssignmentHandler(usecase.NewAssignmentUseCase(
		assignments, orders, technicians,
		func() string { return "generated-id" }, testClock(),
	))
}

func newDiagnosticHandler(t *testing.T, technicians *fakeTechnicianStore, orders *fakeOrderStore, assignments *fakeAssignmentStore) DiagnosticHandler {
	store := &fakeDiagnosticStore{diagnostic: map[string]domain.Diagnostic{}}
	diagnostic, err := domain.NewDiagnostic("diagnostic-1", "order-1", "technician-1", "Ruido en el motor", "Bomba de agua", testMoment)
	if err != nil {
		t.Fatalf("building the diagnostic fixture failed: %v", err)
	}
	if err := store.Save(context.Background(), diagnostic); err != nil {
		t.Fatalf("storing the diagnostic fixture failed: %v", err)
	}
	return NewDiagnosticHandler(usecase.NewDiagnosticUseCase(
		store, orders, assignments, technicians,
		func() string { return "generated-id" }, testClock(),
	))
}

func newInterventionHandler(t *testing.T, technicians *fakeTechnicianStore, orders *fakeOrderStore, assignments *fakeAssignmentStore) InterventionHandler {
	return NewInterventionHandler(usecase.NewInterventionUseCase(
		newFakeInterventionStore(t, "intervention-1"),
		orders, assignments, technicians,
		func() string { return "generated-id" }, testClock(),
	))
}

func TestListReturnsOnlyTheOrdersHeldByTheTechnician(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived), orderFixture(t, "order-2", domain.StatusReceived))
	orders.heldBy("order-1", "technician-1")
	orders.heldBy("order-2", "technician-2")
	technicians := newFakeTechnicianStore(t, "technician-1", "user-1")
	assignments := &fakeAssignmentStore{}
	orderHeldBy(t, "order-1", "technician-1", assignments)
	handler := newOrderHandlerWith(orders, assignments, technicians)
	request := asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order", nil), domain.RoleTechnician)
	recorder := httptest.NewRecorder()

	handler.List(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("listing must answer 200, got %d body %s", recorder.Code, recorder.Body.String())
	}
	var payload []serviceOrderResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("the response must be an order list: %v", err)
	}
	if len(payload) != 1 || payload[0].ID != "order-1" {
		t.Fatalf("a technician must only see the orders they hold, got %+v", payload)
	}
}

func TestListReturnsEveryOrderToAdministrator(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived), orderFixture(t, "order-2", domain.StatusReceived))
	assignments := &fakeAssignmentStore{}
	handler := newOrderHandlerWith(orders, assignments, &fakeTechnicianStore{})
	request := asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order", nil), domain.RoleAdministrator)
	recorder := httptest.NewRecorder()

	handler.List(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("listing must answer 200, got %d body %s", recorder.Code, recorder.Body.String())
	}
	var payload []serviceOrderResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("the response must be an order list: %v", err)
	}
	if len(payload) != 2 {
		t.Fatalf("an administrator must see every order, got %d", len(payload))
	}
}

func TestFindIsForbiddenForATechnicianWhoDoesNotHoldTheOrder(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	technicians := newFakeTechnicianStore(t, "technician-1", "user-1")
	assignments := &fakeAssignmentStore{}
	orderHeldBy(t, "order-1", "technician-2", assignments)
	handler := newOrderHandlerWith(orders, assignments, technicians)
	request := asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order/order-1", nil), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	recorder := httptest.NewRecorder()

	handler.Find(recorder, request)

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("an unrelated technician must get 403 on the order detail, got %d", recorder.Code)
	}
}

func TestFindIsAcceptedForTheAssignedTechnician(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	technicians := newFakeTechnicianStore(t, "technician-1", "user-1")
	assignments := &fakeAssignmentStore{}
	orderHeldBy(t, "order-1", "technician-1", assignments)
	handler := newOrderHandlerWith(orders, assignments, technicians)
	request := asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order/order-1", nil), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	recorder := httptest.NewRecorder()

	handler.Find(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("the assigned technician must read the order detail, got %d body %s", recorder.Code, recorder.Body.String())
	}
}

func TestListTransitionIsForbiddenForAnUnrelatedTechnician(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	technicians := newFakeTechnicianStore(t, "technician-1", "user-1")
	assignments := &fakeAssignmentStore{}
	orderHeldBy(t, "order-1", "technician-2", assignments)
	handler := newOrderHandlerWith(orders, assignments, technicians)
	request := asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order/order-1/transition", nil), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	recorder := httptest.NewRecorder()

	handler.ListTransition(recorder, request)

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("an unrelated technician must get 403 on the status history, got %d", recorder.Code)
	}
}

func TestAssignmentFindFollowsTheOrderOwnership(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	technicians := newFakeTechnicianStore(t, "technician-1", "user-1")
	stranger := &fakeAssignmentStore{}
	orderHeldBy(t, "order-1", "technician-2", stranger)
	denied := httptest.NewRecorder()
	strangerHandler := newAssignmentHandlerWith(technicians, orders, stranger)
	request := asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order/order-1/assignment", nil), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	strangerHandler.Find(denied, request)
	if denied.Code != http.StatusForbidden {
		t.Fatalf("an unrelated technician must get 403 on the assignment, got %d", denied.Code)
	}

	owner := &fakeAssignmentStore{}
	orderHeldBy(t, "order-1", "technician-1", owner)
	allowed := httptest.NewRecorder()
	ownerHandler := newAssignmentHandlerWith(technicians, orders, owner)
	request = asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order/order-1/assignment", nil), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	ownerHandler.Find(allowed, request)
	if allowed.Code != http.StatusOK {
		t.Fatalf("the assigned technician must read the assignment, got %d body %s", allowed.Code, allowed.Body.String())
	}
}

func TestDiagnosticFindFollowsTheOrderOwnership(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	technicians := newFakeTechnicianStore(t, "technician-1", "user-1")
	stranger := &fakeAssignmentStore{}
	orderHeldBy(t, "order-1", "technician-2", stranger)
	denied := httptest.NewRecorder()
	strangerHandler := newDiagnosticHandler(t, technicians, orders, stranger)
	request := asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order/order-1/diagnostic", nil), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	strangerHandler.Find(denied, request)
	if denied.Code != http.StatusForbidden {
		t.Fatalf("an unrelated technician must get 403 on the diagnostic, got %d", denied.Code)
	}

	owner := &fakeAssignmentStore{}
	orderHeldBy(t, "order-1", "technician-1", owner)
	allowed := httptest.NewRecorder()
	ownerHandler := newDiagnosticHandler(t, technicians, orders, owner)
	request = asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order/order-1/diagnostic", nil), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	ownerHandler.Find(allowed, request)
	if allowed.Code != http.StatusOK {
		t.Fatalf("the assigned technician must read the diagnostic, got %d body %s", allowed.Code, allowed.Body.String())
	}
}

func TestInterventionListFollowsTheOrderOwnership(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	technicians := newFakeTechnicianStore(t, "technician-1", "user-1")
	stranger := &fakeAssignmentStore{}
	orderHeldBy(t, "order-1", "technician-2", stranger)
	denied := httptest.NewRecorder()
	strangerHandler := newInterventionHandler(t, technicians, orders, stranger)
	request := asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order/order-1/intervention", nil), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	strangerHandler.List(denied, request)
	if denied.Code != http.StatusForbidden {
		t.Fatalf("an unrelated technician must get 403 on the interventions, got %d", denied.Code)
	}

	owner := &fakeAssignmentStore{}
	orderHeldBy(t, "order-1", "technician-1", owner)
	allowed := httptest.NewRecorder()
	ownerHandler := newInterventionHandler(t, technicians, orders, owner)
	request = asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order/order-1/intervention", nil), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	ownerHandler.List(allowed, request)
	if allowed.Code != http.StatusOK {
		t.Fatalf("the assigned technician must read the interventions, got %d body %s", allowed.Code, allowed.Body.String())
	}
}
