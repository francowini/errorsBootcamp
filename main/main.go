package main

import (
	"fmt"
	"os"

	"github.com/francowini/readerExample/internal/commands"
	"github.com/francowini/readerExample/internal/files"
)

const (
	exitCodeFailGetFile = iota
)

func main() {
	fileService := files.NewService()
	commandsService := commands.NewService(fileService)
	file, err := commandsService.GetArgumentAndReturnFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create fury application: %s", err)
		os.Exit(exitCodeFailGetFile)
	}
	fmt.Printf("file name opened is %s\n", file.Name())
}
