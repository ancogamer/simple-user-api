package user

import (
	"app/models"
	"context"
	"log"
)

// Patch patchs a expecific field of user
func (svc *UserSVC) Patch(ctx context.Context, in models.UserReq, id string) (err error) {

	err = svc.UUIDInterface.IsValid(id)
	if err != nil {
		return models.NewErr(400, "invalid_id")
	}
	// validating new fields info, for now since this op dont go outside
	// it comes first than repo ones
	err = validateUpdate(in)
	if err != nil {
		return
	}

	// we could also check by db constraints, and since i love postgres,
	// i would do that
	ok, err := svc.RepoInterface.ExistID(ctx, id)
	if err != nil {
		log.Default().Print("patch user err: ", err) // not the fast log out there, but will do the work for now
		return models.NewErr(500, "something_went_wrong")
	}
	if !ok {
		return models.NewErr(404, "id_not_found")
	}
	user := models.User{
		ID:       id,
		Name:     in.Name,
		LastName: &in.LastName,
		Age:      in.Age,
	}
	if in.LastName == "" {
		user.LastName = nil
	}

	err = svc.RepoInterface.Patch(ctx, user)
	if err != nil {
		log.Default().Print("patch user err: ", err) // not the fast log out there, but will do the work for now
		return models.NewErr(500, "something_went_wrong")
	}

	return
}

func validateUpdate(in models.UserReq) (err error) {
	okUpdate := false
	if in.Name != "" {

		if len([]rune(in.Name)) < 3 {
			err = models.NewErr(400, "invalid_name")
			return
		}

		if len([]rune(in.Name)) > 50 {
			err = models.NewErr(400, "invalid_name")
			return
		}
		okUpdate = true
	}

	if in.LastName != "" {
		if len([]rune(in.LastName)) > 50 {
			err = models.NewErr(400, "invalid_last_name")
			return
		}
		okUpdate = true
	}

	if in.Age != 0 {
		if in.Age < 0 {
			err = models.NewErr(400, "invalid_age")
			return
		}
		okUpdate = true
	}

	if !okUpdate {
		err = models.NewErr(400, "nothing_to_update")
		return
	}

	return
}
