package data_structures

import (
	"fmt"
)

// binaryCube define el tipo resultado para operadores binarios
// binaryCube[op][leftType][rightType] = resultType
var binaryCube = map[string]map[Tipo]map[Tipo]Tipo{
	"+": {
		Int:   {Int: Int, Float: Float},
		Float: {Int: Float, Float: Float},
	},
	"-": {
		Int:   {Int: Int, Float: Float},
		Float: {Int: Float, Float: Float},
	},
	"*": {
		Int:   {Int: Int, Float: Float},
		Float: {Int: Float, Float: Float},
	},
	"/": {
		Int:   {Int: Float, Float: Float},
		Float: {Int: Float, Float: Float},
	},
	"<": {
		Int:   {Int: Bool, Float: Bool},
		Float: {Int: Bool, Float: Bool},
	},
	">": {
		Int:   {Int: Bool, Float: Bool},
		Float: {Int: Bool, Float: Bool},
	},
	"!=": {
		Int:   {Int: Bool, Float: Bool},
		Float: {Int: Bool, Float: Bool},
	},
}

// unaryCube define el tipo resultado para operadores unarios
// unaryCube[op][operandType] = resultType
var unaryCube = map[string]map[Tipo]Tipo{
	"+": {Int: Int, Float: Float},
	"-": {Int: Int, Float: Float},
	"!": {Bool: Bool},
}

// ResultBinary devuelve el tipo resultante de una operación binaria o un error si no es válida
func ResultBinary(op string, left, right Tipo) (Tipo, error) {
	opMap, ok := binaryCube[op]
	if !ok {
		return 0, fmt.Errorf("operador binario desconocido: %s", op)
	}
	rightMap, ok := opMap[left]
	if !ok {
		return 0, fmt.Errorf("operador %s no soportado para tipo izquierdo %v", op, left)
	}
	resType, ok := rightMap[right]
	if !ok {
		return 0, fmt.Errorf("operador %s no soportado para tipos %v, %v", op, left, right)
	}
	return resType, nil
}

// ResultUnary devuelve el tipo resultante de una operación unaria o un error si no es válida
func ResultUnary(op string, operand Tipo) (Tipo, error) {
	opMap, ok := unaryCube[op]
	if !ok {
		return 0, fmt.Errorf("operador unario desconocido: %s", op)
	}
	resType, ok := opMap[operand]
	if !ok {
		return 0, fmt.Errorf("operador %s no soportado para tipo %v", op, operand)
	}
	return resType, nil
}
