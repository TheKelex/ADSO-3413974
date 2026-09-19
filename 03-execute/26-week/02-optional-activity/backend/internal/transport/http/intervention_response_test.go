package http

import (
	"testing"
	"time"

	"workshop/internal/domain"
	"workshop/internal/usecase"
)

func TestToInterventionResponseIncludesWarrantyState(t *testing.T) {
	intervention := domain.Intervention{
		ID:             "intervention-1",
		ServiceOrderID: "order-1",
		TechnicianID:   "technician-1",
		Description:    "Cambio de bujias",
		LaborHourCount: 1.5,
		PerformedAt:    time.Date(2026, 9, 18, 10, 0, 0, 0, time.UTC),
	}
	kind := domain.WarrantyLabor
	valid := true
	response := toInterventionResponse(usecase.InterventionView{
		Intervention:  intervention,
		WarrantyID:    stringPointer("warranty-1"),
		WarrantyKind:  &kind,
		WarrantyValid: &valid,
	})

	if response.WarrantyID == nil || *response.WarrantyID != "warranty-1" {
		t.Fatalf("expected the warranty id, got %v", response.WarrantyID)
	}
	if response.WarrantyKind == nil || *response.WarrantyKind != "LABOR" {
		t.Fatalf("expected the warranty kind, got %v", response.WarrantyKind)
	}
	if response.WarrantyValid == nil || !*response.WarrantyValid {
		t.Fatalf("expected a valid warranty, got %v", response.WarrantyValid)
	}
}

func TestToInterventionResponseLeavesMissingWarrantyEmpty(t *testing.T) {
	response := toInterventionResponse(usecase.InterventionView{
		Intervention: domain.Intervention{ID: "intervention-1"},
	})

	if response.WarrantyID != nil || response.WarrantyKind != nil || response.WarrantyValid != nil {
		t.Fatalf("expected no warranty fields for an uncovered intervention, got %#v", response)
	}
}

func stringPointer(value string) *string {
	return &value
}
