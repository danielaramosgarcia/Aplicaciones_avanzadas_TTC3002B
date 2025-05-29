package VM

import "babyduck/data_structures"

type Machine struct {
	globalMemory map[int]interface{} // Global variables are stored in a map with their addresses.
	// Memory is a map that holds the global variables.
	memoryStack []*FuncInstance
	Quads       []data_structures.Quadruple
	IP          int
	Name        string // Name of the program being executed.
}

type FuncInstance struct {
	// FuncInstance represents an instance of a function with its local variables.
	tempMemory  map[int]int
	localMemory map[int]interface{}
}

const (
	ADD       = 10
	SUB       = 20
	MUL       = 30
	DIV       = 40
	LT        = 50
	GT        = 60
	NEQ       = 70
	GOTO      = 80
	GOTOFALSE = 81
	PRINT     = 90
	ASSIGN    = 100
	END       = 110
)
