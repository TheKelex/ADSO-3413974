package domain

import (
	"fmt"
	"strings"
	"time"
)

// Diagnostic is the technical evaluation of a service order, written by the
// technician who holds its active assignment.
type Diagnostic struct {
	ID                string
	ServiceOrderID    string
	TechnicianID      string
	Finding           string
	ComponentToRepair string
	CreatedAt         time.Time
}

// NewDiagnostic builds a diagnostic after validating its content.
func NewDiagnostic(id, serviceOrderID, technicianID, finding, componentToRepair string, createdAt time.Time) (Diagnostic, error) {
	finding = strings.TrimSpace(finding)
	componentToRepair = strings.TrimSpace(componentToRepair)
	if id == "" {
		return Diagnostic{}, fmt.Errorf("%w: diagnostic identifier is required", ErrInvalidInput)
	}
	if serviceOrderID == "" || technicianID == "" {
		return Diagnostic{}, fmt.Errorf("%w: the service order and the technician are required", ErrInvalidInput)
	}
	if finding == "" {
		return Diagnostic{}, fmt.Errorf("%w: the diagnostic finding is required", ErrInvalidInput)
	}
	if componentToRepair == "" {
		return Diagnostic{}, fmt.Errorf("%w: the component to repair is required", ErrInvalidInput)
	}
	if err := rejectControlCharacter(finding, "diagnostic finding"); err != nil {
		return Diagnostic{}, err
	}
	if err := rejectControlCharacter(componentToRepair, "component to repair"); err != nil {
		return Diagnostic{}, err
	}
	return Diagnostic{
		ID:                id,
		ServiceOrderID:    serviceOrderID,
		TechnicianID:      technicianID,
		Finding:           finding,
		ComponentToRepair: componentToRepair,
		CreatedAt:         createdAt,
	}, nil
}
