package handles

import "errors"

var (
	ErrEmptyBody        = errors.New("empty body")
	ErrEmptyID          = errors.New("empty id")
	ErrWrongContentType = errors.New("wrong content type")
)
