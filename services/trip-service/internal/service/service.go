package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/erickgreco/micro-ride/services/trip-service/internal/domain"
	"github.com/erickgreco/micro-ride/shared/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserID:   fare.UserID,
		Status:   "pending",
		RideFare: fare,
	}

	return s.repo.CreateTrip(ctx, t)
}

func (s *Service) GetRoute(ctx context.Context, pickUp, destination *types.Coordinate) (*types.OsrmApiResponse, error) {
	url := fmt.Sprintf(
		"http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=full&geometries=geojson",
		pickUp.Longitude, pickUp.Latitude,
		destination.Longitude, destination.Latitude,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch from OSRM API: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var routeResponse types.OsrmApiResponse

	if err := json.Unmarshal(body, &routeResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &routeResponse, nil

}
