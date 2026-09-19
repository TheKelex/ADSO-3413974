package domain

import (
	"fmt"
	"time"
)

// Assignment binds one technician to one service order. ActiveMarker carries
// the technician identifier while the assignment is active and is nil once it
// is released; the storage layer makes that column unique, so a technician can
// never hold two active orders, not even under a concurrent retry.
type Assignment struct {
	ID             string
	ServiceOrderID string
	TechnicianID   string
	IsActive       bool
	ActiveMarker   *string
	AssignedAt     time.Time
	ReleasedAt     *time.Time
}

// NewAssignment creates an active assignment with its marker set.
func NewAssignment(id, serviceOrderID, technicianID string, assignedAt time.Time) (Assignment, error) {
	if id == "" {
		return Assignment{}, fmt.Errorf("%w: assignment identifier is required", ErrInvalidInput)
	}
	if serviceOrderID == "" {
		return Assignment{}, fmt.Errorf("%w: the assigned service order is required", ErrInvalidInput)
	}
	if technicianID == "" {
		return Assignment{}, fmt.Errorf("%w: the assigned technician is required", ErrInvalidInput)
	}
	marker := technicianID
	return Assignment{
		ID:             id,
		ServiceOrderID: serviceOrderID,
		TechnicianID:   technicianID,
		IsActive:       true,
		ActiveMarker:   &marker,
		AssignedAt:     assignedAt,
	}, nil
}

// Release closes the assignment and frees the technician.
func (a *Assignment) Release(at time.Time) {
	a.IsActive = false
	a.ActiveMarker = nil
	a.ReleasedAt = &at
}
