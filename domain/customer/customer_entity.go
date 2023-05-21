package customer

import (
	"github.com/PBKKE08/FP-BE/domain/common"
	"github.com/google/uuid"
)

type Customer struct {
	id        uuid.UUID
	email     string // bisa jadi ini value object
	telephone common.TelephoneNumber
	gender    common.Gender
	city      common.City
	role      common.Role
}
