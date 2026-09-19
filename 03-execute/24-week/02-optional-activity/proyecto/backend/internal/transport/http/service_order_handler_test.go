package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

type fakeOrderStore struct {
	order      map[string]domain.ServiceOrder
	transition []domain.StatusTransition
	writes     int
	holder     map[string]string
}

func newFakeOrderStore(order ...domain.ServiceOrder) *fakeOrderStore {
	stored := make(map[string]domain.ServiceOrder, len(order))
	for _, item := range order {
		stored[item.ID] = item
	}
	return &fakeOrderStore{order: stored, holder: make(map[string]string)}
}

func (f *fakeOrderStore) heldBy(orderID, technicianID string) {
	f.holder[orderID] = technicianID
}

func (f *fakeOrderStore) Save(_ context.Context, order domain.ServiceOrder) error {
	f.order[order.ID] = order
	return nil
}

func (f *fakeOrderStore) FindByID(_ context.Context, id string) (domain.ServiceOrder, error) {
	found, ok := f.order[id]
	if !ok {
		return domain.ServiceOrder{}, domain.ErrNotFound
	}
	return found, nil
}

func (f *fakeOrderStore) FindSummary(ctx context.Context, id string) (usecase.ServiceOrderSummary, error) {
	order, err := f.FindByID(ctx, id)
	if err != nil {
		return usecase.ServiceOrderSummary{}, err
	}
	return usecase.ServiceOrderSummary{Order: order, VehiclePlate: "ABC123", TechnicianName: "Juan Perez"}, nil
}

func (f *fakeOrderStore) List(_ context.Context, _ string) ([]usecase.ServiceOrderSummary, error) {
	listing := make([]usecase.ServiceOrderSummary, 0, len(f.order))
	for _, order := range f.order {
		listing = append(listing, usecase.ServiceOrderSummary{Order: order})
	}
	return listing, nil
}

func (f *fakeOrderStore) ListByTechnician(_ context.Context, technicianID, _ string) ([]usecase.ServiceOrderSummary, error) {
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

func (f *fakeOrderStore) ListByVehicle(_ context.Context, _ string) ([]domain.ServiceOrder, error) {
	return nil, nil
}

func (f *fakeOrderStore) UpdateStatus(_ context.Context, order domain.ServiceOrder, transition domain.StatusTransition) error {
	f.order[order.ID] = order
	f.transition = append(f.transition, transition)
	f.writes++
	return nil
}

func (f *fakeOrderStore) ListTransition(_ context.Context, _ string) ([]domain.StatusTransition, error) {
	return f.transition, nil
}

func (f *fakeOrderStore) CountByStatus(_ context.Context) (map[string]int, error) {
	return map[string]int{}, nil
}

func (f *fakeOrderStore) NextOrderNumber(_ context.Context) (string, error) {
	return "OS-0001", nil
}

type fakeVehicleStore struct {
	vehicle map[string]usecase.VehicleWithOwner
}

func newFakeVehicleStore(id string) *fakeVehicleStore {
	vehicle, _ := domain.NewVehicle(id, "customer-1", "ABC123", "VIN0001", "Mazda", "3", 2019, testMoment)
	return &fakeVehicleStore{vehicle: map[string]usecase.VehicleWithOwner{
		id: {Vehicle: vehicle, OwnerID: "customer-1", OwnerName: "Ana Gomez"},
	}}
}

func (f *fakeVehicleStore) Save(_ context.Context, _ domain.Vehicle) error { return nil }

func (f *fakeVehicleStore) List(_ context.Context) ([]usecase.VehicleWithOwner, error) {
	return nil, nil
}

func (f *fakeVehicleStore) FindByID(_ context.Context, id string) (usecase.VehicleWithOwner, error) {
	found, ok := f.vehicle[id]
	if !ok {
		return usecase.VehicleWithOwner{}, domain.ErrNotFound
	}
	return found, nil
}

type fakeAssignmentStore struct {
	assignment []domain.Assignment
}

func (f *fakeAssignmentStore) Save(_ context.Context, assignment domain.Assignment) error {
	f.assignment = append(f.assignment, assignment)
	return nil
}

func (f *fakeAssignmentStore) FindActiveByServiceOrder(_ context.Context, serviceOrderID string) (domain.Assignment, error) {
	for _, item := range f.assignment {
		if item.IsActive && item.ServiceOrderID == serviceOrderID {
			return item, nil
		}
	}
	return domain.Assignment{}, domain.ErrNotFound
}

func (f *fakeAssignmentStore) FindActiveByTechnician(_ context.Context, technicianID string) (domain.Assignment, error) {
	for _, item := range f.assignment {
		if item.IsActive && item.TechnicianID == technicianID {
			return item, nil
		}
	}
	return domain.Assignment{}, domain.ErrNotFound
}

func (f *fakeAssignmentStore) ReleaseByServiceOrder(_ context.Context, serviceOrderID string, releasedAt time.Time) error {
	for index := range f.assignment {
		if f.assignment[index].ServiceOrderID == serviceOrderID && f.assignment[index].IsActive {
			f.assignment[index].Release(releasedAt)
		}
	}
	return nil
}

func orderFixture(t *testing.T, id string, status domain.ServiceOrderStatus) domain.ServiceOrder {
	t.Helper()
	order, err := domain.NewServiceOrder(id, "OS-0001", "vehicle-1", "Ruido en el motor", testMoment)
	if err != nil {
		t.Fatalf("building the order fixture failed: %v", err)
	}
	order.Status = status
	return order
}

// asCaller runs a request through the identity the handlers expect.
func asCaller(request *http.Request, role domain.Role) *http.Request {
	return request.WithContext(context.WithValue(
		request.Context(), callerContextKey, caller{UserID: "user-1", Role: role},
	))
}

func newOrderHandler(orders *fakeOrderStore) ServiceOrderHandler {
	return newOrderHandlerWith(orders, &fakeAssignmentStore{}, &fakeTechnicianStore{})
}

func newOrderHandlerWith(orders *fakeOrderStore, assignments *fakeAssignmentStore, technicians *fakeTechnicianStore) ServiceOrderHandler {
	return NewServiceOrderHandler(usecase.NewServiceOrderUseCase(
		orders, newFakeVehicleStore("vehicle-1"), assignments, technicians,
		func() string { return "generated-id" }, testClock(),
	))
}

func TestCreateServiceOrderAnswersCreatedInReceivedStatus(t *testing.T) {
	handler := newOrderHandler(newFakeOrderStore())
	request := asCaller(httptest.NewRequest(
		http.MethodPost, "/api/service-order",
		strings.NewReader(`{"vehicleId":"vehicle-1","reportedFailure":"Ruido en el motor"}`),
	), domain.RoleAdministrator)
	recorder := httptest.NewRecorder()

	handler.Create(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("a valid check-in must answer 201, got %d body %s", recorder.Code, recorder.Body.String())
	}
	var payload serviceOrderResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("the response must be a service order payload: %v", err)
	}
	if payload.Status != string(domain.StatusReceived) {
		t.Fatalf("a new order must be RECEIVED, got %s", payload.Status)
	}
}

func TestCreateServiceOrderIsForbiddenForATechnician(t *testing.T) {
	handler := newOrderHandler(newFakeOrderStore())
	request := asCaller(httptest.NewRequest(
		http.MethodPost, "/api/service-order",
		strings.NewReader(`{"vehicleId":"vehicle-1","reportedFailure":"Ruido"}`),
	), domain.RoleTechnician)
	recorder := httptest.NewRecorder()

	handler.Create(recorder, request)

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("a technician must not open a service order, got %d", recorder.Code)
	}
}

func TestAdvanceOutOfLifecycleAnswersUnprocessableInSpanish(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	handler := newOrderHandler(orders)
	request := asCaller(httptest.NewRequest(
		http.MethodPost, "/api/service-order/order-1/status", strings.NewReader(`{"status":"READY"}`),
	), domain.RoleAdministrator)
	request.SetPathValue("serviceOrderId", "order-1")
	recorder := httptest.NewRecorder()

	handler.Advance(recorder, request)

	if recorder.Code != http.StatusUnprocessableEntity {
		t.Fatalf("an out of lifecycle move must answer 422, got %d body %s", recorder.Code, recorder.Body.String())
	}
	var payload errorPayload
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("the response must be an error payload: %v", err)
	}
	if payload.Message != "Transicion de estado no permitida." {
		t.Fatalf("the message shown to the user must be in Spanish, got %q", payload.Message)
	}
	if orders.writes != 0 {
		t.Fatalf("a rejected advance must write nothing, got %d writes", orders.writes)
	}
}

func TestAdvanceWithinTheLifecycleAnswersTheNewStatus(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	handler := newOrderHandler(orders)
	request := asCaller(httptest.NewRequest(
		http.MethodPost, "/api/service-order/order-1/status", strings.NewReader(`{"status":"IN_DIAGNOSIS"}`),
	), domain.RoleAdministrator)
	request.SetPathValue("serviceOrderId", "order-1")
	recorder := httptest.NewRecorder()

	handler.Advance(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("an allowed move must answer 200, got %d body %s", recorder.Code, recorder.Body.String())
	}
	if len(orders.transition) != 1 || orders.transition[0].ToStatus != domain.StatusInDiagnosis {
		t.Fatalf("the transition record must be written, got %+v", orders.transition)
	}
}

func TestFindServiceOrderAnswersNotFoundForAnUnknownIdentifier(t *testing.T) {
	handler := newOrderHandler(newFakeOrderStore())
	request := asCaller(httptest.NewRequest(http.MethodGet, "/api/service-order/order-9", nil), domain.RoleAdministrator)
	request.SetPathValue("serviceOrderId", "order-9")
	recorder := httptest.NewRecorder()

	handler.Find(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("an unknown order must answer 404, got %d", recorder.Code)
	}
}

func TestAdvanceIsForbiddenForATechnicianWhoDoesNotHoldTheOrder(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	technicians := newFakeTechnicianStore(t, "technician-1", "user-1")
	assignments := &fakeAssignmentStore{}
	assignment, err := domain.NewAssignment("assignment-1", "order-1", "technician-2", testMoment)
	if err != nil {
		t.Fatalf("building the assignment fixture failed: %v", err)
	}
	if err := assignments.Save(context.Background(), assignment); err != nil {
		t.Fatalf("storing the assignment fixture failed: %v", err)
	}
	handler := newOrderHandlerWith(orders, assignments, technicians)
	request := asCaller(httptest.NewRequest(
		http.MethodPost, "/api/service-order/order-1/status", strings.NewReader(`{"status":"IN_DIAGNOSIS"}`),
	), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	recorder := httptest.NewRecorder()

	handler.Advance(recorder, request)

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("a technician must not move an order they do not hold, got %d body %s", recorder.Code, recorder.Body.String())
	}
	if orders.writes != 0 {
		t.Fatalf("a forbidden advance must write nothing, got %d writes", orders.writes)
	}
}

func TestAdvanceIsAcceptedForTheAssignedTechnician(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	technicians := newFakeTechnicianStore(t, "technician-1", "user-1")
	assignments := &fakeAssignmentStore{}
	assignment, err := domain.NewAssignment("assignment-1", "order-1", "technician-1", testMoment)
	if err != nil {
		t.Fatalf("building the assignment fixture failed: %v", err)
	}
	if err := assignments.Save(context.Background(), assignment); err != nil {
		t.Fatalf("storing the assignment fixture failed: %v", err)
	}
	handler := newOrderHandlerWith(orders, assignments, technicians)
	request := asCaller(httptest.NewRequest(
		http.MethodPost, "/api/service-order/order-1/status", strings.NewReader(`{"status":"IN_DIAGNOSIS"}`),
	), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	recorder := httptest.NewRecorder()

	handler.Advance(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("the assigned technician must move the order, got %d body %s", recorder.Code, recorder.Body.String())
	}
	if len(orders.transition) != 1 || orders.transition[0].ChangedByUserID != "user-1" {
		t.Fatalf("the transition must record the assigned technician as author, got %+v", orders.transition)
	}
}

func TestAdvanceIsForbiddenForATechnicianWhenNoOneIsAssigned(t *testing.T) {
	orders := newFakeOrderStore(orderFixture(t, "order-1", domain.StatusReceived))
	technicians := newFakeTechnicianStore(t, "technician-1", "user-1")
	handler := newOrderHandlerWith(orders, &fakeAssignmentStore{}, technicians)
	request := asCaller(httptest.NewRequest(
		http.MethodPost, "/api/service-order/order-1/status", strings.NewReader(`{"status":"IN_DIAGNOSIS"}`),
	), domain.RoleTechnician)
	request.SetPathValue("serviceOrderId", "order-1")
	recorder := httptest.NewRecorder()

	handler.Advance(recorder, request)

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("an unassigned technician must not move an unassigned order, got %d", recorder.Code)
	}
}
