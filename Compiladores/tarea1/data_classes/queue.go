package data_classes

import "fmt"

// Struck para la estructura de datos Queue
type Queue struct {
	Values []int
	Front  int
}

// Metodo para agregar un elemento a la cola
func (q *Queue) Enqueue(value int) {
	q.Values = append(q.Values, value)
	q.Front++
}

// Metodo para obtener y eliminar un elemento de la cola
func (q *Queue) Dequeue() (int, bool) {
	if q.Front == 0 {
		return 0, false
	}
	value := q.Values[0]
	q.Values = q.Values[1:]
	q.Front--
	return value, true
}

// Metodo para imprimir la cola
func (q *Queue) Print() {
	for i := 0; i < q.Front; i++ {
		fmt.Printf("%d ", q.Values[i])
	}
	fmt.Println()
}

// Metodo para verificar si la cola está vacía
func (q *Queue) IsEmpty() bool {
	return q.Front == 0
}

// Metodo para obtener el tamaño de la cola
func (q *Queue) Size() int {
	return q.Front
}

// Metodo para obtener el elemento en la parte frontal de la cola sin eliminarlo
func (q *Queue) Peek() (int, bool) {
	if q.Front == 0 {
		return 0, false
	}
	return q.Values[0], true
}
