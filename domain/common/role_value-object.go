package common

type Role string

const (
	Admin   Role = "admin"
	User    Role = "user"
	Partner Role = "partner"
	Unknown Role = "unknown"
)
