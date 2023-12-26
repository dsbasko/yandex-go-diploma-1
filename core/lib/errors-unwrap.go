package lib

import "errors"

func ErrorsUnwrap(err error) error {
	for {
		unwrapped := errors.Unwrap(err)

		if unwrapped == nil {
			return err
		}

		err = unwrapped
	}
}
