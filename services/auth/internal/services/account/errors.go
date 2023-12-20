package account

import (
	"errors"
	"fmt"
)

var (
	ErrUnauthorized       = errors.New("unauthorized")
	ErrArgumentsNotFilled = errors.New("not all arguments are filled in")

	ErrUsernameMinLength = fmt.Errorf("username must be at least %v characters long", UsernameMinLength)
	ErrUsernameMaxLength = fmt.Errorf("username cannot exceed a maximum length of %v characters", UsernameMaxLength)

	ErrPasswordMinLength = fmt.Errorf("password must be at least %v characters long", PasswordMinLength)
	ErrPasswordMaxLength = fmt.Errorf("password cannot exceed a maximum length of %v characters", PasswordMaxLength)
)
