package files

import "errors"

var (
	ErrNoSlash = errors.New("no slash in path")
	ErrInvalidCharacter = errors.New("there are some invalid character in path")
)
