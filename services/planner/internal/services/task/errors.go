package task

import (
	"errors"
	"fmt"
)

var (
	ErrArgumentsNotFilled = errors.New("not all arguments are filled in")
	ErrEmptyID            = errors.New("empty id")
	ErrEmptyUserID        = errors.New("empty user id")

	ErrNameMinLength = fmt.Errorf("name must be at least %v characters long", NameMinLength)
	ErrNameMaxLength = fmt.Errorf("name cannot exceed a maximum length of %v characters", NameMaxLength)
)
