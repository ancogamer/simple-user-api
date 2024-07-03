package address

import (
	"app/models"
	"context"
	"log"
)

func (svc *AddressSVC) Get(ctx context.Context, userID, addressID string) (out []models.Address, err error) {

	err = svc.UUIDInterface.IsValid(userID)
	if err != nil {
		return nil, models.NewErr(400, "invalid_id")
	}

	if addressID != "" {
		err = svc.UUIDInterface.IsValid(addressID)
		if err != nil {
			return nil, models.NewErr(400, "invalid_id")
		}
	}

	out, err = svc.RepoInterface.Get(ctx, userID, addressID)
	if err != nil {
		log.Default().Print("delete user err: ", err) // not the fast log out there, but will do the work for now
		return nil, models.NewErr(500, "something_went_wrong")
	}

	return
}
