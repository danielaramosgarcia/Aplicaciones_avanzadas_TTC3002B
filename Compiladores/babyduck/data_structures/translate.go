package data_structures

import "fmt"

/*
Traduccion de tipo a su representacion en numero

	int -> 0
	float -> 1
	bool -> 2
*/
func TranslateType(typ string) Tipo {
	switch typ {
	case "int":
		return 0
	case "float":
		return 1
	case "bool":
		return 2
	default:
		return -1
	}
}

/*
Traduccion de operadores a su representacion en numero
  - -> 10
  - -> 20
  - -> 30
    /  -> 40
    <  -> 50
    >  -> 60
    != -> 70
    ( -> 80
    ) -> 90
    GOTO -> 90
    GOTOFALSE -> 91
    = -> 92
*/
const (
	ADD       = 10
	SUB       = 20
	MUL       = 30
	DIV       = 40
	LT        = 50
	GT        = 60
	NEQ       = 70
	RPAR      = 80
	GOTO      = 90 // salto incondicional
	GOTOFALSE = 91 // salto si condición es falsa
	GOTOTRUE  = 91 // salto si condición es falsa
	EQ        = 92 // asignación
	PRINT     = 93 // imprimir
)

// TanslateOp recibe un operador como string y devuelve su representación numérica
func (ctx *Context) TranslateOp(op string) (interface{}, error) {
	switch op {
	case "+":
		ctx.PushOperator(ADD)
		return nil, nil
	case "-":
		ctx.PushOperator(SUB)
		return nil, nil
	case "*":
		ctx.PushOperator(MUL)
		return nil, nil
	case "/":
		ctx.PushOperator(DIV)
		return nil, nil
	case "<":
		ctx.PushOperator(LT)
		return nil, nil
	case ">":
		ctx.PushOperator(GT)
		return nil, nil
	case "!=":
		ctx.PushOperator(NEQ)
		return nil, nil
	case "(":
		ctx.PushOperator(RPAR) // Asumimos que "(" es un operador especial
		return nil, nil
	case ")":
		return ctx.HandleRightParen()
		// return nil, nil
	default:
		return nil, fmt.Errorf("operador desconocido: %s", op)
	}
}

// Translate back fron number to string
func TranslateBackOp(op int) string {
	switch op {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case LT:
		return "<"
	case GT:
		return ">"
	case NEQ:
		return "!="
	case RPAR:
		return "("
	case GOTO:
		return "GOTO"
	case GOTOFALSE:
		return "GOTOFALSE"
	case EQ:
		return "="
	case PRINT:
		return "PRINT"
	default:
		return ""
	}
}
