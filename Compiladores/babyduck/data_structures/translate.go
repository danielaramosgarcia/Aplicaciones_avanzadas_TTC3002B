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
*/

// TanslateOp recibe un operador como string y devuelve su representación numérica
func TranslateOp(op string) (int, error) {
	switch op {
	case "+":
		return 10, nil
	case "-":
		return 20, nil
	case "*":
		return 30, nil
	case "/":
		return 40, nil
	case "<":
		return 50, nil
	case ">":
		return 60, nil
	case "!=":
		return 70, nil
	default:
		return -1, fmt.Errorf("operador desconocido: %s", op)
	}
}
