package commands

import (
	"fmt"
	"os"
)

type FilesOpener interface {
	OpenFile(path string) (*os.File, error)
}

type Service struct {
	filesOpener FilesOpener
}

func NewService (opener FilesOpener) *Service {
	return &Service{filesOpener: opener}
}

func (s *Service) GetArgumentAndReturnFile () (*os.File, error) {
	fmt.Printf("build file name is %s\n", os.Args[0])
	file, err := s.filesOpener.OpenFile(os.Args[1])
	if err != nil {
		return nil, err
	}
	return file, nil
}