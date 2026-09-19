package http

import (
	"net/http"

	"workshop/internal/usecase"
)

// createVehicleRequest is the payload of the vehicle form.
type createVehicleRequest struct {
	CustomerID string `json:"customerId"`
	Plate      string `json:"plate"`
	VIN        string `json:"vin"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	ModelYear  int    `json:"modelYear"`
}

// vehicleResponse is one row of the vehicle table.
type vehicleResponse struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
	OwnerName  string `json:"ownerName"`
	Plate      string `json:"plate"`
	VIN        string `json:"vin"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	ModelYear  int    `json:"modelYear"`
	CreatedAt  string `json:"createdAt"`
}

// VehicleHandler exposes the vehicle registry.
type VehicleHandler struct {
	vehicle usecase.VehicleUseCase
}

// NewVehicleHandler wires the vehicle handler.
func NewVehicleHandler(vehicle usecase.VehicleUseCase) VehicleHandler {
	return VehicleHandler{vehicle: vehicle}
}

// Create registers a vehicle for an existing customer.
func (h VehicleHandler) Create(writer http.ResponseWriter, request *http.Request) {
	if _, err := requireAdministrator(request.Context()); err != nil {
		failure(writer, err)
		return
	}
	var payload createVehicleRequest
	if err := decode(writer, request, &payload); err != nil {
		failure(writer, err)
		return
	}
	created, err := h.vehicle.Register(
		request.Context(), payload.CustomerID, payload.Plate, payload.VIN,
		payload.Brand, payload.Model, payload.ModelYear,
	)
	if err != nil {
		failure(writer, err)
		return
	}
	respond(writer, http.StatusCreated, vehicleResponse{
		ID:         created.ID,
		CustomerID: created.CustomerID,
		Plate:      created.Plate,
		VIN:        created.VIN,
		Brand:      created.Brand,
		Model:      created.Model,
		ModelYear:  created.ModelYear,
		CreatedAt:  formatTime(created.CreatedAt),
	})
}

// List returns every vehicle with the name of its owner.
func (h VehicleHandler) List(writer http.ResponseWriter, request *http.Request) {
	if _, err := requireAdministrator(request.Context()); err != nil {
		failure(writer, err)
		return
	}
	listed, err := h.vehicle.List(request.Context())
	if err != nil {
		failure(writer, err)
		return
	}
	payload := make([]vehicleResponse, 0, len(listed))
	for _, item := range listed {
		payload = append(payload, toVehicleResponse(item))
	}
	respond(writer, http.StatusOK, payload)
}

func toVehicleResponse(item usecase.VehicleWithOwner) vehicleResponse {
	return vehicleResponse{
		ID:         item.Vehicle.ID,
		CustomerID: item.OwnerID,
		OwnerName:  item.OwnerName,
		Plate:      item.Vehicle.Plate,
		VIN:        item.Vehicle.VIN,
		Brand:      item.Vehicle.Brand,
		Model:      item.Vehicle.Model,
		ModelYear:  item.Vehicle.ModelYear,
		CreatedAt:  formatTime(item.Vehicle.CreatedAt),
	}
}
