// Snippet 2 para mostrarse como recuperarse de un panic
package main

import (
	"fmt"
	"log"
)

// someStruct implementa una función para utilizarla en el defer
type someStruct struct {}
func (ss *someStruct) foo3() {
	fmt.Println("hello from the deferred function 3!")
}

// definición de una función como variable
var foo2 = func(str string) {
	fmt.Println(str)
}

func main() {
	doStuffs()
	fmt.Println("hello post do stuff panicking")
}

func doStuffs () {
	testStruct := someStruct{}
	// puedo hacer un defer de una función definida en la misma linea
	defer func() {
		// recuperación del error
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
		fmt.Println("hello from the deferred function 1!")
	}()
	// puedo hacer un defer de una función predefinida como variable
	defer foo2("hello from the deferred function 2!")
	// o puedo hacer un defer de un método de alguna estructura
	defer testStruct.foo3()
	fmt.Println("hello previus panicking")
	throwPanic()
	fmt.Println("do stuff finish")
}

func throwPanic () {
	var zero int
	_ = 1/zero
}