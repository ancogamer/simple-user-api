package user

import (
	"app/models"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

// go test -v -failfast -run ^TestUserSVC_Patch$
func TestUserSVC_Patch(t *testing.T) {
	type fields struct {
		UUIDInterface UUIDInterface
		RepoInterface RepoInterface
	}
	type args struct {
		ctx context.Context
		in  models.UserReq
		id  string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		errString string
	}{
		{
			name: "Fail valid id",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_not_so_nice_id").Return(errors.New("kabum"))

				f.UUIDInterface = id
				return
			}(),
			args: args{
				ctx: nil,
				in:  models.UserReq{},
				id:  "my_not_so_nice_id",
			},
			wantErr:   true,
			errString: "400:invalid_id",
		},
		{
			name: "Fail validateUpdate",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_nice_id").Return(nil)

				f.UUIDInterface = id
				return
			}(),
			args: args{
				ctx: nil,
				in:  models.UserReq{},
				id:  "my_nice_id",
			},
			wantErr:   true,
			errString: "400:nothing_to_update",
		},
		{
			name: "Fail ExistID error",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_nice_id").Return(nil)

				repo := &repoMock{}
				repo.On("ExistID", nil, "my_nice_id").Return(false, errors.New("batata"))

				f.RepoInterface = repo
				f.UUIDInterface = id
				return
			}(),
			args: args{
				ctx: nil,
				in: models.UserReq{
					Name:     "aaa",
					LastName: "",
					Age:      0,
				},
				id: "my_nice_id",
			},
			wantErr:   true,
			errString: "500:something_went_wrong",
		},
		{
			name: "Fail ExistID not found",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_nice_id").Return(nil)

				repo := &repoMock{}
				repo.On("ExistID", nil, "my_nice_id").Return(false, nil)

				f.RepoInterface = repo
				f.UUIDInterface = id
				return
			}(),
			args: args{
				ctx: nil,
				in: models.UserReq{
					Name:     "aaa",
					LastName: "",
					Age:      0,
				},
				id: "my_nice_id",
			},
			wantErr:   true,
			errString: "404:id_not_found",
		},
		{
			name: "Fail Patch",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_nice_id").Return(nil)

				repo := &repoMock{}
				repo.On("ExistID", nil, "my_nice_id").Return(true, nil)
				repo.On("Patch", nil, mock.Anything).Return(errors.New("well, the DB is gone"))

				f.RepoInterface = repo
				f.UUIDInterface = id
				return
			}(),
			args: args{
				ctx: nil,
				in: models.UserReq{
					Name:     "aaa",
					LastName: "",
					Age:      0,
				},
				id: "my_nice_id",
			},
			wantErr:   true,
			errString: "500:something_went_wrong",
		},
		{
			name: "ok",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_nice_id").Return(nil)

				repo := &repoMock{}
				repo.On("ExistID", nil, "my_nice_id").Return(true, nil)
				repo.On("Patch", nil, mock.Anything).Return(nil)

				f.RepoInterface = repo
				f.UUIDInterface = id
				return
			}(),
			args: args{
				ctx: nil,
				in: models.UserReq{
					Name:     "aaa",
					LastName: "",
					Age:      0,
				},
				id: "my_nice_id",
			},
			wantErr:   false,
			errString: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &UserSVC{
				UUIDInterface: tt.fields.UUIDInterface,
				RepoInterface: tt.fields.RepoInterface,
			}
			err := svc.Patch(tt.args.ctx, tt.args.in, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserSVC.Patch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err.Error() != tt.errString {
					t.Errorf("UserSVC.Patch() error = %s, wantErr %s", err.Error(), tt.errString)
					return
				}
			}
		})
	}
}

// go test -v -failfast -run ^Test_validateUpdate$
func Test_validateUpdate(t *testing.T) {
	type args struct {
		in models.UserReq
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		errString string
	}{
		{
			name: "invalid name less than 3",
			args: args{
				in: models.UserReq{
					Name:     "a",
					LastName: "",
					Age:      0,
				},
			},
			wantErr:   true,
			errString: "400:invalid_name",
		},
		{
			name: "invalid name more than 50",
			args: args{
				in: models.UserReq{
					Name:     "012345678901234567890012345678901234567890123456789",
					LastName: "",
					Age:      0,
				},
			},
			wantErr:   true,
			errString: "400:invalid_name",
		},
		{
			name: "invalid last name name more than 50",
			args: args{
				in: models.UserReq{
					Name:     "test",
					LastName: "012345678901234567890012345678901234567890123456789",
					Age:      0,
				},
			},
			wantErr:   true,
			errString: "400:invalid_last_name",
		},
		{
			name: "invalid last name name more than 50",
			args: args{
				in: models.UserReq{
					Name:     "test",
					LastName: "02345678901234567890012345678901234567890123456789",
					Age:      -1,
				},
			},
			wantErr:   true,
			errString: "400:invalid_age",
		},
		{
			name: "not to update",
			args: args{
				in: models.UserReq{
					Name:     "",
					LastName: "",
					Age:      0,
				},
			},
			wantErr:   true,
			errString: "400:nothing_to_update",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateUpdate(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err.Error() != tt.errString {
					t.Errorf("Test_validateUpdate() error = %s, wantErr %s", err.Error(), tt.errString)
					return
				}
			}
		})
	}
}
