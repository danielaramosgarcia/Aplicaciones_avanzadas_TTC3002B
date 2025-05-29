package VM

import (
	"babyduck/data_structures"
	"babyduck/lexer"
	"babyduck/parser"
	"fmt"
	"os"
)

func uploadInfo(ctx *data_structures.Context) *Machine {
	return &Machine{
		Quads:        ctx.Quads.Quads,
		IP:           0,
		memoryStack:  []*FuncInstance{},
		globalMemory: make(map[int]interface{}),
		Name:         ctx.programName
	}
}

func codeInput(code string) string {

	// 2) Tokeniza y parsea
	l := lexer.NewLexer([]byte(code))
	p := parser.NewParser()
	result, err := p.Parse(l)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error de parseo: %v\n", err)
		os.Exit(1)
	}
	ctx := result.(*data_structures.Context)
	machine := uploadInfo(ctx)

	output, err := machine.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error en VM: %v\n", err)
		os.Exit(1)
	}

	return output
}
