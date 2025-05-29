package VM

import (
	"fmt"
)

// Run aplica todos los cuadruplos y devuelve todo lo impreso.
func (m *Machine) Run() (string, error) {
	callStack := []string{m.Ctx.ProgramName}

	for m.IP < len(m.Quads) {
		q := m.Quads[m.IP]
		switch q.Op {
		// 1) Aritmética
		case ADD, SUB, MUL, DIV:
			lv, _ := m.readMem(q.Arg1)
			rv, _ := m.readMem(q.Arg2)
			l, r := lv.(int), rv.(int)
			var out int
			switch q.Op {
			case ADD:
				out = l + r
			case SUB:
				out = l - r
			case MUL:
				out = l * r
			case DIV:
				out = l / r
			}
			m.writeMem(q.Result, out)
			m.IP++

		// 2) Comparaciones
		case LT, GT, NEQ:
			lv, _ := m.readMem(q.Arg1)
			rv, _ := m.readMem(q.Arg2)
			l, r := lv.(int), rv.(int)
			var b int
			switch q.Op {
			case LT:
				if l < r {
					b = 1
				}
			case GT:
				if l > r {
					b = 1
				}
			case NEQ:
				if l != r {
					b = 1
				}
			}
			m.writeMem(q.Result, b)
			m.IP++

		// 3) Saltos
		case GOTOFALSE:
			cv, _ := m.readMem(q.Arg1)
			if cv.(int) == 0 {
				m.IP = q.Result
			} else {
				m.IP++
			}
		case GOTO:
			m.IP = q.Result

			// 4) Call stack
		case ERA:
			fe, err := m.lookupFuncEntry(q.Arg1)
			if err != nil {
				return "", err
			}
			// crear marco local
			local := make(map[int]interface{})
			for _, ve := range fe.VarTable.Vars {
				local[ve.DirInt] = 0
			}
			m.memStack = append(m.memStack, local)
			// ¡empujamos el nombre de la función!
			callStack = append(callStack, fe.Name)
			m.IP++

		case PARAM:
			// q.Arg1 = dir del argumento, q.Result = índice 0-based del parámetro
			val, ok := m.readMem(q.Arg1)
			if !ok {
				return "", fmt.Errorf("PARAM: valor no inicializado en %d", q.Arg1)
			}
			// función activa = cima de callStack
			fnName := callStack[len(callStack)-1]
			addrs := m.ParamAddrs[fnName]
			idx := q.Result - 1
			if idx < 0 || idx >= len(addrs) {
				return "", fmt.Errorf("PARAM: índice %d fuera de rango para %q", idx, fnName)
			}
			targetDir := addrs[idx]
			// asignamos en el marco actual
			m.memStack[len(m.memStack)-1][targetDir] = val
			m.IP++

		case GOSUB:
			m.retStack = append(m.retStack, m.IP+1)
			fe, err := m.lookupFuncEntry(q.Arg1)
			if err != nil {
				return "", err
			}
			m.IP = fe.CuadStart

		case ENDF:
			// restaurar retorno y marco
			ret := m.retStack[len(m.retStack)-1]
			m.retStack = m.retStack[:len(m.retStack)-1]
			m.memStack = m.memStack[:len(m.memStack)-1]
			// desapilar el nombre de la función
			if len(callStack) > 1 {
				callStack = callStack[:len(callStack)-1]
			}
			m.IP = ret

		// 5) Asignación
		case EQ:
			// Lee el valor fuente en Arg1 y lo escribe en la dirección Result
			val, ok := m.readMem(q.Arg1)
			if !ok {
				return "", fmt.Errorf("EQ: dirección fuente %d no inicializada", q.Arg1)
			}
			m.writeMem(q.Result, val)
			m.IP++

		// 6) Print
		case PRINT:
			val, _ := m.readMem(q.Result)
			m.PrintOutput.WriteString(fmt.Sprintf("%v\n", val))
			m.IP++

		// 7) Fin de programa
		case END:

			return m.PrintOutput.String(), nil

		default:
			return "", fmt.Errorf("operación desconocida %d", q.Op)
		}
	}
	// si se acaba sin END explícito
	return m.PrintOutput.String(), nil
}
