package utils

import "github.com/google/uuid"

type UUIDSvc struct{}

// NewString generates a new uuid as str, this is a mustDo or panic
func (svc *UUIDSvc) NewString() (str string) {
	return uuid.New().String()
}

// IsValid checks if the str is a valid uuid
func (svc *UUIDSvc) IsValid(str string) (err error) {
	return uuid.Validate(str)
}
