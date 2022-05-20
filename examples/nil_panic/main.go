// snippet para ejemplificar un panic por usar un método de un puntero que apunta a un nil
package main
import (
	"fmt"
)

type Shark struct {
	Name string
}

func (s *Shark) SayHello() {
	fmt.Println("Hi! My name is", s.Name)
}

func main() {
	// s se instancia como tipo &Shark (puntero a shark)
	s := &Shark{"Sammy"}
	// le asigno a s una dirección vacía
	s = nil
	// paaaaanic
	s.SayHello()
}