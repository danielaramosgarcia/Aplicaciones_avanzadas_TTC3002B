package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

type Lexer struct {
	s       scanner.Scanner
	lastNum int
	Result  int // Aquí guardamos el resultado final
}

func (l *Lexer) Init(input string) {
	l.s.Init(strings.NewReader(input))
	l.s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanChars | scanner.ScanStrings | scanner.ScanRawStrings
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok := l.s.Scan()
	switch tok {
	case scanner.EOF:
		return 0
	case '+', '-', '*', '/', '(', ')':
		return int(tok)
	case scanner.Int:
		var val int
		fmt.Sscanf(l.s.TokenText(), "%d", &val)
		lval.num = val
		return NUMBER
	default:
		return int(tok)
	}
}

func (l *Lexer) Error(s string) {
	fmt.Printf("Error léxico: %s\n", s)
}
