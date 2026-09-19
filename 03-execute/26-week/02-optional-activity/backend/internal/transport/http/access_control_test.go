package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

type fakeCustomerStore struct {
	customer map[string]domain.Customer
}

func (f *fakeCustomerStore) Save(_ context.Context, customer domain.Customer) error {
	f.customer[customer.ID] = customer
	return nil
}

func (f *fakeCustomerStore) List(_ context.Context) ([]domain.Customer, error) {
	return nil, nil
}

func (f *fakeCustomerStore) FindByID(_ context.Context, _ string) (domain.Customer, error) {
	return domain.Customer{}, domain.ErrNotFound
}

type fakeDiagnosticStore struct {
	diagnostic map[string]domain.Diagnostic
}

func (f *fakeDiagnosticStore) Save(_ context.Context, diagnostic domain.Diagnostic) error {
	f.diagnostic[diagnostic.ID] = diagnostic
	return nil
}

func (f *fakeDiagnosticStore) FindByServiceOrder(_ context.Context, serviceOrderID string) (domain.Diagnostic, error) {
	for _, item := range f.diagnostic {
		if item.ServiceOrderID == serviceOrderID {
			return item, nil
		}
	}
	return domain.Diagnostic{}, domain.ErrNotFound
}

func (f *fakeDiagnosticStore) ListByVehicle(_ context.Context, _ string) ([]domain.Diagnostic, error) {
	return nil, nil
}

// assertListAccess checks that a list endpoint answers 403 to a technician and
// 200 to an administrator, as required by A1.
func assertListAccess(t *testing.T, list func(http.ResponseWriter, *http.Request), path string) {
	t.Helper()
	denied := httptest.NewRecorder()
	list(denied, asCaller(httptest.NewRequest(http.MethodGet, path, nil), domain.RoleTechnician))
	if denied.Code != http.StatusForbidden {
		t.Fatalf("a technician must get 403 on %s, got %d", path, denied.Code)
	}
	allowed := httptest.NewRecorder()
	list(allowed, asCaller(httptest.NewRequest(http.MethodGet, path, nil), domain.RoleAdministrator))
	if allowed.Code != http.StatusOK {
		t.Fatalf("an administrator must get 200 on %s, got %d body %s", path, allowed.Code, allowed.Body.String())
	}
}

func TestCustomerListRequiresAdministrator(t *testing.T) {
	customers := &fakeCustomerStore{customer: map[string]domain.Customer{}}
	handler := NewCustomerHandler(usecase.NewCustomerUseCase(customers, func() string { return "generated-id" }, testClock()))
	assertListAccess(t, handler.List, "/api/customer")
}

func TestVehicleListRequiresAdministrator(t *testing.T) {
	handler := NewVehicleHandler(usecase.NewVehicleUseCase(
		newFakeVehicleStore("vehicle-1"), &fakeCustomerStore{customer: map[string]domain.Customer{}}, func() string { return "generated-id" }, testClock(),
	))
	assertListAccess(t, handler.List, "/api/vehicle")
}

func TestTechnicianListRequiresAdministrator(t *testing.T) {
	handler := NewTechnicianHandler(usecase.NewTechnicianUseCase(&fakeTechnicianStore{}))
	assertListAccess(t, handler.List, "/api/technician")
}

func TestWarrantyListRequiresAdministrator(t *testing.T) {
	handler := newWarrantyHandler(t, newFakeWarrantyStore())
	assertListAccess(t, handler.List, "/api/warranty")
}

func TestTimelineBuildRequiresAdministrator(t *testing.T) {
	handler := NewTimelineHandler(usecase.NewTimelineUseCase(
		newFakeVehicleStore("vehicle-1"),
		newFakeOrderStore(),
		&fakeDiagnosticStore{diagnostic: map[string]domain.Diagnostic{}},
		newFakeInterventionStore(t, "intervention-1"),
		newFakeWarrantyStore(),
	))
	beat := func(role domain.Role) *httptest.ResponseRecorder {
		request := asCaller(httptest.NewRequest(http.MethodGet, "/api/vehicle/vehicle-1/timeline", nil), role)
		request.SetPathValue("vehicleId", "vehicle-1")
		recorder := httptest.NewRecorder()
		handler.Build(recorder, request)
		return recorder
	}

	if recorder := beat(domain.RoleTechnician); recorder.Code != http.StatusForbidden {
		t.Fatalf("a technician must get 403 on the timeline, got %d", recorder.Code)
	}
	if recorder := beat(domain.RoleAdministrator); recorder.Code != http.StatusOK {
		t.Fatalf("an administrator must get 200 on the timeline, got %d body %s", recorder.Code, recorder.Body.String())
	}
}
