package data_structures

func (q *QuadQueue) Enqueue(qd Quadruple) { q.Quads = append(q.Quads, qd) }
func (q *QuadQueue) List() []Quadruple    { return q.Quads }

// PushOperator apila un operador en OperatorStack.
func (ctx *Context) PushOperator(op int) {
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
func (ctx *Context) GenerateQuad() error {
	// 1) Sacar operandos y tipos
	rightOp := ctx.PopOperand()
	rightType := ctx.PopType()
	leftOp := ctx.PopOperand()
	leftType := ctx.PopType()
	op := ctx.PopOperator()

	// 2) Chequeo semántico: obtener resultado del cubo
	resType, err := ResultBinary(op, leftType, rightType)
	if err != nil {
		return err
	}

	// 3) Crear temporal y actualizar pilas
	// temp := fmt.Sprintf("t%d", ctx.TempCounter)
	temp := ctx.TempCounter
	ctx.TempCounter++
	ctx.PushOperand(temp)
	ctx.PushType(resType)

	// 4) Encolar cuádruplo
	ctx.Quads.Enqueue(Quadruple{
		Op:     op,
		Arg1:   leftOp,
		Arg2:   rightOp,
		Result: temp,
	})
	return nil
}

// HandleOperand registra un operando (id o literal) y su tipo en las pilas.
func (ctx *Context) HandleOperand(lexeme int, typ int) {
	ctx.OperandStack = append(ctx.OperandStack, lexeme)
	ctx.TypeStack = append(ctx.TypeStack, typ)
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

	// 3) Temporal
	temp := ctx.TempCounter
	ctx.TempCounter++
	ctx.OperandStack = append(ctx.OperandStack, temp)
	ctx.TypeStack = append(ctx.TypeStack, resTyp)

	// 4) Cuádruplo
	ctx.Quads.Enqueue(Quadruple{Op: op, Arg1: leftOp, Arg2: rightOp, Result: temp})
	return nil
}

// HandleRelational funciona igual que HandleBinary, pero para <, >, !=
func (ctx *Context) HandleRelational(op int) error {
	return ctx.HandleBinary(op)
}
