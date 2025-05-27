package data_structures

// “enum” de tipos básicos en BabyDuck
type Tipo int

// const (
// 	Int    Tipo = iota // 0
// 	Float              // 1
// 	Bool               // 2
// 	Void               // 3
// 	String             // 4
// )

var Int = 0
var Float = 1
var Bool = 2
var Void = 3
var String = 4
var NP = 5

type Quadruple struct {
	Op     int
	Arg1   int
	Arg2   int
	Result int
}

type QuadQueue struct {
	Quads []Quadruple
}

// type Jump struct {
// 	Op     int // Operación de salto (GOTO, GOTOFALSE, etc.)

// Context mantiene el estado global del compilador:
type Context struct {
	// Tablas de variables y funciones
	// GlobalVars  *VarTable
	FuncDir     *FuncDir
	currentFunc *FuncEntry

	//Pilas y cola para cuádruplos
	OperatorStack []int
	OperandStack  []int
	TypeStack     []int

	Quads     QuadQueue
	JumpStack []int // Pila de saltos pendientes

	TempCounter  int
	LabelCounter int
}

// Param representa un parámetro de función: nombre y tipo.
type Param struct {
	Name string
	Type int
}

type SpaceVariables struct {
	Var, Temp int
}

// Representa la firma y el ámbito de una función
type FuncEntry struct {
	Name       string    // identificador de la función
	ReturnType int       // tipo de retorno
	ParamTypes []int     // tipos de parámetros en orden
	VarTable   *VarTable // tabla de variables locales
	CuadStart  int
	Space      *SpaceVariables
}

// Contenedor global de todas las funciones
type FuncDir struct {
	funcs map[string]*FuncEntry // nombre → FuncEntry
}

// VarEntry representa una variable con su nombre, tipo y dirección.
type VarEntry struct {
	Name   string
	Type   int
	DirInt int // dirección en memoria
}

// VarTable gestiona un ámbito de variables, indexadas por dirección.
type VarTable struct {
	vars   map[int]*VarEntry // DirInt → VarEntry
	parent *VarTable         // ámbito superior
}
