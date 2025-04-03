//Tarea 1 de Desarrollo de Aplicaciones Avanzadas TTC3002B
//Daniela Ramos A01174259

package main

import (
	"fmt"
	"tarea1/data_classes"
)

func main() {

	// Manipulando la estructura de datos de Stack
	stack := data_classes.Stack{
		Values: []int{},
		Top:    -1,
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Print("Stack: ")
	stack.Print()

	value, ok := stack.Pop()
	if ok {
		fmt.Printf("Popped: %d\n", value)
		fmt.Print("Stack after pop: ")
		stack.Print()
	} else {
		fmt.Println("Stack is empty")
	}
}
