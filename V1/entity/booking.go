package entity

import (
	"github.com/google/uuid"
	"time"
)

type Booking struct {
	ID         uuid.UUID
	BookerID   uuid.UUID
	PartnerID  uuid.UUID
	Schedule   time.Time
	Duration   time.Duration
	CityID     uuid.UUID
	CategoryID uuid.UUID
	Open       bool
	Closed     bool
	Completed  bool
	Voided     bool
}
