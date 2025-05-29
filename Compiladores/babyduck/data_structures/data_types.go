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
	FuncDir       *FuncDir
	currentFunc   *FuncEntry
	FuncSignature *FuncSignature // Firma de la función actual
	//Pilas y cola para cuádruplos
	OperatorStack []int
	OperandStack  []int
	TypeStack     []int

	Quads       QuadQueue
	JumpStack   []int          // Pila de saltos pendientes
	FuncCount   int            // Contador de funciones
	FuncIndex   map[int]string // Mapa de índices a nombres de funciones
	ConstTable  Constable      // Tabla de constantes
	AddedConst  []string
	programName string // Nombre del programa
}

// Constable
type Constable struct {
	Num   map[int]int
	Float map[int]float64
	Str   map[int]string
	Bool  map[int]bool
}

// Struct para llevar la firma de una función
type FuncSignature struct {
	ParamSignature []int // Firma de parámetros para la función actual
	ParamLength    int   // Cantidad de parámetros
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
	index      int // índice de la función en FuncDir
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
