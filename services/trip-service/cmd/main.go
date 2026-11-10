package main

import (
	"context"
	"log"
	"time"

	"github.com/erickgreco/micro-ride/services/trip-service/internal/domain"
	"github.com/erickgreco/micro-ride/services/trip-service/internal/infrastructure/repository"
	"github.com/erickgreco/micro-ride/services/trip-service/internal/service"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()

	inMemRepo := repository.NewInmemRepository()
	srvc := service.NewService(inMemRepo)

	fare := &domain.RideFareModel{
		UserID: uuid.New(),
	}

	t, err := srvc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	for {
		time.Sleep(time.Second)
	}
}
