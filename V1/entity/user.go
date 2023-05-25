package entity

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Telephone string
	Password  string
	Gender    string
	Address   string
}
