package user

import (
	"app/models"
	"context"
	"log"
)

// Get Gets a user
func (svc *UserSVC) Get(ctx context.Context, id string) (out models.UserResp, err error) {
	err = svc.UUIDInterface.IsValid(id)
	if err != nil {
		return models.UserResp{}, models.NewErr(400, "invalid_id")
	}

	user, err := svc.RepoInterface.Get(ctx, id)
	if err != nil {
		// only logs the error and return a empty struct for user, since we need to know what happen, but the user dont
		// and this is a get, so we cover the 'not found' problem without checking if the record exist, saving 1 call
		log.Default().Print("get user err: ", err) // not the fast log out there, but will do the work for now
		return models.UserResp{}, nil
	}
	out.ID = user.ID
	out.Name = user.Name
	out.LastName = user.LastName
	out.Age = user.Age
	return
}
