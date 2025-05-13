package data_structures

import (
	"fmt"
)

// Context mantiene el estado global del compilador:
// - GlobalVars: tabla global de variables
// - FuncDir: directorio de funciones
// - currentFunc: la función activa (no anidada)
// En BabyDuck no hay funciones anidadas ni scope más profundo que el de funciones.

type Context struct {
	GlobalVars  *VarTable
	FuncDir     *FuncDir
	currentFunc *FuncEntry
}

// Param representa un parámetro de función: nombre y tipo.
// Se especifica el tipo Param para poder reconocer el arreglo
// de Param cuando se reconoce una funcion en la gramática.
type Param struct {
	Name string
	Type Tipo
}

// NewContext inicializa el contexto con tablas vacías.
func NewContext() *Context {
	return &Context{
		GlobalVars:  NewVarTable(nil),
		FuncDir:     NewFuncDir(),
		currentFunc: nil,
	}
}

// AddGlobalVar agrega una variable al ámbito global, error si ya existe.
func (ctx *Context) AddGlobalVar(name string, typ Tipo) error {
	return ctx.GlobalVars.Add(name, typ)
}

// AddFunction registra una función en el directorio global.
// Crea la entrada con su propia tabla local para variables.
func (ctx *Context) AddFunction(name string, ret Tipo, params []Param) error {
	// Prepara tabla de variables locales
	vt := NewVarTable(ctx.GlobalVars)

	// Crea la entrada de función
	f := &FuncEntry{
		Name:       name,
		ReturnType: ret,
		ParamTypes: []Tipo{},
		VarTable:   vt,
	}

	// Agrega al directorio
	if err := ctx.FuncDir.Add(f); err != nil {
		return err
	}
	// Inserta parámetros en la tabla local
	for _, p := range params {
		f.ParamTypes = append(f.ParamTypes, p.Type)
		if err := f.VarTable.Add(p.Name, p.Type); err != nil {
			return fmt.Errorf("parámetro %s: %w", p.Name, err)
		}
	}
	return nil
}

// EnterFunction activa el contexto local de la función nombrada.
func (ctx *Context) EnterFunction(name string) (interface{}, error) {
	f, ok := ctx.FuncDir.Get(name)
	if !ok {
		return nil, fmt.Errorf("función %s no declarada", name)
	}
	ctx.currentFunc = f
	return nil, nil
}

// ExitFunction cierra el scope de la función activa.
func (ctx *Context) ExitFunction() (interface{}, error) {
	ctx.currentFunc = nil
	return nil, nil
}

// CurrentVarTable devuelve la tabla de variables del scope actual,
// que será la tabla local si estamos dentro de una función, o la global en caso contrario.
func (ctx *Context) CurrentVarTable() *VarTable {
	if ctx.currentFunc != nil {
		return ctx.currentFunc.VarTable
	}
	return ctx.GlobalVars
}

// AddLocalVar agrega una variable a la tabla de la función activa.
func (ctx *Context) AddLocalVar(name string, typ Tipo) error {
	return ctx.currentFunc.VarTable.Add(name, typ)
}

// ReturnContext devuelve el contexto completo al finalizar el parseo.
func (ctx *Context) ReturnContext() (interface{}, error) {
	return ctx, nil
}

// RegisterGlobalVars añade varias variables globales de un mismo tipo.
func (ctx *Context) RegisterGlobalVars(names []string, typ Tipo) (interface{}, error) {
	for _, name := range names {
		if err := ctx.AddGlobalVar(name, typ); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// RegisterFunction registra una función (firma) en el directorio de funciones.
func (ctx *Context) RegisterFunction(name string, ret Tipo, params []Param) (interface{}, error) {
	err := ctx.AddFunction(name, ret, params)
	return nil, err
}

// MakeVarList construye un slice con un único identificador.
func MakeVarList(name string) (interface{}, error) {
	return []string{name}, nil
}

// ConcatVarList agrega un identificador al frente de una lista existente.
func ConcatVarList(head string, tail []string) (interface{}, error) {
	return append([]string{head}, tail...), nil
}

// PrependParam añade un nuevo Param al frente de la lista existente.
func PrependParam(name string, typ Tipo, tail []Param) (interface{}, error) {
	return append([]Param{{Name: name, Type: typ}}, tail...), nil
}

// MakeParam construye un slice de Param con un solo elemento.
func MakeParam(name string, typ Tipo) (interface{}, error) {
	return []Param{{Name: name, Type: typ}}, nil
}

// TODO VER SI LO NECESITAMOS
// MakeParamList construye un slice de parámetros (tipos) con un único elemento.
func MakeParamList(param Tipo) (interface{}, error) {
	return []Tipo{param}, nil
}

// ConcatParamList concatena un tipo de parámetro al frente de la lista existente.
func ConcatParamList(head Tipo, tail []Tipo) (interface{}, error) {
	return append([]Tipo{head}, tail...), nil
}

// ValidateAssign comprueba en tiempo de parseo que la variable exista y que el tipo coincida.
func (ctx *Context) ValidateAssign(name string, typ Tipo) (interface{}, error) {
	entry, ok := ctx.CurrentVarTable().Get(name)
	if !ok {
		return nil, fmt.Errorf("variable %s no declarada", name)
	}
	if entry.Type != typ {
		return nil, fmt.Errorf("tipos incompatibles en asignación: %v != %v", entry.Type, typ)
	}
	return nil, nil
}

// ResolveVarType consulta el tipo de una variable en el contexto actual.
func (ctx *Context) ResolveVarType(name string) (interface{}, error) {
	entry, ok := ctx.CurrentVarTable().Get(name)
	if !ok {
		return nil, fmt.Errorf("variable %s no declarada", name)
	}
	return entry.Type, nil
}

// ReturnExpression simplemente devuelve el tipo inferido de una subexpresión.
func ReturnExpression(expr Tipo) (interface{}, error) {
	return expr, nil
}

// Reset reinicia el contexto para un nuevo parseo,
// limpiando la tabla global y el directorio de funciones.
func (ctx *Context) Reset() (interface{}, error) {
	ctx.GlobalVars = NewVarTable(nil)
	ctx.FuncDir = NewFuncDir()
	ctx.currentFunc = nil
	return ctx, nil
}

func (ctx *Context) RegisterAndEnterFunction(
	name string,
	ret Tipo,
	params []Param,
) (interface{}, error) {
	// 1) registra firma + parámetros
	if _, err := ctx.RegisterFunction(name, ret, params); err != nil {
		return nil, err
	}
	// 2) activa currentFunc
	return ctx.EnterFunction(name)
}
