package domain

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            uuid.UUID
	PackageSlug       string // ex: van, luxury, sedan
	TotalPriceInCents float64
	ExpiresAt         time.Time
}
