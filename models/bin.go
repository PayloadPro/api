package models

import (
	"errors"
)

// Bin is a designated space to partition requests
type Bin struct {
	ID string
}

// ErrBinNotFound is returned when an bin cannot be found
var ErrBinNotFound = errors.New("Bin could not be found")
