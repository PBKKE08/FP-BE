package entity

import "github.com/google/uuid"

type City struct {
	ID   uuid.UUID
	Name string
}
