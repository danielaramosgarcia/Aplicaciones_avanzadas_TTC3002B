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
	+ -> 10
	- -> 20
	* -> 30
    /  -> 40
    <  -> 50
    >  -> 60
    != -> 70
	( -> 80
	) -> 90
*/

// TanslateOp recibe un operador como string y devuelve su representación numérica
func (ctx *Context) TranslateOp(op string) (interface{}, error) {
	fmt.Println("=== ENTRO A TRANSLATEOP ===")
	fmt.Printf("Operador recibido: %s\n", op)
	switch op {
	case "+":
		ctx.PushOperator(10)
		return nil, nil
	case "-":
		ctx.PushOperator(20)
		return nil, nil
	case "*":
		ctx.PushOperator(30)
		return nil, nil
	case "/":
		ctx.PushOperator(40)
		return nil, nil
	case "<":
		ctx.PushOperator(50)
		return nil, nil
	case ">":
		ctx.PushOperator(60)
		return nil, nil
	case "!=":
		ctx.PushOperator(70)
		return nil, nil
	case "(":
		ctx.PushOperator(80) // Asumimos que "(" es un operador especial
		return nil, nil
	case ")":
		return ctx.HandleRightParen()
		// return nil, nil
	default:
		return nil, fmt.Errorf("operador desconocido: %s", op)
	}
}
