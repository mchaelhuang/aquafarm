package constant

import (
	"errors"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrDuplicateRecord = errors.New("record is duplicate")
	ErrNotImplemented  = errors.New("not implemented yet")
	ErrIncorrectFarmID = errors.New("incorrect farm id")
)
