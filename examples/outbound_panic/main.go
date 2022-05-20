// snippet para ejemplificar panic por acceder a un índice que no existe en un array/slice
package main

import (
	"fmt"
)

func main() {
	names := []string{
		"lobster",
		"sea urchin",
		"sea cucumber",
	}
	// accedo a names[3], el cual no existe porque se comienza indexando desde 0. Paaaaanic
	fmt.Println("My favorite sea creature is:", names[len(names)])
}