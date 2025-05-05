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

// Representa la firma y el ámbito de una función
type FuncEntry struct {
	Name       string    // identificador
	ReturnType Tipo      // enum { Int, Float, Bool, String, Void }
	ParamTypes []Tipo    // tipos de parámetros en orden
	VarTable   *VarTable // tabla de variables locales
}

// Contenedor global de todas las funciones
type FuncDir struct {
	funcs map[string]*FuncEntry
}

// Representa una variable declarada
type VarEntry struct {
	Name string
	Type Tipo
}

// Tabla de variables de un ámbito (global o función)
type VarTable struct {
	vars   map[string]*VarEntry
	parent *VarTable // para encadenar a tabla global
}
