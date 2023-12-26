package account

import (
	"testing"

	"github.com/dsbasko/yandex-go-diploma-1/core/lib"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func Test_passwordEncode(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantStr  bool
		wantErr  bool
	}{
		{name: "Success", password: "test", wantStr: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passHash, err := passwordEncode(tt.password)

			if tt.wantStr {
				assert.NotZero(t, passHash)
			}

			if tt.wantErr {
				assert.NotZero(t, err)
			}
		})
	}
}

func Test_passwordCompare(t *testing.T) {
	hashedPassword, _ := passwordEncode("test")

	tests := []struct {
		name     string
		password string
		wantErr  error
	}{
		{name: "Success", password: "test", wantErr: nil},
		{name: "Password Not Valid", password: "not_test", wantErr: bcrypt.ErrMismatchedHashAndPassword},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := passwordCompare(hashedPassword, tt.password)

			if err != nil || tt.wantErr != nil {
				assert.Equal(t, lib.ErrorsUnwrap(err), tt.wantErr)
			}
		})
	}
}
