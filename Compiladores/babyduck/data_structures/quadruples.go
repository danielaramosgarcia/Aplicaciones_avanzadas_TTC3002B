package data_structures

import "fmt"

func (q *QuadQueue) Enqueue(qd Quadruple) { q.Quads = append(q.Quads, qd) }
func (q *QuadQueue) List() []Quadruple    { return q.Quads }

// PushOperator apila un operador en OperatorStack.
func (ctx *Context) PushOperator(op int) {
	// fmt.Printf("=== PushOperator: %d ===\n", op)
	ctx.OperatorStack = append(ctx.OperatorStack, op)
}

// PopOperator desapila y devuelve el operador tope.
func (ctx *Context) PopOperator() int {
	n := len(ctx.OperatorStack) - 1
	op := ctx.OperatorStack[n]
	ctx.OperatorStack = ctx.OperatorStack[:n]
	return op
}

// PushOperand apila un operando (nombre o temporal).
func (ctx *Context) PushOperand(operand int) {
	ctx.OperandStack = append(ctx.OperandStack, operand)
}

// PopOperand desapila un operando.
func (ctx *Context) PopOperand() int {
	n := len(ctx.OperandStack) - 1
	val := ctx.OperandStack[n]
	ctx.OperandStack = ctx.OperandStack[:n]
	return val
}

// PushType / PopType para la pila de tipos
func (ctx *Context) PushType(t int) {
	ctx.TypeStack = append(ctx.TypeStack, t)
}
func (ctx *Context) PopType() int {
	n := len(ctx.TypeStack) - 1
	t := ctx.TypeStack[n]
	ctx.TypeStack = ctx.TypeStack[:n]
	return t
}

// GenerateQuad procesa el tope de operator/operandas y encola un cuádruplo.
func (ctx *Context) GenerateQuad() (interface{}, error) {
	println("\n === ENTRO A GENERATE QUAD ===")
	// 1) Sacar operandos y tipos
	rightOp := ctx.PopOperand()
	fmt.Println("rightOp:", rightOp)
	rightType := ctx.PopType()
	fmt.Println("rightType:", rightType)
	leftOp := ctx.PopOperand()
	fmt.Println("leftOp:", leftOp)
	leftType := ctx.PopType()
	fmt.Println("leftType:", leftType)
	op := ctx.PopOperator()
	fmt.Println("op:", op)
	// 2) Chequeo semántico: obtener resultado del cubo
	resType, err := ResultBinary(op, leftType, rightType)
	if err != nil {
		return nil, fmt.Errorf("TYPE mismatch!: %s", err)
	}

	// 3) Agregar temp y actualizar pilas
	tempDir, err := ctx.RegisterTemp(resType)
	if err != nil {
		return nil, err
	}
	ctx.PushOperand(tempDir)
	ctx.PushType(resType)

	// 4) Encolar cuádruplo
	ctx.Quads.Enqueue(Quadruple{
		Op:     op,
		Arg1:   leftOp,
		Arg2:   rightOp,
		Result: tempDir,
	})
	fmt.Printf("Cuádruplo generado: (%d, %d, %d, %d)", op, leftOp, rightOp, tempDir)
	return resType, nil
}

// HandleOperand registra un operando (id o literal) y su tipo en las pilas.
func (ctx *Context) HandleOperand(dir int, typ int) {
	// fmt.Printf("=== HandleOperand: dir=%d, typ=%d ===\n", dir, typ)
	ctx.OperandStack = append(ctx.OperandStack, dir)
	// fmt.Printf("OperandStack en handleOperand: %v\n", ctx.OperandStack)
	ctx.TypeStack = append(ctx.TypeStack, typ)
	// fmt.Printf("TypeStack en handleOperand: %v\n", ctx.TypeStack)
	// fmt.Printf("OperadorStack en handleOperand: %v\n", ctx.OperatorStack)
}

// HandleBinary genera un cuádruplo para un operador binario.
// Desapila dos operandos+tipos, chequea en el cubo y encola el quad.
func (ctx *Context) HandleBinary(op int) error {
	// 1) Sacar operandos y tipos
	n := len(ctx.OperandStack) - 1
	rightOp := ctx.OperandStack[n]
	ctx.OperandStack = ctx.OperandStack[:n]
	rightTyp := ctx.TypeStack[n]
	ctx.TypeStack = ctx.TypeStack[:n]
	n = len(ctx.OperandStack) - 1
	leftOp := ctx.OperandStack[n]
	ctx.OperandStack = ctx.OperandStack[:n]
	leftTyp := ctx.TypeStack[n]
	ctx.TypeStack = ctx.TypeStack[:n]

	// 2) Semántica
	resTyp, err := ResultBinary(op, leftTyp, rightTyp)
	if err != nil {
		return err
	}

	// 3) Agregar temp y actualizar pilas
	tempDir, err := ctx.RegisterTemp(resTyp)
	if err != nil {
		return err
	}

	// 3) Temporal
	// temp := ctx.TempCounter
	// ctx.TempCounter++
	ctx.OperandStack = append(ctx.OperandStack, tempDir)
	ctx.TypeStack = append(ctx.TypeStack, resTyp)

	// 4) Cuádruplo
	ctx.Quads.Enqueue(Quadruple{Op: op, Arg1: leftOp, Arg2: rightOp, Result: tempDir})
	return nil
}

// HandleRelational funciona igual que HandleBinary, pero para <, >, !=
func (ctx *Context) HandleRelational(op int) error {
	return ctx.HandleBinary(op)
}

// HandleRightParen saca operadores hasta el "(" (código 80)
// y por cada uno invoca GenerateQuad para encolar su cuádruplo.
func (ctx *Context) HandleRightParen() (interface{}, error) {
	for {
		if len(ctx.OperatorStack) == 0 {
			return nil, fmt.Errorf("paréntesis desbalanceados")
		}
		op := ctx.PopOperator()
		if op == 80 { // 80 = código que usaste para "("
			break
		}
		// Como GenerateQuad espera tener el operador en la pila,
		// lo volvemos a empujar y dejamos que GenerateQuad lo saque:
		ctx.PushOperator(op)
		if _, err := ctx.GenerateQuad(); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// Fill auxiliar para llenar un cuádruplo pendiente con el índice actual.
// Se usa para completar saltos condicionales.
func (ctx *Context) FillJump() (interface{}, error) {
	println("\n === ENTRO A FILLJUMP ===")
	idx := len(ctx.Quads.Quads) // Índice del cuádruplo actual

	// Chequear si hay un salto pendiente
	if len(ctx.JumpStack) == 0 {
		return nil, fmt.Errorf("no hay salto pendiente para llenar")
	}
	// Sacar el índice del cuádruplo a llenar
	jumpIndex := ctx.JumpStack[len(ctx.JumpStack)-1]
	fmt.Printf("Índice del jump: %d\n", jumpIndex)
	ctx.JumpStack = ctx.JumpStack[:len(ctx.JumpStack)-1]
	// Actualizar el cuádruplo con el resultado
	ctx.Quads.Quads[jumpIndex].Result = idx
	fmt.Printf("Cuádruplo en índice %d actualizado con resultado %d\n", jumpIndex, idx)
	return nil, nil
}

func (ctx *Context) MakeGFQuad(typ int) (interface{}, error) {
	// Checar si es de tipo booleano
	println("\n === ENTRO A MAKEGFQUAD QUADS ===")
	// 1) Sacar operandos y tipos
	rightOp := ctx.PopOperand()
	fmt.Println("rightOp:", rightOp)
	rightType := ctx.PopType()
	fmt.Println("rightType:", rightType)
	if rightType != 2 { // 2 = tipo booleano
		return nil, fmt.Errorf("tipo %d no es booleano para condición if", rightType)
	}
	// Generar cuádruplo de salto condicional
	quad := Quadruple{
		Op:     GOTOFALSE,
		Arg1:   rightOp, // Último operando
		Arg2:   -1,      // No se usa
		Result: 0,       // Será llenado después
	}
	// 2) Agregar a la pila de saltos pendientes
	ctx.JumpStack = append(ctx.JumpStack, len(ctx.Quads.Quads)) // Índice del cuádruplo a llenar
	fmt.Printf("JumpStack actualizado: %v\n", ctx.JumpStack)
	ctx.Quads.Quads = append(ctx.Quads.Quads, quad)
	fmt.Printf("Cuádruplo generado: (%d, %d, %d, %d)\n", quad.Op, quad.Arg1, quad.Arg2, quad.Result)
	return nil, nil
}

func (ctx *Context) ElseJumpIf() (interface{}, error) {
	fmt.Println("\n=== ENTRO A ELSEJUMPIF ===")
	// Chequear si hay un salto pendiente
	if len(ctx.JumpStack) == 0 {
		return nil, fmt.Errorf("no hay salto pendiente para llenar")
	}
	// Sacar el índice del cuádruplo a llenar
	jumpIndex := ctx.JumpStack[len(ctx.JumpStack)-1]
	fmt.Printf("Índice del jump: %d\n", jumpIndex)
	ctx.JumpStack = ctx.JumpStack[:len(ctx.JumpStack)-1]
	// Actualizar el cuádruplo con el resultado
	ctx.Quads.Quads[jumpIndex].Result = len(ctx.Quads.Quads) + 1 // Actualizar con el índice actual

	// Generar cuádruplo de salto GOTO
	quad := Quadruple{
		Op:     GOTO,
		Arg1:   -1, // Último operando
		Arg2:   -1, // No se usa
		Result: 0,  // Será llenado después
	}

	// 2) Agregar a la pila de saltos pendientes
	ctx.JumpStack = append(ctx.JumpStack, len(ctx.Quads.Quads)) // Índice del cuádruplo a llenar
	ctx.Quads.Quads = append(ctx.Quads.Quads, quad)
	fmt.Printf("Cuádruplo generado: (%d, %d, %d, %d)\n", quad.Op, quad.Arg1, quad.Arg2, quad.Result)
	fmt.Printf("JumpStack actualizado: %v\n", ctx.JumpStack)

	return nil, nil
}

func (ctx *Context) CycleJump() (interface{}, error) {

	fmt.Println(" \n === CYCLEJUMP ===")
	// Sacar el índice del cuádruplo a llenar
	ctx.JumpStack = append(ctx.JumpStack, len(ctx.Quads.Quads)) // Índice del cuádruplo a llenar
	fmt.Printf("JumpStack actualizado: %v\n", ctx.JumpStack)

	return nil, nil
}

func (ctx *Context) WhileJump() (interface{}, error) {
	fmt.Println("\n=== ENTRO A WhileJump ===")
	// Chequear si hay un salto pendiente
	fmt.Printf("Lista de jumps pendientes: %v\n", ctx.JumpStack)

	if len(ctx.JumpStack) == 0 {
		return nil, fmt.Errorf("no hay salto pendiente para llenar")
	}

	quad := Quadruple{
		Op:     GOTO,
		Arg1:   -1,
		Arg2:   -1, // No se usa
		Result: 0,
	}
	ctx.Quads.Quads = append(ctx.Quads.Quads, quad)

	ctx.FillJump()
	fmt.Printf("Lista de jumps pendientes: %v\n", ctx.JumpStack)
	// Sacar indicie para cuad del final
	jumpIndexStart := ctx.JumpStack[len(ctx.JumpStack)-1]
	fmt.Printf("INDICE A DONDE CREO QUE SALTARA EN EL GO TO: %d\n", jumpIndexStart)
	ctx.JumpStack = ctx.JumpStack[:len(ctx.JumpStack)-1]

	// Actualizar el cuádruplo con el resultado
	ctx.Quads.Quads[len(ctx.Quads.Quads)-1].Result = jumpIndexStart // Actualizar al indice del while

	fmt.Printf("Lista de jumps pendientes: %v\n", ctx.JumpStack)
	fmt.Printf("Cuádruplo generado: (%d, %d, %d, %d)\n", quad.Op, quad.Arg1, quad.Arg2, quad.Result)

	return nil, nil
}
