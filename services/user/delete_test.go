package user

import (
	"context"
	"errors"
	"testing"
)

// go test -v -failfast -run ^TestUserSVC_Delete$
func TestUserSVC_Delete(t *testing.T) {
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
		wantErr   bool
		errString string
	}{
		{
			name: "fail check id",
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
			wantErr:   true,
			errString: "400:invalid_id",
		},
		{
			name: "fail id not found",
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
				id:  "my_nice_id",
			},
			wantErr:   true,
			errString: "404:id_not_found",
		},
		{
			name: "fail ExistID error",
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
				id:  "my_nice_id",
			},
			wantErr:   true,
			errString: "500:something_went_wrong",
		},
		{
			name: "fail Delete error",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_nice_id").Return(nil)

				repo := &repoMock{}
				repo.On("ExistID", nil, "my_nice_id").Return(true, nil)
				repo.On("Delete", nil, "my_nice_id").Return(errors.New("batata"))

				f.RepoInterface = repo
				f.UUIDInterface = id
				return
			}(),
			args: args{
				ctx: nil,
				id:  "my_nice_id",
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
				repo.On("Delete", nil, "my_nice_id").Return(nil)

				f.RepoInterface = repo
				f.UUIDInterface = id
				return
			}(),
			args: args{
				ctx: nil,
				id:  "my_nice_id",
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
			err := svc.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserSVC.Delete() error = %v, wantErr %v", err, tt.wantErr)
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
