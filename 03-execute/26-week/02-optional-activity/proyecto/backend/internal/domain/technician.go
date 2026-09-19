package domain

import (
	"fmt"
	"strings"
	"time"
)

// Technician is the mechanic profile bound to exactly one user account.
type Technician struct {
	ID        string
	UserID    string
	Specialty string
	CreatedAt time.Time
}

// NewTechnician builds a technician profile after validating its data.
func NewTechnician(id, userID, specialty string, createdAt time.Time) (Technician, error) {
	specialty = strings.TrimSpace(specialty)
	if id == "" {
		return Technician{}, fmt.Errorf("%w: technician identifier is required", ErrInvalidInput)
	}
	if userID == "" {
		return Technician{}, fmt.Errorf("%w: the technician user account is required", ErrInvalidInput)
	}
	if specialty == "" {
		return Technician{}, fmt.Errorf("%w: technician specialty is required", ErrInvalidInput)
	}
	return Technician{ID: id, UserID: userID, Specialty: specialty, CreatedAt: createdAt}, nil
}

// TechnicianWorkload is the read model the allocation panel and the dashboard
// consume: who the technician is and whether they already hold an active order.
type TechnicianWorkload struct {
	Technician         Technician
	FullName           string
	Busy               bool
	ActiveOrderID      string
	ActiveOrderNumber  string
	ActiveVehiclePlate string
}
