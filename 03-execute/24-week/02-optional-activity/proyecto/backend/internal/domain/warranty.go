package domain

import (
	"fmt"
	"time"
)

// WarrantyKind says what the coverage protects.
type WarrantyKind string

const (
	// WarrantyLabor covers the work the technician performed.
	WarrantyLabor WarrantyKind = "LABOR"
	// WarrantyPart covers a part installed during the intervention.
	WarrantyPart WarrantyKind = "PART"
)

// Valid reports whether the kind is one the product recognizes.
func (k WarrantyKind) Valid() bool {
	return k == WarrantyLabor || k == WarrantyPart
}

// Warranty is the coverage issued over one intervention. The expiration is
// computed once, at issue time, and then persisted: validity is read from the
// stored date, never recomputed from a coverage rule that may have changed.
type Warranty struct {
	ID                 string
	InterventionID     string
	Kind               WarrantyKind
	CoverageMonthCount int
	IssuedAt           time.Time
	ExpirationDate     time.Time
}

// NewWarranty issues a warranty, deriving the expiration from the issue date
// plus the coverage in months.
func NewWarranty(id, interventionID string, kind WarrantyKind, coverageMonthCount int, issuedAt time.Time) (Warranty, error) {
	if id == "" || interventionID == "" {
		return Warranty{}, fmt.Errorf("%w: the warranty identifiers are required", ErrInvalidInput)
	}
	if !kind.Valid() {
		return Warranty{}, fmt.Errorf("%w: unknown warranty kind %q", ErrInvalidInput, kind)
	}
	if coverageMonthCount <= 0 {
		return Warranty{}, fmt.Errorf("%w: the coverage in months must be greater than zero", ErrInvalidInput)
	}
	return Warranty{
		ID:                 id,
		InterventionID:     interventionID,
		Kind:               kind,
		CoverageMonthCount: coverageMonthCount,
		IssuedAt:           issuedAt,
		ExpirationDate:     issuedAt.AddDate(0, coverageMonthCount, 0),
	}, nil
}

// IsValidAt reports whether the warranty still covers the intervention at the
// consulted date.
func (w Warranty) IsValidAt(consultedAt time.Time) bool {
	return consultedAt.Before(w.ExpirationDate)
}
