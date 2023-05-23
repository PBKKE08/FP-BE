package entity

import "github.com/google/uuid"

type Review struct {
	ID          uuid.UUID
	BookingID   uuid.UUID
	Rating      int
	Description string
}
