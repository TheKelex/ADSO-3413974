package domain

import (
	"fmt"
	"strings"
	"time"
)

// PartUsage is one part consumed inside an intervention. It has no meaning
// outside its intervention, so it lives inside that aggregate.
type PartUsage struct {
	ID             string
	InterventionID string
	PartName       string
	Quantity       int
	CreatedAt      time.Time
}

// Intervention is one physical action executed on the vehicle inside a
// service order, with the labor hours it consumed and the parts it used.
type Intervention struct {
	ID             string
	ServiceOrderID string
	TechnicianID   string
	Description    string
	LaborHourCount float64
	PerformedAt    time.Time
	CreatedAt      time.Time
	Part           []PartUsage
}

// NewPartUsage builds a part usage after validating its quantity.
func NewPartUsage(id, interventionID, partName string, quantity int, createdAt time.Time) (PartUsage, error) {
	partName = strings.TrimSpace(partName)
	if id == "" || interventionID == "" {
		return PartUsage{}, fmt.Errorf("%w: the part usage identifiers are required", ErrInvalidInput)
	}
	if partName == "" {
		return PartUsage{}, fmt.Errorf("%w: the part name is required", ErrInvalidInput)
	}
	if quantity <= 0 {
		return PartUsage{}, fmt.Errorf("%w: the part quantity must be greater than zero", ErrInvalidInput)
	}
	if err := rejectInvalidText(partName, "part name", 160, false); err != nil {
		return PartUsage{}, err
	}
	return PartUsage{
		ID:             id,
		InterventionID: interventionID,
		PartName:       partName,
		Quantity:       quantity,
		CreatedAt:      createdAt,
	}, nil
}

// NewIntervention builds an intervention with its parts after validating the
// labor hours and every part quantity.
func NewIntervention(id, serviceOrderID, technicianID, description string, laborHourCount float64, part []PartUsage, performedAt time.Time) (Intervention, error) {
	description = strings.TrimSpace(description)
	if id == "" {
		return Intervention{}, fmt.Errorf("%w: intervention identifier is required", ErrInvalidInput)
	}
	if serviceOrderID == "" || technicianID == "" {
		return Intervention{}, fmt.Errorf("%w: the service order and the technician are required", ErrInvalidInput)
	}
	if description == "" {
		return Intervention{}, fmt.Errorf("%w: the intervention description is required", ErrInvalidInput)
	}
	if err := rejectControlCharacter(description, "intervention description"); err != nil {
		return Intervention{}, err
	}
	if laborHourCount <= 0 {
		return Intervention{}, fmt.Errorf("%w: the labor hour count must be greater than zero", ErrInvalidInput)
	}
	for _, item := range part {
		if item.Quantity <= 0 {
			return Intervention{}, fmt.Errorf("%w: the part quantity must be greater than zero", ErrInvalidInput)
		}
	}
	return Intervention{
		ID:             id,
		ServiceOrderID: serviceOrderID,
		TechnicianID:   technicianID,
		Description:    description,
		LaborHourCount: laborHourCount,
		PerformedAt:    performedAt,
		CreatedAt:      performedAt,
		Part:           part,
	}, nil
}
