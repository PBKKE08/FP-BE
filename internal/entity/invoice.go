package entity

import "github.com/google/uuid"

type Invoice struct {
	ID        uuid.UUID
	BookingID uuid.UUID
	BookerID  uuid.UUID
	Open      bool
	Closed    bool
	Voided    bool
}
