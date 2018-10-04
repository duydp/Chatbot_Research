package ui

import (
	"errors"
)

var NotFound error = errors.New("FAQ not found")
var ErrEmptyIntent error = errors.New("Intent is empty")
var ErrNoAnswer error = errors.New("There's no answer")