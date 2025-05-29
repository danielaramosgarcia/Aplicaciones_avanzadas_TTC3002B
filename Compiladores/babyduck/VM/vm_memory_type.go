package VM

import (
	"babyduck/data_structures"
	"strings"
)

// type Machine struct {
// 	globalMemory map[int]interface{} // Global variables and constants.
// 	// Memory is a map that holds the global variables.
// 	MS     []string
// 	Quads  []data_structures.Quadruple
// 	IP     int
// 	Name   string              // Name of the program being executed.
// 	Memory map[int]interface{} // Memory for the current function instance.
// }

// type FuncInstance struct {
// 	// FuncInstance represents an instance of a function with its local variables.
// 	tempMemory  map[int]int
// 	localMemory map[int]interface{}
// }

type Machine struct {
	Ctx         *data_structures.Context
	memStack    []map[int]interface{}
	retStack    []int
	Quads       []data_structures.Quadruple
	IP          int
	ParamAddrs  map[string][]int // <-- nuevo
	PrintOutput strings.Builder
}

const (
	ADD       = 10 // suma
	SUB       = 20 // resta
	MUL       = 30 // multiplicación
	DIV       = 40 // división
	LT        = 50 // menor que
	GT        = 60 // mayor que
	NEQ       = 70 // no igual
	RPAR      = 80 // paréntesis derecho
	GOTO      = 90 // salto incondicional
	GOTOFALSE = 91 // salto si condición es falsa
	EQ        = 92 // asignación
	PRINT     = 93 // imprimir
	ENDF      = 94 // fin de función
	ERA       = 95 // inicio de función
	PARAM     = 96 // parámetro de función
	GOSUB     = 97 // llamada a función
	END       = 98 // fin del programa
)
