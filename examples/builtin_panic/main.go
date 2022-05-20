// Snippet para mostra un custom panic y como quiebra en cascada cada una de las funciones.
package main

func main() {
	foo()
	print("llegue hasta el final de la función main")
}

func foo() {
	foo2()
	print("llegue hasta el final de la función foo")
}

func foo2(){
	panic("oh no!")
	print("llegue hasta el final de la función foo 2")
}