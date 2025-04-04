//Tarea 1 de Desarrollo de Aplicaciones Avanzadas TTC3002B
//Daniela Ramos A01174259

package main

import (
	"fmt"
	"tarea1/data_classes"
)

func main() {

	fmt.Println("-----------------------------------------------------")
	fmt.Println("Tarea 1 de Desarrollo de Aplicaciones Avanzadas TTC3002B")
	fmt.Println("Daniela Ramos A01174259")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Estructura de datos: Stack")

	// Test cases para la clase de Stack
	// Manipulando la estructura de datos de Stack
	stack := data_classes.Stack{
		Values: []int{},
		Top:    -1,
	}

	// Agregando elementos a la pila
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Debe imprimir: 1 2 3
	fmt.Print("Stack: ")
	stack.Print()

	// Verificando el tamaño de la pila
	value, ok := stack.Pop()

	// Debe imprimir: 3
	if ok {
		fmt.Printf("Popped: %d\n", value)
		fmt.Print("Stack despues de pop: ")
		stack.Print()
	} else {
		fmt.Println("Stack esta vacio")
	}

	// Verificando si la pila está vacía
	if stack.IsEmpty() {
		fmt.Println("Stack esta vacio")
	} else {
		fmt.Println("Stack no esta vacio")
	}
	// Verificando el tamaño de la pila
	fmt.Printf("Tamano de Stack: %d\n", stack.Size())
	// Verificando el elemento en la parte superior de la pila
	value, ok = stack.Peek()
	if ok {
		fmt.Printf("Primer elemento: %d\n", value)
	} else {
		fmt.Println("Stack esta vacio")
	}

	// Test cases para la clase de Queue
	// Manipulando la estructura de datos de Queue
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Estructura de datos: Queue")

	queue := data_classes.Queue{
		Values: []int{},
		Front:  0,
	}
	// Agregando elementos a la cola
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Debe imprimir: 10 20 30
	fmt.Print("Queue: ")
	queue.Print()

	// Debe eliminar el primer elemento de la cola
	value, ok = queue.Dequeue()
	value1, ok1 := queue.Dequeue()
	// Debe imprimir: 10
	if ok && ok1 {
		fmt.Printf("Dequeued: %d\n", value)
		fmt.Printf("Dequeued: %d\n", value1)
		fmt.Print("Queue despues de dequeue: ")
		queue.Print()
	} else {
		fmt.Println("Queue esta vacio")
	}

	//Vuelve a imprimir la cola
	fmt.Print("Queue: ")
	queue.Print()
	// Verificando si la cola está vacía
	if queue.IsEmpty() {
		fmt.Println("Queue esta vacio")
	} else {
		fmt.Println("Queue no esta vacio")
	}
	// Verificando el tamaño de la cola
	fmt.Printf("Tamano de Queue: %d\n", queue.Size())
	// Verificando el elemento en la parte frontal de la cola
	value, ok = queue.Peek()
	if ok {
		fmt.Printf("Primer elemento: %d\n", value)
	} else {
		fmt.Println("Queue esta vacio")
	}

	// Test cases para la clase de Dictionary
	// Manipulando la estructura de datos de Dictionary
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Estructura de datos: Dictionary")

	// Crear un nuevo diccionario
	dictionary := data_classes.Dictionary{
		Keys:   []string{},
		Values: []string{},
	}
	// Agregar elementos al diccionario
	dictionary.Add("manzana", "roja")
	dictionary.Add("durazno", "naranja")
	dictionary.Add("platano", "amarillo")
	// Imprimir el diccionario
	fmt.Print("Diccionario: ")
	fmt.Println()
	dictionary.Print()

	// Eliminar un elemento del diccionario
	ok = dictionary.Remove("durazno")
	if ok {
		fmt.Println("durazno removido")
	} else {
		fmt.Println("Llave no encontrada")
	}
	// Imprimir el diccionario después de eliminar un elemento
	fmt.Print("Dictionary despues de quitar durazno: ")
	fmt.Println()
	dictionary.Print()
	// Verificando si el diccionario está vacío
	if dictionary.IsEmpty() {
		fmt.Println("Dictionary esta vacio")
	} else {
		fmt.Println("Dictionary no esta vacio")
	}
	// Verificando el tamaño del diccionario
	fmt.Printf("Tamano de Dictionary: %d\n", dictionary.Size())
	// Verificando si una clave existe en el diccionario
	ok = dictionary.ContainsKey("platano")
	if ok {
		fmt.Println("Llave platano existe")
	} else {
		fmt.Println("Llave platano no existe")
	}
	// Verificando si un valor existe en el diccionario
	ok = dictionary.ContainsValue("roja")
	if ok {
		fmt.Println("Value roja existe")
	} else {
		fmt.Println("Value roja no existe")
	}
	// Verificando si el diccionario está vacío
	if dictionary.IsEmpty() {
		fmt.Println("Dictionary esta vacio")
	} else {
		fmt.Println("Dictionary no esta vacio")
	}

}
