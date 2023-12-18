package account

import (
	"fmt"
)

var (
	ErrUsernameMinLength = fmt.Errorf("username must be at least %v characters long", UsernameMinLength)
	ErrUsernameMaxLength = fmt.Errorf("username cannot exceed a maximum length of %v characters", UsernameMaxLength)

	ErrPasswordMinLength = fmt.Errorf("password must be at least %v characters long", PasswordMinLength)
	ErrPasswordMaxLength = fmt.Errorf("password cannot exceed a maximum length of %v characters", PasswordMaxLength)
)
