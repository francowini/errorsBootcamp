// Snippet para mostrar el funcionamiento  de los defer. 
package main

import "fmt"

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
	testStruct := someStruct{}
	// puedo hacer un defer de una función definida en la misma linea
	defer func() {
		fmt.Println("hello from the deferred function 1!")
	}()
	// puedo hacer un defer de una función predefinida como variable
	defer foo2("hello from the deferred function 2!")
	// o puedo hacer un defer de un método de alguna estructura
	defer testStruct.foo3()
	fmt.Println("hello previus panicking")
	panic("oh no!")
	fmt.Println("hello post panicking")
}