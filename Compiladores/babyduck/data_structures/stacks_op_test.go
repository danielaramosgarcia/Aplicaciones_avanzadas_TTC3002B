package data_structures

import "testing"

func TestQuadrupleGeneration(t *testing.T) {
	ctx := NewContext()

	// 1) Generar cuádruplo para 2 + 3
	ctx.HandleOperand(1601, 0)
	ctx.HandleOperand(1602, 0)
	ctx.PushOperator(10)
	if err := ctx.HandleBinary(10); err != nil {
		t.Fatalf("HandleBinary(+) falló: %v", err)
	}

	t.Logf("=== Después de 2+3 ===")
	t.Logf("OperandStack: %v", ctx.OperandStack)
	t.Logf("TypeStack:    %v", ctx.TypeStack)
	t.Logf("Cuádruplos:")
	for i, q := range ctx.Quads.List() {
		t.Logf("%2d: (%d, %d, %d, %d)", i, q.Op, q.Arg1, q.Arg2, q.Result)
	}

	// 2) Generar cuádruplo para t1 * 4
	// Tomamos el resultado previo (último operando)
	// agregar temporales
	last := ctx.OperandStack[len(ctx.OperandStack)-1]
	//TODO: AGREGAR GENERACION DE TEMPORALES
	ctx.HandleOperand(last, 0)
	ctx.HandleOperand(1603, 0)
	ctx.PushOperator(30)
	if err := ctx.HandleBinary(30); err != nil {
		t.Fatalf("HandleBinary(*) falló: %v", err)
	}

	t.Logf("=== Después de %d *4 ===", last)
	t.Logf("OperandStack: %v", ctx.OperandStack)
	t.Logf("TypeStack:    %v", ctx.TypeStack)
	t.Logf("Cuádruplos:")
	for i, q := range ctx.Quads.List() {
		t.Logf("%2d: (%d, %d, %d, %d)", i, q.Op, q.Arg1, q.Arg2, q.Result)
	}
}

func TestRelationalQuadruple(t *testing.T) {
	ctx := NewContext()

	// Generar cuádruplo para 5 < 6
	ctx.HandleOperand(1601, 0)
	ctx.HandleOperand(1602, 0)
	ctx.PushOperator(50)
	if err := ctx.HandleBinary(50); err != nil {
		t.Fatalf("HandleBinary(<) falló: %v", err)
	}

	t.Logf("=== Después de 5<6 ===")
	t.Logf("OperandStack: %v", ctx.OperandStack)
	t.Logf("TypeStack:    %v", ctx.TypeStack)
	t.Logf("Cuádruplos:")
	for i, q := range ctx.Quads.List() {
		t.Logf("%2d: (%d, %d, %d, %d)", i, q.Op, q.Arg1, q.Arg2, q.Result)
	}
}

func TestMixedPrecedence(t *testing.T) {
	ctx := NewContext()

	// 3 * 4 primero
	ctx.HandleOperand(1601, 0)
	ctx.HandleOperand(1602, 0)
	ctx.PushOperator(30)
	if err := ctx.HandleBinary(30); err != nil {
		t.Fatalf("HandleBinary(*) falló: %v", err)
	}

	// Luego suma con 2: t1 + 2
	t1 := ctx.OperandStack[len(ctx.OperandStack)-1]
	ctx.HandleOperand(t1, 0)
	ctx.HandleOperand(1603, 0)
	ctx.PushOperator(10)
	if err := ctx.HandleBinary(10); err != nil {
		t.Fatalf("HandleBinary(+) falló: %v", err)
	}

	t.Log("OperandStack:", ctx.OperandStack)
	t.Log("TypeStack:   ", ctx.TypeStack)
	t.Log("Cuádruplos:")
	for i, q := range ctx.Quads.List() {
		t.Logf("%2d: (%d, %d, %d, %d)", i, q.Op, q.Arg1, q.Arg2, q.Result)
	}
}

// func TestChainAddition(t *testing.T) {
// 	ctx := NewContext()

// 	// (1 + 2) + 3
// 	ctx.HandleOperand("1", Int)
// 	ctx.HandleOperand("2", Int)
// 	ctx.PushOperator("+")
// 	if err := ctx.HandleBinary("+"); err != nil {
// 		t.Fatalf("HandleBinary(+) falló: %v", err)
// 	}

// 	ctx.HandleOperand("3", Int)
// 	ctx.PushOperator("+")
// 	if err := ctx.HandleBinary("+"); err != nil {
// 		t.Fatalf("HandleBinary(+) falló: %v", err)
// 	}

// 	t.Log("OperandStack:", ctx.OperandStack)
// 	t.Log("TypeStack:   ", ctx.TypeStack)
// 	t.Log("Cuádruplos:")
// 	for i, q := range ctx.Quads.List() {
// 		t.Logf("%2d: (%s, %s, %s, %s)", i, q.Op, q.Arg1, q.Arg2, q.Result)
// 	}
// }

// func TestIntFloatPromotion(t *testing.T) {
// 	ctx := NewContext()

// 	// 2 + 3.14 debería promover a Float y no fallar
// 	ctx.HandleOperand("2", Int)
// 	ctx.HandleOperand("3.14", Float)
// 	ctx.PushOperator("+")
// 	err := ctx.HandleBinary("+")
// 	if err != nil {
// 		t.Fatalf("HandleBinary(+) inesperado error: %v", err)
// 	}
// 	// Verificamos el tipo en la cima de TypeStack
// 	topType := ctx.TypeStack[len(ctx.TypeStack)-1]
// 	if topType != Float {
// 		t.Errorf("TypeStack tope = %v; quiero %v", topType, Float)
// 	}
// }

// func TestTypeMismatch(t *testing.T) {
// 	ctx := NewContext()

// 	// Intento inválido: Bool + Int
// 	ctx.HandleOperand("true", Bool)
// 	ctx.HandleOperand("1", Int)
// 	ctx.PushOperator("+")
// 	err := ctx.HandleBinary("+")
// 	if err == nil {
// 		t.Errorf("HandleBinary(+) debería fallar por mismatch Bool+Int")
// 	} else {
// 		t.Logf("Error esperado: %v", err)
// 	}
// }

// func TestResultBinary_Valid(t *testing.T) {
// 	cases := []struct {
// 		op    string
// 		left  Tipo
// 		right Tipo
// 		want  Tipo
// 	}{
// 		{"+", Int, Int, Int},
// 		{"+", Int, Float, Float},
// 		{"+", Float, Int, Float},
// 		{"+", Float, Float, Float},
// 		{"-", Int, Int, Int},
// 		{"-", Float, Int, Float},
// 		{"*", Int, Float, Float},
// 		{"/", Int, Int, Float}, // integer division promotes to float
// 		{"/", Float, Float, Float},
// 		{"<", Int, Int, Bool},
// 		{">", Float, Int, Bool},
// 		{"!=", Int, Float, Bool},
// 	}
// 	for _, c := range cases {
// 		got, err := ResultBinary(c.op, c.left, c.right)
// 		if err != nil {
// 			t.Errorf("ResultBinary(%q, %v, %v) unexpected error: %v", c.op, c.left, c.right, err)
// 			continue
// 		}
// 		if got != c.want {
// 			t.Errorf("ResultBinary(%q, %v, %v) = %v; want %v", c.op, c.left, c.right, got, c.want)
// 		}
// 	}
// }

// func TestResultBinary_Invalid(t *testing.T) {
// 	cases := []struct {
// 		op    string
// 		left  Tipo
// 		right Tipo
// 	}{
// 		{"+", Bool, Int},
// 		{"/", Bool, Bool},
// 		{"<", Bool, Int},
// 		{"unknown", Int, Int},
// 		{"*", Int, Bool},
// 	}
// 	for _, c := range cases {
// 		_, err := ResultBinary(c.op, c.left, c.right)
// 		if err == nil {
// 			t.Errorf("ResultBinary(%q, %v, %v) expected error, got nil", c.op, c.left, c.right)
// 		}
// 	}
// }

// func TestResultUnary_Valid(t *testing.T) {
// 	cases := []struct {
// 		op      string
// 		operand Tipo
// 		want    Tipo
// 	}{
// 		{"+", Int, Int},
// 		{"+", Float, Float},
// 		{"-", Int, Int},
// 		{"-", Float, Float},
// 		{"!", Bool, Bool},
// 	}
// 	for _, c := range cases {
// 		got, err := ResultUnary(c.op, c.operand)
// 		if err != nil {
// 			t.Errorf("ResultUnary(%q, %v) unexpected error: %v", c.op, c.operand, err)
// 			continue
// 		}
// 		if got != c.want {
// 			t.Errorf("ResultUnary(%q, %v) = %v; want %v", c.op, c.operand, got, c.want)
// 		}
// 	}
// }

func TestResultUnary_Invalid(t *testing.T) {
	cases := []struct {
		op      int
		operand int
	}{
		{-1, 2},
		{10, 2},
	}
	for _, c := range cases {
		_, err := ResultUnary(c.op, c.operand)
		if err == nil {
			t.Errorf("ResultUnary(%q, %v) expected error, got nil", c.op, c.operand)
		}
	}
}

// func TestComplexExpression(t *testing.T) {
// 	ctx := NewContext()

// 	// Simular: (1 + 2) * (3 - 4) / 2
// 	// Paso 1: 1 + 2
// 	ctx.HandleOperand("1", Int)
// 	ctx.HandleOperand("2", Int)
// 	ctx.PushOperator("+")
// 	if err := ctx.HandleBinary("+"); err != nil {
// 		t.Fatalf("1+2 falló: %v", err)
// 	}
// 	t1 := ctx.OperandStack[len(ctx.OperandStack)-1]

// 	// Paso 2: 3 - 4
// 	ctx.HandleOperand("3", Int)
// 	ctx.HandleOperand("4", Int)
// 	ctx.PushOperator("-")
// 	if err := ctx.HandleBinary("-"); err != nil {
// 		t.Fatalf("3-4 falló: %v", err)
// 	}
// 	t2 := ctx.OperandStack[len(ctx.OperandStack)-1]

// 	// Paso 3: t1 * t2
// 	ctx.HandleOperand(t1, Int)
// 	ctx.HandleOperand(t2, Int)
// 	ctx.PushOperator("*")
// 	if err := ctx.HandleBinary("*"); err != nil {
// 		t.Fatalf("%s*%s falló: %v", t1, t2, err)
// 	}
// 	t3 := ctx.OperandStack[len(ctx.OperandStack)-1]

// 	// Paso 4: t3 / 2
// 	ctx.HandleOperand(t3, Int)
// 	ctx.HandleOperand("2", Int)
// 	ctx.PushOperator("/")
// 	if err := ctx.HandleBinary("/"); err != nil {
// 		t.Fatalf("%s/2 falló: %v", t3, err)
// 	}

// 	// Validaciones finales
// 	if len(ctx.Quads.List()) != 4 {
// 		t.Errorf("Se esperaban 4 cuádruplos, se obtuvieron %d", len(ctx.Quads.List()))
// 	}
// 	// El tipo final debe ser Float, por la división Int/Int → Float
// 	finalType := ctx.TypeStack[len(ctx.TypeStack)-1]
// 	if finalType != Float {
// 		t.Errorf("Tipo final = %v; se esperaba %v", finalType, Float)
// 	}
// 	t.Logf("OperandStack final: %v", ctx.OperandStack)
// 	t.Logf("TypeStack final:    %v", ctx.TypeStack)
// 	t.Logf("Cuádruplos generados:")
// 	for i, q := range ctx.Quads.List() {
// 		t.Logf("%2d: (%s, %s, %s, %s)", i, q.Op, q.Arg1, q.Arg2, q.Result)
// 	}
// }
