package data_structures

// “enum” de tipos básicos en BabyDuck
type Tipo int

const (
	Int    Tipo = iota // 0
	Float              // 1
	Bool               // 2
	Void               // 3
	String             // 4
)

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

// Representa la firma y el ámbito de una función
type FuncEntry struct {
	Name       string    // identificador de la función
	ReturnType Tipo      // tipo de retorno
	ParamTypes []Tipo    // tipos de parámetros en orden
	VarTable   *VarTable // tabla de variables locales
}

// Contenedor global de todas las funciones
type FuncDir struct {
	funcs map[string]*FuncEntry
}

// VarEntry representa una variable con su nombre, tipo y dirección.
type VarEntry struct {
	Name   string
	Type   Tipo
	DirInt int // dirección en memoria
}

// VarTable gestiona un ámbito de variables, indexadas por dirección.
type VarTable struct {
	vars   map[int]*VarEntry // DirInt → VarEntry
	parent *VarTable         // ámbito superior
}
