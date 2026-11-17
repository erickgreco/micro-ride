package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erickgreco/micro-ride/services/trip-service/internal/domain"
	"github.com/erickgreco/micro-ride/shared/types"
)

type HttpHandler struct {
	Service domain.TripService
}

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	PickUp      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (h *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	t, err := h.Service.GetRoute(ctx, &reqBody.PickUp, &reqBody.Destination)
	if err != nil {
		log.Println(err)
	}

	writeJSON(w, http.StatusOK, t)

}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
