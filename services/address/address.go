package address

import (
	"app/models"
	"context"
)

// Address users service
type AddressSVC struct {
	UUIDInterface UUIDInterface
	RepoInterface RepoInterface
}

// RepoInterface Definies what is a Repo provider the user service
type RepoInterface interface {
	Create(ctx context.Context, in models.Address) (err error)                           // Creates a address for a user
	Get(ctx context.Context, userID, addressID string) (out []models.Address, err error) // Gets address from user, or get 1 address, in that case the slice will only have 1 obj
	Delete(ctx context.Context, userID, addressID string) (err error)                    //  Deletes a user address based on userID and address ID
	Patch(ctx context.Context, addressID string, in models.Address) (err error)          // Patch patchs a address info
}

// UUIDInterface definies what is a UUID provider for service
type UUIDInterface interface {
	NewString() (str string)        //NewString Creates new uuid as string
	IsValid(str string) (err error) //IsValid checks if the str is a valid uuid
}
