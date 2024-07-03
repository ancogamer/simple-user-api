package address

import (
	"app/models"
	"context"
	"log"
)

func (svc *AddressSVC) Patch(ctx context.Context, userID, addressID string, in models.AddressReq) (err error) {

	err = svc.UUIDInterface.IsValid(userID)
	if err != nil {
		return models.NewErr(400, "invalid_id")
	}

	if addressID != "" {
		err = svc.UUIDInterface.IsValid(addressID)
		if err != nil {
			return models.NewErr(400, "invalid_id")
		}
	}

	in.UserID = ""
	err = svc.validateUpdate(in)
	if err != nil {
		return

	}
	// checks if this addressID is from that user
	out, err := svc.RepoInterface.Get(ctx, userID, addressID)
	if err != nil {
		log.Default().Print("delete user err: ", err) // not the fast log out there, but will do the work for now
		return models.NewErr(500, "something_went_wrong")
	}
	if len(out) == 0 {
		return models.NewErr(403, "address_not_from_user")
	}

	err = svc.RepoInterface.Patch(ctx, addressID, models.Address{
		UserID:  userID,
		ID:      addressID,
		Details: in.Details,
	})
	if err != nil {
		log.Default().Print("delete user err: ", err) // not the fast log out there, but will do the work for now
		return models.NewErr(500, "something_went_wrong")
	}
	return
}

func (svc *AddressSVC) validateUpdate(in models.AddressReq) (err error) {

	if in.Details != "" {
		if len(in.Details) > 100 {
			err = models.NewErr(400, "invalid_details")
			return
		}
	}

	return
}
