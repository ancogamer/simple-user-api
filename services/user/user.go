package user

import (
	"app/models"
	"context"
)

// UserSVC users service
type UserSVC struct {
	UUIDInterface UUIDInterface
	RepoInterface RepoInterface
}

// RepoInterface Definies what is a Repo provider the user service
type RepoInterface interface {
	Create(ctx context.Context, in models.User) (err error)
	ExistID(ctx context.Context, id string) (ok bool, err error)
	Get(ctx context.Context, id string) (out models.User, err error)
	Delete(ctx context.Context, id string) (err error)
	Patch(ctx context.Context, in models.User) (err error)
}

// UUIDInterface definies what is a UUID provider for service
type UUIDInterface interface {
	NewString() (str string)        //NewString Creates new uuid as string
	IsValid(str string) (err error) //IsValid checks if the str is a valid uuid
}
