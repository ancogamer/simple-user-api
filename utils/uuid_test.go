package utils

import (
	"testing"

	"github.com/google/uuid"
)

// go test -v -failfast -run ^TestUUIDSvc_IsValid$
func TestUUIDSvc_IsValid(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name      string
		svc       *UUIDSvc
		args      args
		wantErr   bool
		errString string
	}{
		{
			name: "fail",
			svc:  &UUIDSvc{},
			args: args{
				str: "batata",
			},
			wantErr:   true,
			errString: "invalid UUID length: 6",
		},
		{
			name: "ok",
			svc:  &UUIDSvc{},
			args: args{
				str: uuid.New().String(),
			},
			wantErr:   false,
			errString: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &UUIDSvc{}
			err := svc.IsValid(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("UUIDSvc.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				if err.Error() != tt.errString {
					t.Errorf("UUIDSvc.IsValid()  error = %s, wantErr %s", err.Error(), tt.errString)
					return
				}
			}
		})
	}
}
