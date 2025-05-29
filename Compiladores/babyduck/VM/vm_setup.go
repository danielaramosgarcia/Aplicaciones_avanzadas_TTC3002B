package VM

import (
	"babyduck/data_structures"
	"babyduck/lexer"
	"babyduck/parser"
	"fmt"
	"os"
	"sort"
)

// func ReturnVarTable(vt *data_structures.VarTable) map[int]interface{} {
// 	memory := make(map[int]interface{})
// 	for _, v := range vt.Vars {
// 		memory[v.DirInt] = 0
// 	}
// 	return memory
// }

func NewMachine(ctx *data_structures.Context) *Machine {
	// --- memoria global igual que antes ---
	global := make(map[int]interface{})
	for _, ve := range ctx.FuncDir.Funcs[ctx.ProgramName].VarTable.Vars {
		global[ve.DirInt] = 0
	}
	for dir, v := range ctx.ConstTable.Num {
		global[dir] = v
	}
	for dir, v := range ctx.ConstTable.Float {
		global[dir] = v
	}
	for dir, v := range ctx.ConstTable.Str {
		global[dir] = v
	}
	for dir, v := range ctx.ConstTable.Bool {
		global[dir] = v
	}

	// --- construyo ParamAddrs ---
	paramAddrs := make(map[string][]int)
	for name, fe := range ctx.FuncDir.Funcs {
		n := len(fe.ParamTypes)
		if n == 0 {
			paramAddrs[name] = nil
			continue
		}
		// reuno todas las direcciones y las ordeno
		var dirs []int
		for dir := range fe.VarTable.Vars {
			dirs = append(dirs, dir)
		}
		sort.Ints(dirs)
		// los primeros n son los parámetros (porque se insertaron primero)
		paramAddrs[name] = dirs[:n]
	}

	return &Machine{
		Ctx:        ctx,
		memStack:   []map[int]interface{}{global},
		retStack:   []int{},
		Quads:      ctx.Quads.Quads,
		IP:         0,
		ParamAddrs: paramAddrs,
	}
}

func (m *Machine) currentMem() map[int]interface{} {
	return m.memStack[len(m.memStack)-1]
}

// // NewMachine crea el marco global y carga constantes + vars globales.
// func NewMachine(ctx *data_structures.Context) *Machine {
// 	global := make(map[int]interface{})
// 	// 1) Variables globales
// 	for _, ve := range ctx.FuncDir.Funcs[ctx.ProgramName].VarTable.Vars {
// 		global[ve.DirInt] = 0
// 	}
// 	// 2) Constantes
// 	for dir, v := range ctx.ConstTable.Num {
// 		global[dir] = v
// 	}
// 	for dir, v := range ctx.ConstTable.Float {
// 		global[dir] = v
// 	}
// 	for dir, v := range ctx.ConstTable.Str {
// 		global[dir] = v
// 	}
// 	for dir, v := range ctx.ConstTable.Bool {
// 		global[dir] = v
// 	}
// 	return &Machine{
// 		Ctx:         ctx,
// 		memStack:    []map[int]interface{}{global},
// 		retStack:    []int{},
// 		Quads:       ctx.Quads.Quads,
// 		IP:          0,
// 		PrintOutput: strings.Builder{},
// 	}
// }

// readMem busca addr del tope de memStack hacia abajo.
func (m *Machine) readMem(addr int) (interface{}, bool) {
	// println("LLEGO A READ MEM Con addr", addr)
	// for i, mem := range m.memStack {
	// 	fmt.Printf("memStack[%d]: %v\n", i, mem)
	// }
	for i := len(m.memStack) - 1; i >= 0; i-- {
		if v, ok := m.memStack[i][addr]; ok {
			return v, true
		}
	}
	return nil, false
}

// writeMem actualiza el primer marco que contenga addr,
// o global si solo existe allí.
func (m *Machine) writeMem(addr int, val interface{}) {
	for i := len(m.memStack) - 1; i >= 0; i-- {
		if _, ok := m.memStack[i][addr]; ok {
			m.memStack[i][addr] = val
			return
		}
	}
	// si no existía, va a global
	m.memStack[0][addr] = val
}

// lookupFuncEntry por índice de función
func (m *Machine) lookupFuncEntry(idx int) (*data_structures.FuncEntry, error) {
	name, ok := m.Ctx.FuncIndex[idx]
	if !ok {
		return nil, fmt.Errorf("índice función %d no registrado", idx)
	}
	fe, ok := m.Ctx.FuncDir.Funcs[name]
	if !ok {
		return nil, fmt.Errorf("función %q no encontrada", name)
	}
	return fe, nil
}

func CodeInput(code string) string {
	// 2) Tokeniza y parsea
	l := lexer.NewLexer([]byte(code))
	p := parser.NewParser()
	result, err := p.Parse(l)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error de parseo: %v\n", err)
		os.Exit(1)
	}
	ctx := result.(*data_structures.Context)
	machine := NewMachine(ctx)
	output, err := machine.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error en VM: %v\n", err)
		os.Exit(1)
	}
	// machine = NewMachine(ctx) // Reinicializa la máquina para evitar problemas de estado
	return output
}

// func main() {
// 	// Ejemplo de uso
// 	code := `
// 	program p;
// 	var x: float;
// 	void f(a: int)
// 	[
// 		var b, c: int;
// 		{
// 			x = (a + 3) / c - 5;
// 		}
// 	];
// 	main { }
// 	end`
// 	output := CodeInput(code)
// 	fmt.Println("Output:", output)
// }
