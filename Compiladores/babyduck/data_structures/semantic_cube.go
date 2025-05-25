package data_structures

import (
	"fmt"
)

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
func TanslateOp(op string) (int, error) {
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
binaryCube define el tipo resultado para operadores binarios
binaryCube[op][leftType][rightType] = resultType
*/
var binaryCube = map[int]map[int]map[int]int{
	10: {
		0: {0: 0, 1: 1},
		1: {0: 1, 1: 1},
	},
	20: {
		0: {0: 0, 1: 1},
		1: {0: 1, 1: 1},
	},
	30: {
		0: {0: 0, 1: 1},
		1: {0: 1, 1: 1},
	},
	40: {
		0: {0: 1, 1: 1},
		1: {0: 1, 1: 1},
	},
	50: {
		0: {0: 2, 1: 2},
		1: {0: 2, 1: 2},
	},
	60: {
		10: {10: 2, 1: 2},
		1:  {10: 2, 1: 2},
	},
	70: {
		0: {0: 2, 1: 2},
		1: {0: 2, 1: 2},
	},
}

/*
unaryCube define el tipo resultado para operadores unarios
unaryCube[op][operandType] = resultType
*/
var unaryCube = map[int]map[int]int{
	0: {0: 0, 1: 1},
	1: {0: 0, 1: 1},
}

// ResultBinary devuelve el tipo resultante de una operación binaria o un error si no es válida
func ResultBinary(op int, left, right int) (int, error) {
	opMap, ok := binaryCube[op]
	if !ok {
		return 0, fmt.Errorf("operador binario desconocido: %d", op)
	}
	rightMap, ok := opMap[left]
	if !ok {
		return 0, fmt.Errorf("operador %d no soportado para tipo izquierdo %v", op, left)
	}
	resType, ok := rightMap[right]
	if !ok {
		return 0, fmt.Errorf("operador %d no soportado para tipos %v, %v", op, left, right)
	}
	return resType, nil
}

// ResultUnary devuelve el tipo resultante de una operación unaria o un error si no es válida
func ResultUnary(op int, operand int) (int, error) {
	opMap, ok := unaryCube[op]
	if !ok {
		return 0, fmt.Errorf("operador unario desconocido: %d", op)
	}
	resType, ok := opMap[operand]
	if !ok {
		return 0, fmt.Errorf("operador %d no soportado para tipo %v", op, operand)
	}
	return resType, nil
}
