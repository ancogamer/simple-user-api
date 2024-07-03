package user

import (
	"app/models"
	"context"
	"log"
)

// Delete deles a user
func (svc *UserSVC) Delete(ctx context.Context, id string) (err error) {

	err = svc.UUIDInterface.IsValid(id)
	if err != nil {
		return models.NewErr(400, "invalid_id")
	}
	// we could also check by db constraints, and since i love postgres
	// i would do that

	ok, err := svc.RepoInterface.ExistID(ctx, id)
	if err != nil {
		log.Default().Print("delete user err: ", err) // not the fast log out there, but will do the work for now
		return models.NewErr(500, "something_went_wrong")
	}
	if !ok {
		return models.NewErr(404, "id_not_found")
	}

	err = svc.RepoInterface.Delete(ctx, id)
	if err != nil {
		log.Default().Print("delete user err: ", err) // not the fast log out there, but will do the work for now
		return models.NewErr(500, "something_went_wrong")
	}

	return
}
