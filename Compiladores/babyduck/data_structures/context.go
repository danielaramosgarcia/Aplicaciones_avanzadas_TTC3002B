package data_structures

import (
	"fmt"
)

// NewContext inicializa el contexto con tablas vacías.
func NewContext() *Context {
	return &Context{
		FuncDir:       NewFuncDir(),
		currentFunc:   nil,
		OperatorStack: make([]int, 0),
		OperandStack:  make([]int, 0),
		TypeStack:     make([]int, 0),
		Quads:         QuadQueue{Quads: make([]Quadruple, 0)},
		FuncCount:     0,
		FuncIndex:     make(map[int]string),
		FuncSignature: &FuncSignature{
			ParamSignature: make([]int, 0),
			ParamLength:    0,
		},
		ConstTable: Constable{
			Num: make(map[int]int),
			Str: make(map[int]string),
			Bool: map[int]bool{
				0: false,
				1: true,
			},
			Float: make(map[int]float64),
		},
		AddedConst: make([]string, 0),
	}
}

// AddFunction registra una función en el directorio global.
// Crea la entrada con su propia tabla local para variables.
func (ctx *Context) AddFunction(name string, ret int, params []Param) error {
	// Prepara tabla de variables locales
	vt := NewVarTable(nil)

	if ret == 5 {
		programName = name
	} else {
		vt.parent = ctx.FuncDir.funcs[programName].VarTable
	}

	// Crea la entrada de función
	f := &FuncEntry{
		Name:       name,
		ReturnType: ret,
		ParamTypes: []int{},
		VarTable:   vt,
		Space: &SpaceVariables{
			Var:  0, // Inicializa espacio de variables
			Temp: 0, // Inicializa espacio de temporales
		},
		index:     ctx.FuncCount,
		CuadStart: len(ctx.Quads.Quads), // Marca el inicio de los cuádruplos para esta función
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
	ctx.FuncIndex[ctx.FuncCount] = name
	ctx.FuncCount++
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

func (ctx *Context) ResetCounters() {
	nextGlobalIntAddr = GlobalIntBase
	nextGlobalFloatAddr = GlobalFloatBase

	nextLocalIntAddr = LocalIntBase
	nextLocalFloatAddr = LocalFloatBase

	nextTempIntAddr = TempIntBase
	nextTempFloatAddr = TempFloatBase
	nextTempBoolAddr = TempBoolBase

	nextConstIntAddr = ConstIntBase
	nextConstFloatAddr = ConstFloatBase
	nextConstStringAddr = ConstStringBase
	ctx.FuncSignature = &FuncSignature{
		ParamSignature: make([]int, 0),
		ParamLength:    0,
	}
}

// ExitFunction cierra el scope de la función activa.
func (ctx *Context) ExitFunction() (interface{}, error) {
	ctx.currentFunc = ctx.FuncDir.funcs[programName]
	ctx.MakeEndFQuad()

	ctx.ResetCounters()
	return nil, nil
}

// CurrentVarTable devuelve la tabla de variables del scope actual,
// que será la tabla local si estamos dentro de una función, o la global en caso contrario.
func (ctx *Context) CurrentVarTable() *VarTable {

	if ctx.currentFunc != nil {
		return ctx.currentFunc.VarTable
	}
	return ctx.FuncDir.funcs[programName].VarTable
}

// ReturnContext devuelve el contexto completo al finalizar el parseo.
func (ctx *Context) ReturnContext() (interface{}, error) {
	fmt.Println(" ____________________________________________________")
	fmt.Printf("\n| Operador | Operando Izq | Operando Der | Resultado |\n")
	fmt.Println(" ---------------------------------------------------- ")
	for i, quad := range ctx.Quads.Quads {
		fmt.Printf("|  Quad %d: | %s | %d | %d | %d | \n", i, TranslateBackOp(quad.Op), quad.Arg1, quad.Arg2, quad.Result)
	}
	fmt.Printf("\nLista de operandos: %v\n", ctx.OperandStack)
	fmt.Printf("\nLista de operadores: %v\n", ctx.OperatorStack)
	fmt.Printf("\n Mapa de FuncCount: %v\n", ctx.FuncIndex)
	fmt.Printf("\n Mapa de Constantes int: %v\n", ctx.ConstTable.Num)
	fmt.Printf("\n Mapa de Constantes float: %v\n", ctx.ConstTable.Float)
	fmt.Printf("\n Mapa de Constantes str: %v\n", ctx.ConstTable.Str)
	fmt.Printf("\n Funciones y sus propiedades:\n")
	for name, entry := range ctx.FuncDir.funcs {
		fmt.Println(" ____________________________________________________")
		fmt.Printf("Función %s: Retorno=%d, Parámetros=%v, Variables=%d, Temporales=%d CuadStart=%d  \n",
			name, entry.ReturnType, entry.ParamTypes, entry.Space.Var, entry.Space.Temp, entry.CuadStart)
		fmt.Printf("Tabla de variables de %s:\n", name)
		for _, entry := range entry.VarTable.vars {
			fmt.Printf("  - %s: Dir=%d, Tipo=%d\n", entry.Name, entry.DirInt, entry.Type)
		}
		fmt.Println(" ____________________________________________________")

	}
	return ctx, nil
}

// TODO: CAMBIAR NOMBRE EN EL BNF Y ACA PARA Q SEA LA UNICA DE AGREGAR VARS.
// RegisterVars añade varias variables globales de un mismo tipo.
func (ctx *Context) RegisterVars(names []string, typ int) (interface{}, error) {
	for _, name := range names {
		if err := ctx.currentFunc.VarTable.Add(name, typ); err != nil {
			return nil, err
		}
		ctx.currentFunc.Space.Var++
	}
	return nil, nil
}

// RegisterTemp una variable temporal a la tabla actual y suma su contador.
func (ctx *Context) RegisterTemp(typ int) (int, error) {
	dir, err := ctx.currentFunc.VarTable.AddTemp(typ)
	if err != nil {
		return dir, err
	}
	ctx.currentFunc.Space.Temp++
	return dir, nil
}

// RegisterFunction registra una función (firma) en el directorio de funciones.
func (ctx *Context) RegisterFunction(name string, ret int, params []Param) (interface{}, error) {
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
func PrependParam(name string, typ int, tail []Param) (interface{}, error) {
	return append([]Param{{Name: name, Type: typ}}, tail...), nil
}

// MakeParam construye un slice de Param con un solo elemento.
func MakeParam(name string, typ int) (interface{}, error) {
	return []Param{{Name: name, Type: typ}}, nil
}

// ConcatParamList concatena un tipo de parámetro al frente de la lista existente.
func ConcatParamList(head int, tail []int) (interface{}, error) {
	return append([]int{head}, tail...), nil
}

// GetByName busca una variable por su identificador (Name) recorriendo las entradas
// en este ámbito y, si no la encuentra, en los padres. Devuelve VarEntry y true si existe.
func (vt *VarTable) GetByName(name string) (*VarEntry, bool) {
	// 1) Busca en este ámbito
	for _, entry := range vt.vars {
		if entry.Name == name {
			return entry, true
		}
	}
	// 2) Si no está aquí, y hay padre, busca recursivamente
	if vt.parent != nil {
		return vt.parent.GetByName(name)
	}
	// 3) No existe en ningún ámbito
	return nil, false
}

// ValidateAssign comprueba en tiempo de parseo que la variable identificada por name
// existe y que su tipo coincida. Internamente usa la dirección asignada en VarEntry.
func (ctx *Context) ValidateAssign(name string, typ int) (interface{}, error) {
	vt := ctx.CurrentVarTable()
	entry, ok := vt.GetByName(name)
	if !ok {
		return nil, fmt.Errorf("variable %q no declarada", name)
	}
	if entry.Type != typ {
		return nil, fmt.Errorf(
			"MISMATCH ERR: Tipos incompatibles en asignación de %q: %v != %v",
			name, entry.Type, typ,
		)
	}
	// Empuja el operando y su tipo en las pilas
	ctx.HandleOperand(entry.DirInt, entry.Type)
	// Crear quad de validate
	ctx.AssignQuad()

	return nil, nil
}

// ResolveVarType consulta el tipo de la variable identificada por name.
func (ctx *Context) ResolveVarType(name string) (interface{}, error) {
	vt := ctx.CurrentVarTable()
	entry, ok := vt.GetByName(name)
	if !ok {
		return nil, fmt.Errorf("variable %q no declarada", name)
	}
	// Empuja el operando y su tipo en las pilas
	ctx.HandleOperand(entry.DirInt, entry.Type)
	return entry.Type, nil
}

// TODO CONVERTIR A INT PARA ALMACENAR VALOR
// ResolveVarType consulta el tipo de la variable identificada por name.
func (ctx *Context) ResolveCteInt(cte string) (interface{}, error) {

	dir, err := ctx.AddConst(0, cte)
	if err != nil {
		return nil, fmt.Errorf("error al agregar constante: %w", err)
	}
	// Empuja el operando y su tipo en las pilas
	ctx.HandleOperand(dir, 0)
	return Int, nil
}

// TODO CONVERTIR A FLOAT PARA ALMACENAR VALOR
// ResolveVarType consulta el tipo de la variable identificada por name.
func (ctx *Context) ResolveCteFloat(cte string) (interface{}, error) {

	dir, err := ctx.AddConst(1, cte)
	if err != nil {
		return nil, fmt.Errorf("error al agregar constante: %w", err)
	}
	// Empuja el operando y su tipo en las pilas
	ctx.HandleOperand(dir, 1)
	return Float, nil
}

// TODO CONVERTIR A FLOAT PARA ALMACENAR VALOR
// ResolveVarType consulta el tipo de la variable identificada por name.
func (ctx *Context) ResolveCteSting(cte string) (interface{}, error) {

	dir, err := ctx.AddConst(4, cte)
	if err != nil {
		return nil, fmt.Errorf("error al agregar constante: %w", err)
	}
	// Empuja el operando y su tipo en las pilas
	ctx.HandleOperand(dir, 4)
	return String, nil
}

// ReturnExpression simplemente devuelve el tipo inferido de una subexpresión.
func ReturnExpression(expr int) (interface{}, error) {
	return expr, nil
}

// Reset reinicia el contexto para un nuevo parseo,
// limpiando la tabla global y el directorio de funciones.
func (ctx *Context) Reset() (interface{}, error) {
	ctx.FuncDir = NewFuncDir()
	ctx.currentFunc = nil
	return ctx, nil
}

func (ctx *Context) RegisterAndEnterFunction(
	name string,
	ret int,
	params []Param,
) (interface{}, error) {
	// 1) registra firma + parámetros
	if _, err := ctx.RegisterFunction(name, ret, params); err != nil {
		return nil, err
	}
	// ctx.FuncDir.funcs[name].CuadStart = len(ctx.Quads.Quads)
	// ctx.FuncDir.funcs[name].index = ctx.FuncCount
	// ctx.
	// 2) activa currentFunc
	return ctx.EnterFunction(name)
}

func (ctx *Context) RegisterProgramId(name string) (interface{}, error) {
	if _, err := ctx.RegisterFunction(name, NP, nil); err != nil {
		return nil, err
	}
	return ctx.EnterFunction(name)
}

func (ctx *Context) FunctionCall(id string) (interface{}, error) {
	if _, ok := ctx.FuncDir.Get(id); !ok {
		return nil, fmt.Errorf("función %s no declarada", id)
	}
	ctx.FuncSignature.ParamLength = len(ctx.FuncDir.funcs[id].ParamTypes)
	ctx.FuncSignature.ParamSignature = ctx.FuncDir.funcs[id].ParamTypes
	ctx.MakeEraQuad(id)
	return id, nil
}

// FunctionCallEnd checa si hay parametros pendientes
func (ctx *Context) FunctionCallEnd(id string) (interface{}, error) {
	if len(ctx.FuncSignature.ParamSignature) > 0 {
		return nil, fmt.Errorf("ERR Le faltan %d parámetros a la funcion, se esperaban %d, pero se recibieron %d",
			len(ctx.FuncSignature.ParamSignature), ctx.FuncSignature.ParamLength, ctx.FuncSignature.ParamLength-len(ctx.FuncSignature.ParamSignature))
	}
	ctx.MakeGOSUBQuad(id)
	return nil, nil
}
