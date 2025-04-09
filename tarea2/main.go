package main

import "fmt"

func main() {
	input := "3 + 4 * (2 + 1)"
	lexer := &Lexer{}
	lexer.Init(input)

	if yyParse(lexer) == 0 {
		fmt.Printf("Resultado: %d\n", result) // Usamos la variable global del parser
	} else {
		fmt.Println("Error al analizar la expresión.")
	}
}
