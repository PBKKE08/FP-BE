package common

import "github.com/google/uuid"

type City struct {
	id   uuid.UUID
	name string
}

func (c City) GetID() string {
	return c.id.String()
}

func (c City) GetName() string {
	return c.name
}
