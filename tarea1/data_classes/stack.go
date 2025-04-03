package data_classes

import "fmt"

// Struck para la estructura de datos Stack
type Stack struct {
	Values []int
	Top    int
}

// Metodo para agregar un elemento a la pila
func (s *Stack) Push(value int) {
	s.Values = append(s.Values, value)
	s.Top++
}

// Metodo para obtener y eliminar un elemento de la pila
func (s *Stack) Pop() (int, bool) {
	if s.Top == -1 {
		return 0, false
	}
	value := s.Values[s.Top]
	s.Values = s.Values[:s.Top]
	s.Top--
	return value, true
}

// Metodo para imprimir la pila
func (s *Stack) Print() {
	for i := 0; i <= s.Top; i++ {
		fmt.Printf("%d ", s.Values[i])
	}
	fmt.Println()
}

// Metodo para verificar si la pila está vacía
func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

// Metodo para obtener el tamaño de la pila
func (s *Stack) Size() int {
	return s.Top + 1
}

// Metodo para obtener el elemento en la parte superior de la pila sin eliminarlo
func (s *Stack) Peek() (int, bool) {
	if s.Top == -1 {
		return 0, false
	}
	return s.Values[s.Top], true
}
