package http

import (
	"net/http"

	"workshop/internal/usecase"
)

// timelineEntryResponse is one dated fact of the clinical history.
type timelineEntryResponse struct {
	Kind        string `json:"kind"`
	OccurredAt  string `json:"occurredAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Reference   string `json:"reference"`
}

// timelineResponse is the clinical history of one vehicle.
type timelineResponse struct {
	Vehicle vehicleResponse         `json:"vehicle"`
	Entry   []timelineEntryResponse `json:"entry"`
}

// TimelineHandler exposes the clinical history of a vehicle.
type TimelineHandler struct {
	timeline usecase.TimelineUseCase
}

// NewTimelineHandler wires the timeline handler.
func NewTimelineHandler(timeline usecase.TimelineUseCase) TimelineHandler {
	return TimelineHandler{timeline: timeline}
}

// Build returns every fact recorded for a vehicle, oldest first.
func (h TimelineHandler) Build(writer http.ResponseWriter, request *http.Request) {
	if _, err := requireAdministrator(request.Context()); err != nil {
		failure(writer, err)
		return
	}
	timeline, err := h.timeline.Build(request.Context(), request.PathValue("vehicleId"))
	if err != nil {
		failure(writer, err)
		return
	}
	entry := make([]timelineEntryResponse, 0, len(timeline.Entry))
	for _, item := range timeline.Entry {
		entry = append(entry, timelineEntryResponse{
			Kind:        string(item.Kind),
			OccurredAt:  formatTime(item.OccurredAt),
			Title:       item.Title,
			Description: item.Description,
			Reference:   item.Reference,
		})
	}
	respond(writer, http.StatusOK, timelineResponse{
		Vehicle: toVehicleResponse(timeline.Vehicle),
		Entry:   entry,
	})
}
