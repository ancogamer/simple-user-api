package address

import (
	"app/models"
	"context"
	"log"
)

func (svc *AddressSVC) Create(ctx context.Context, in models.AddressReq) (out models.Address, err error) {

	err = svc.validateCreate(in)
	if err != nil {
		return
	}
	out = models.Address{
		UserID:  in.UserID,
		ID:      svc.UUIDInterface.NewString(),
		Details: in.Details,
		ZipCode: in.ZipCode,
		Country: in.Country,
		State:   in.State,
		City:    in.City,
	}
	err = svc.RepoInterface.Create(ctx, out)
	if err != nil {
		log.Default().Print("delete user err: ", err) // not the fast log out there, but will do the work for now
		return models.Address{}, models.NewErr(500, "something_went_wrong")
	}

	return
}

func (svc *AddressSVC) validateCreate(in models.AddressReq) (err error) {

	err = svc.UUIDInterface.IsValid(in.UserID)
	if err != nil {
		return models.NewErr(400, "invalid_id")
	}

	if in.Details == "" {
		err = models.NewErr(400, "invalid_details")
		return
	}

	if len(in.Details) > 100 {
		err = models.NewErr(400, "invalid_details")
		return
	}

	if in.ZipCode == "" {
		err = models.NewErr(400, "invalid_zipcode")
		return
	}

	// todo add more validade if need

	return
}
