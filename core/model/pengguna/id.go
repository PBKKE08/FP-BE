package pengguna

import "github.com/google/uuid"

type ID string

func NewID() ID {
	id := uuid.NewString()
	return ID(id)
}

func NewIDFrom(id string) (ID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return ID(id), nil
}

func (id ID) String() string {
	return string(id)
}