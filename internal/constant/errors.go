package constant

import (
	"errors"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrNotImplemented  = errors.New("not implemented yet")
	ErrIncorrectFarmID = errors.New("incorrect farm id")
)
