package main

import "github.com/erickgreco/micro-ride/shared/types"

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	PickUp      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}
