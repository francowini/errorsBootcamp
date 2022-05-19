package files

import (
	"os"
	"strings"
)

const (
	slash = "/"
	percent = "%"
	ampersand = "&"
	plusSign = "+"
)

type Service struct {}

func NewService () *Service{
	return &Service{}
}

func (s *Service) OpenFile(path string) (*os.File, error) {
	if err := stringValidator(path); err != nil {
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file, nil
}


func stringValidator (str string) error {
	if strings.Contains(str, ampersand) || strings.Contains(str, percent) ||  strings.Contains(str, plusSign) {
		return ErrInvalidCharacter
	}
	if strings.Contains(str, slash) {
		return ErrNoSlash
	}
	return nil
}