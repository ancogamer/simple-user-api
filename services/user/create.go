package user

import (
	"app/models"
	"context"
	"log"
)

// Create creates a user in the DB
func (svc *UserSVC) Create(ctx context.Context, in models.UserReq) (out models.User, err error) {
	err = validateCreate(in)
	if err != nil {
		return
	}

	out.LastName = &in.LastName
	if in.LastName == "" {
		out.LastName = nil
	}

	//out.ID = uuid.New().String() // i know, but since the mock will complain and i dont want to make my own mock logic
	out.ID = svc.UUIDInterface.NewString() // lets inject this (complexy++)
	out.Name = in.Name
	out.Age = in.Age

	err = svc.RepoInterface.Create(ctx, out)
	if err != nil {
		log.Default().Print("delete user err: ", err) // not the fast log out there, but will do the work for now
		return models.User{}, models.NewErr(500, "something_went_wrong")
	}

	return
}

func validateCreate(in models.UserReq) (err error) {
	if in.Name == "" {
		err = models.NewErr(400, "invalid_name")
		return
	}

	if len([]rune(in.Name)) < 3 {
		err = models.NewErr(400, "invalid_name")
		return
	}

	if len([]rune(in.Name)) > 50 {
		err = models.NewErr(400, "invalid_name")
		return
	}

	if in.Age <= 0 {
		err = models.NewErr(400, "invalid_age")
		return
	}

	if in.LastName != "" {
		if len([]rune(in.LastName)) > 50 {
			err = models.NewErr(400, "invalid_last_name")
			return
		}
	}
	return
}
