package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	Username string
	Password string
	Gender   string
	Address  string
}
