package VM

import (
	"babyduck/data_structures"
	"fmt"
	"strings"
)

func NewMachine(quads []data_structures.Quadruple) *Machine {
	return &Machine{
		Memory: make(map[int]interface{}),
		Quads:  quads,
		IP:     0,
	}
}

func (m *Machine) Run() (string, error) {
	var output strings.Builder

	for m.IP < len(m.Quads) {
		q := m.Quads[m.IP]
		switch q.Op {
		case ADD:
			left := m.Memory[q.Arg1].(int)
			right := m.Memory[q.Arg2].(int)
			m.Memory[q.Result] = left + right
			m.IP++

		case SUB:
			left := m.Memory[q.Arg1].(int)
			right := m.Memory[q.Arg2].(int)
			m.Memory[q.Result] = left - right
			m.IP++

		case MUL:
			left := m.Memory[q.Arg1].(int)
			right := m.Memory[q.Arg2].(int)
			m.Memory[q.Result] = left * right
			m.IP++

		case DIV:
			left := m.Memory[q.Arg1].(int)
			right := m.Memory[q.Arg2].(int)
			m.Memory[q.Result] = left / right
			m.IP++

		case LT:
			left := m.Memory[q.Arg1].(int)
			right := m.Memory[q.Arg2].(int)
			if left < right {
				m.Memory[q.Result] = 1
			} else {
				m.Memory[q.Result] = 0
			}
			m.IP++

		case GT:
			left := m.Memory[q.Arg1].(int)
			right := m.Memory[q.Arg2].(int)
			if left > right {
				m.Memory[q.Result] = 1
			} else {
				m.Memory[q.Result] = 0
			}
			m.IP++

		case NEQ:
			left := m.Memory[q.Arg1].(int)
			right := m.Memory[q.Arg2].(int)
			if left != right {
				m.Memory[q.Result] = 1
			} else {
				m.Memory[q.Result] = 0
			}
			m.IP++

		case GOTOFALSE:
			cond := m.Memory[q.Arg1].(int)
			if cond == 0 {
				m.IP = q.Result
			} else {
				m.IP++
			}

		case GOTO:
			m.IP = q.Result

		case PRINT:
			// Arg1 es la dirección del valor a imprimir
			if val, ok := m.Memory[q.Arg1]; ok {
				// Convierte a string la línea de salida (añade salto)
				output.WriteString(fmt.Sprintf("%v\n", val))
			} else {
				return "", fmt.Errorf("PRINT: memoria vacía en dirección %d", q.Arg1)
			}
			m.IP++

		case ASSIGN:
			m.Memory[q.Arg1] = m.Memory[q.Result]
			m.IP++

		case END:
			// Terminamos la ejecución; devolvemos todo lo impreso
			return output.String(), nil

		default:
			return "", fmt.Errorf("código de operación desconocido %d", q.Op)
		}
	}

	// En caso de salirse del bucle sin END explícito
	return output.String(), nil
}
