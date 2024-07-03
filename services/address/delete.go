package address

import (
	"app/models"
	"context"
	"log"
)

func (svc *AddressSVC) Delete(ctx context.Context, userID, addressID string) (err error) {

	err = svc.UUIDInterface.IsValid(addressID)
	if err != nil {
		return models.NewErr(400, "invalid_id")
	}

	err = svc.UUIDInterface.IsValid(userID)
	if err != nil {
		return models.NewErr(400, "invalid_id")
	}

	err = svc.RepoInterface.Delete(ctx, userID, addressID)
	if err != nil {
		log.Default().Print("delete user err: ", err) // not the fast log out there, but will do the work for now
		return models.NewErr(500, "something_went_wrong")
	}

	return
}
