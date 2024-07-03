package address

import (
	"app/models"
	"errors"
	"testing"
)

// go test -v -failfast -run ^TestAddressSVC_validateCreate$
func TestAddressSVC_validateCreate(t *testing.T) {
	type fields struct {
		UUIDInterface UUIDInterface
		RepoInterface RepoInterface
	}
	type args struct {
		in models.AddressReq
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		errString string
	}{
		{
			name: "uuid invalid",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_not_nice_id").Return(errors.New("bata")).Times(1)

				f.UUIDInterface = id
				return
			}(),
			args: args{
				in: models.AddressReq{
					UserID:  "my_not_nice_id",
					ZipCode: "",
					Details: "",
					State:   "",
					Country: "",
					City:    "",
				},
			},
			wantErr:   true,
			errString: "400:invalid_id",
		},
		{
			name: "empty details",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_not_nice_id").Return(nil).Times(1)

				f.UUIDInterface = id
				return
			}(),
			args: args{
				in: models.AddressReq{
					UserID:  "my_not_nice_id",
					ZipCode: "",
					Details: "",
					State:   "",
					Country: "",
					City:    "",
				},
			},
			wantErr:   true,
			errString: "400:invalid_details",
		},
		{
			name: "invalid details",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_not_nice_id").Return(nil).Times(1)

				f.UUIDInterface = id
				return
			}(),
			args: args{
				in: models.AddressReq{
					UserID:  "my_not_nice_id",
					ZipCode: "",
					Details: "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789",
					State:   "",
					Country: "",
					City:    "",
				},
			},
			wantErr:   true,
			errString: "400:invalid_details",
		},
		{
			name: "invalid zip code",
			fields: func() (f fields) {
				id := &uuidMock{}
				id.On("IsValid", "my_not_nice_id").Return(nil).Times(1)

				f.UUIDInterface = id
				return
			}(),
			args: args{
				in: models.AddressReq{
					UserID:  "my_not_nice_id",
					ZipCode: "",
					Details: "teste rua lalala",
					State:   "",
					Country: "",
					City:    "",
				},
			},
			wantErr:   true,
			errString: "400:invalid_zipcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &AddressSVC{
				UUIDInterface: tt.fields.UUIDInterface,
				RepoInterface: tt.fields.RepoInterface,
			}
			err := svc.validateCreate(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressSVC.validateCreate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				if err.Error() != tt.errString {
					t.Errorf("AddressSVC.validateCreate()  error = %s, wantErr %s", err.Error(), tt.errString)
					return
				}
			}
		})
	}
}
