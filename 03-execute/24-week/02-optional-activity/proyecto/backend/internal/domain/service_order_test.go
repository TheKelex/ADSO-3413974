package domain_test

import (
	"errors"
	"testing"
	"time"

	"workshop/internal/domain"
)

func openOrder(t *testing.T) domain.ServiceOrder {
	t.Helper()
	order, err := domain.NewServiceOrder("order-1", "OS-0001", "vehicle-1", "Ruido en el motor", time.Now())
	if err != nil {
		t.Fatalf("opening a valid service order failed: %v", err)
	}
	if order.Status != domain.StatusReceived {
		t.Fatalf("a new service order must start in RECEIVED, got %s", order.Status)
	}
	return order
}

func TestServiceOrderWalksTheWholeLifecycle(t *testing.T) {
	order := openOrder(t)
	step := []domain.ServiceOrderStatus{
		domain.StatusInDiagnosis,
		domain.StatusInRepair,
		domain.StatusReady,
		domain.StatusDelivered,
	}
	for index, next := range step {
		transition, err := order.MoveTo(next, "transition-1", "user-1", time.Now())
		if err != nil {
			t.Fatalf("step %d to %s must be allowed: %v", index, next, err)
		}
		if transition.ToStatus != next || transition.ServiceOrderID != order.ID {
			t.Fatalf("the transition record does not describe the change: %+v", transition)
		}
		if order.Status != next {
			t.Fatalf("the order status is %s after moving to %s", order.Status, next)
		}
	}
}

func TestServiceOrderRejectsAnOutOfLifecycleMove(t *testing.T) {
	order := openOrder(t)
	if _, err := order.MoveTo(domain.StatusReady, "transition-1", "user-1", time.Now()); err == nil {
		t.Fatal("moving from RECEIVED to READY must be rejected")
	} else if !errors.Is(err, domain.ErrInvalidTransition) {
		t.Fatalf("the rejection must be an invalid transition, got %v", err)
	}
	if order.Status != domain.StatusReceived {
		t.Fatalf("a rejected move must leave the order untouched, got %s", order.Status)
	}
}

func TestServiceOrderRejectsAnUnknownStatus(t *testing.T) {
	order := openOrder(t)
	if _, err := order.MoveTo(domain.ServiceOrderStatus("CANCELLED"), "transition-1", "user-1", time.Now()); !errors.Is(err, domain.ErrInvalidTransition) {
		t.Fatalf("an unknown status must be rejected as an invalid transition, got %v", err)
	}
}

func TestDeliveredIsTerminalAndClosesTheOrder(t *testing.T) {
	if domain.StatusDelivered.IsOpen() {
		t.Fatal("a delivered order no longer occupies the workshop")
	}
	if !domain.StatusInRepair.IsOpen() {
		t.Fatal("an order in repair is still open")
	}
	if domain.StatusDelivered.CanMoveTo(domain.StatusReceived) {
		t.Fatal("DELIVERED is terminal and must not move back")
	}
}

func TestServiceOrderRequiresItsReportedFailure(t *testing.T) {
	if _, err := domain.NewServiceOrder("order-1", "OS-0001", "vehicle-1", "   ", time.Now()); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("an empty reported failure must be rejected, got %v", err)
	}
}
