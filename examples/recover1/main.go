// Snippet mostrando como recuperarse de un panic.
package main

import (
	"fmt"
	"log"
)

func main() {
	divideByZero()
	fmt.Println("we survived dividing by zero!")
}

func divideByZero() {
	//recuperación del panic
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
	return a / b
}