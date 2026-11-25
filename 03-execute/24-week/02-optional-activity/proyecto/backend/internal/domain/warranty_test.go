package domain_test

import (
	"errors"
	"testing"
	"time"

	"workshop/internal/domain"
)

func TestWarrantyExpirationIsTheIssueDatePlusTheCoverage(t *testing.T) {
	issuedAt := time.Date(2026, time.January, 15, 10, 0, 0, 0, time.UTC)
	warranty, err := domain.NewWarranty("warranty-1", "intervention-1", domain.WarrantyLabor, 12, issuedAt)
	if err != nil {
		t.Fatalf("issuing a valid warranty failed: %v", err)
	}
	expected := time.Date(2027, time.January, 15, 10, 0, 0, 0, time.UTC)
	if !warranty.ExpirationDate.Equal(expected) {
		t.Fatalf("expected the expiration at %s, got %s", expected, warranty.ExpirationDate)
	}
}

func TestWarrantyValidityFlipsAtTheExpirationDate(t *testing.T) {
	issuedAt := time.Date(2026, time.January, 15, 10, 0, 0, 0, time.UTC)
	warranty, err := domain.NewWarranty("warranty-1", "intervention-1", domain.WarrantyPart, 6, issuedAt)
	if err != nil {
		t.Fatalf("issuing a valid warranty failed: %v", err)
	}
	if !warranty.IsValidAt(issuedAt.AddDate(0, 3, 0)) {
		t.Fatal("a warranty must be valid three months into a six month coverage")
	}
	if warranty.IsValidAt(issuedAt.AddDate(0, 9, 0)) {
		t.Fatal("a warranty must be expired nine months into a six month coverage")
	}
	if warranty.IsValidAt(warranty.ExpirationDate) {
		t.Fatal("a warranty is no longer valid exactly at its expiration date")
	}
}

func TestWarrantyRejectsAnUnknownKindAndANonPositiveCoverage(t *testing.T) {
	issuedAt := time.Now()
	if _, err := domain.NewWarranty("warranty-1", "intervention-1", domain.WarrantyKind("OTHER"), 12, issuedAt); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("an unknown warranty kind must be rejected, got %v", err)
	}
	if _, err := domain.NewWarranty("warranty-1", "intervention-1", domain.WarrantyLabor, 0, issuedAt); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("a coverage of zero months must be rejected, got %v", err)
	}
}
