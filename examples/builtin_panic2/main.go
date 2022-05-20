package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("somefile.txt")
	if err != nil {
        // custom panic transportando el error que me arrojo os.Open
		panic(err)
	}
	fmt.Print("logre llegar al final del programa\n")
}
