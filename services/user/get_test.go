package user

import (
	"app/models"
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// go test -v -failfast -run ^TestUserSVC_Get$
func TestUserSVC_Get(t *testing.T) {
	type fields struct {
		UUIDInterface UUIDInterface
		RepoInterface RepoInterface
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantOut   models.UserResp
		wantErr   bool
		errString string
	}{
		{
			name: "invalid id",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_not_so_nice_id").Return(errors.New("kabum"))

				f.UUIDInterface = id
				return
			}(),
			args: args{
				ctx: nil,
				id:  "my_not_so_nice_id",
			},
			wantOut:   models.UserResp{},
			wantErr:   true,
			errString: "400:invalid_id",
		},
		{
			name: "get failed for any reason",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_nice_id").Return(nil)
				repo := &repoMock{}
				repo.On("Get", nil, "my_nice_id").Return(models.User{}, errors.New("kabum"))

				f.RepoInterface = repo
				f.UUIDInterface = id
				return
			}(),
			args: args{
				ctx: nil,
				id:  "my_nice_id",
			},
			wantOut:   models.UserResp{},
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
			gotOut, err := svc.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserSVC.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(gotOut, tt.wantOut); diff != "" {
				t.Errorf("UserSVC.Get() = %s", diff)
			}

			if tt.wantErr {
				if err.Error() != tt.errString {
					t.Errorf("UserSVC.Delete() error = %s, wantErr %s", err.Error(), tt.errString)
					return
				}
			}
		})
	}
}
