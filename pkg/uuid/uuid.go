package uuid

import "github.com/google/uuid"

type UUID struct {
}

func (u *UUID) GetCode() string {
	code := uuid.New()
	return code.String()[0:12]
}
