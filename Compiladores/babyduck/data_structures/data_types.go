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

const Int = 0
const Float = 1
const Bool = 2
const Void = 3
const String = 4

// Context mantiene el estado global del compilador:
// - GlobalVars: tabla global de variables
// - FuncDir: directorio de funciones3
// - currentFunc: la función activa (no anidada)

type Context struct {
	// Tablas de variables y funciones
	GlobalVars  *VarTable
	FuncDir     *FuncDir
	currentFunc *FuncEntry

	//Pilas y cola para cuádruplos
	OperatorStack []int
	OperandStack  []int
	TypeStack     []int
	Quads         QuadQueue

	TempCounter  int
	LabelCounter int
}

// Param representa un parámetro de función: nombre y tipo.
// Se especifica el tipo Param para poder reconocer el arreglo
// de Param cuando se reconoce una funcion en la gramática.
type Param struct {
	Name string
	Type int
}

// Simulación de direcciones de memoria (stack) por segmentos:
const (
	// Variables globales [0,500]
	GlobalIntBase   = 1
	GlobalFloatBase = 250
	GlobalLimit     = 500

	// Variables locales [501,1000]
	LocalIntBase   = 501
	LocalFloatBase = 751
	LocalLimit     = 1000

	// Variables temporales [1001,1600]
	TempIntBase   = 1001
	TempFloatBase = 1201
	TempBoolBase  = 1401
	TempLimit     = 1600

	// Constantes [1601,2200]
	ConstIntBase    = 1601
	ConstFloatBase  = 1801
	ConstStringBase = 2001
	ConstLimit      = 2200
)

// Contadores para asignar la siguiente dirección disponible.
var (
	nextGlobalIntAddr   = GlobalIntBase
	nextGlobalFloatAddr = GlobalFloatBase

	nextLocalIntAddr   = LocalIntBase
	nextLocalFloatAddr = LocalFloatBase

	nextTempIntAddr   = TempIntBase
	nextTempFloatAddr = TempFloatBase
	nextTempBoolAddr  = TempBoolBase

	nextConstIntAddr    = ConstIntBase
	nextConstFloatAddr  = ConstFloatBase
	nextConstStringAddr = ConstStringBase
)

type Quadruple struct {
	Op     int
	Arg1   int
	Arg2   int
	Result int
}

type QuadQueue struct {
	Quads []Quadruple
}

// Representa la firma y el ámbito de una función
type FuncEntry struct {
	Name       string    // identificador de la función
	ReturnType int       // tipo de retorno
	ParamTypes []int     // tipos de parámetros en orden
	VarTable   *VarTable // tabla de variables locales
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
