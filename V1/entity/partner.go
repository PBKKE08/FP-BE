package entity

import "github.com/google/uuid"

type Partner struct {
	user_id   	uuid.UUID
	price		float32
	rating 		float32
}
