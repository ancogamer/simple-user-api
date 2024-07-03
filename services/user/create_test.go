package user

import (
	"app/models"
	"app/utils"
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// go test -v -failfast -run ^TestUserSVC_Create$
func TestUserSVC_Create(t *testing.T) {
	type fields struct {
		RepoInterface RepoInterface
		UUIDInterface UUIDInterface
	}
	type args struct {
		ctx context.Context
		in  models.UserReq
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantOut   models.User
		wantErr   bool
		errString string
	}{
		{
			name: "fail validade create",
			fields: func() (f fields) {
				return
			}(),
			args: args{
				ctx: context.Background(),
				in: models.UserReq{
					Name:     "test",
					LastName: "01234aa56789012345678901234567890123456789012345678",
					Age:      1,
				},
			},
			wantOut:   models.User{},
			wantErr:   true,
			errString: "400:invalid_last_name",
		},
		{
			name: "fail create",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("NewString").Return("my_nice_id").Times(1)

				user := &repoMock{}
				user.On("Create", context.Background(), models.User{
					ID:       "my_nice_id",
					Name:     "test",
					LastName: utils.StringToPointer("0123456789012345678901234567890123456789012345678"),
					Age:      1,
				}).Return(errors.New("fail create")).Times(1)

				f.UUIDInterface = id
				f.RepoInterface = user
				return
			}(),
			args: args{
				ctx: context.Background(),
				in:  models.UserReq{},
			},
			wantOut:   models.User{},
			wantErr:   true,
			errString: "400:invalid_name",
		},
		{
			name: "ok",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("NewString").Return("my_nice_id").Times(1)

				user := &repoMock{}
				user.On("Create", context.Background(), models.User{
					ID:       "my_nice_id",
					Name:     "test",
					LastName: utils.StringToPointer("0123456789012345678901234567890123456789012345678"),
					Age:      1,
				}).Return(nil).Times(1)

				f.UUIDInterface = id
				f.RepoInterface = user
				return
			}(),
			args: args{
				ctx: context.Background(),
				in: models.UserReq{
					Name:     "test",
					LastName: "0123456789012345678901234567890123456789012345678",
					Age:      1,
				},
			},
			wantOut: models.User{
				ID:       "my_nice_id",
				Name:     "test",
				LastName: utils.StringToPointer("0123456789012345678901234567890123456789012345678"),
				Age:      1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &UserSVC{
				UUIDInterface: tt.fields.UUIDInterface,
				RepoInterface: tt.fields.RepoInterface,
			}
			gotOut, err := svc.Create(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserSVC.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(gotOut, tt.wantOut); diff != "" {
				t.Errorf("UserSVC.Create() = %s", diff)
				return
			}

			if tt.wantErr {
				if err.Error() != tt.errString {
					t.Errorf("vUserSVC.Create()  error = %s, wantErr %s", err.Error(), tt.errString)
					return
				}
			}
		})
	}
}

// go test -v -failfast -run ^Test_validateCreate$
func Test_validateCreate(t *testing.T) {
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
			name:      "no name empty",
			args:      args{},
			wantErr:   true,
			errString: "400:invalid_name",
		},
		{
			name: "no name smaller than 3",
			args: args{
				in: models.UserReq{
					Name:     "01",
					LastName: "",
					Age:      0,
				},
			},
			wantErr:   true,
			errString: "400:invalid_name",
		},
		{
			name: "no name bigger than 50",
			args: args{
				in: models.UserReq{
					Name:     "012345678901234567890123456789012345678901234567890",
					LastName: "",
					Age:      0,
				},
			},
			wantErr:   true,
			errString: "400:invalid_name",
		},
		{
			name: "no age",
			args: args{
				in: models.UserReq{
					Name:     "test",
					LastName: "",
					Age:      0,
				},
			},
			wantErr:   true,
			errString: "400:invalid_age",
		},
		{
			name: "-1 age",
			args: args{
				in: models.UserReq{
					Name:     "test",
					LastName: "",
					Age:      -1,
				},
			},
			wantErr:   true,
			errString: "400:invalid_age",
		},
		{
			name: "lastname bigger than 50",
			args: args{
				in: models.UserReq{
					Name:     "test",
					LastName: "012345678901234567890123456789012345678901234567890",
					Age:      1,
				},
			},
			wantErr:   true,
			errString: "400:invalid_last_name",
		},
		{
			name: "ok",
			args: args{
				in: models.UserReq{
					Name:     "test",
					LastName: "01234567890123456789012345678901234567890123456789",
					Age:      1,
				},
			},
			wantErr:   false,
			errString: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreate(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateCreate() error = %s, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err.Error() != tt.errString {
					t.Errorf("validateCreate() error = %s, wantErr %s", err.Error(), tt.errString)
					return
				}
			}
		})
	}
}

var err error

// go test -benchmem -run=^$ -bench ^Benchmark_validateCreate$ app/services/user
func Benchmark_validateCreate(b *testing.B) {
	b.Run("ok", func(b *testing.B) {
		err = validateCreate(models.UserReq{
			Name:     "test",
			LastName: "01234567890123456789012345678901234567890123456789",
			Age:      1,
		})
	})

}
