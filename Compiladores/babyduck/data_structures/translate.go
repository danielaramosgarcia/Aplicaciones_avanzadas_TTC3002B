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
	ADD       = 10 // suma
	SUB       = 20 // resta
	MUL       = 30 // multiplicación
	DIV       = 40 // división
	LT        = 50 // menor que
	GT        = 60 // mayor que
	NEQ       = 70 // no igual
	RPAR      = 80 // paréntesis derecho
	GOTO      = 90 // salto incondicional
	GOTOFALSE = 91 // salto si condición es falsa
	EQ        = 92 // asignación
	PRINT     = 93 // imprimir
	ENDF      = 94 // fin de función
	ERA       = 95 // inicio de función
	PARAM     = 96 // parámetro de función
	GOSUB     = 97 // llamada a función
	END       = 98 // fin del programa
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
	case ENDF:
		return "ENDF"
	case ERA:
		return "ERA"
	case PARAM:
		return "PARAM"
	case GOSUB:
		return "GOSUB"
	case END:
		return "END"
	default:
		return ""
	}
}
