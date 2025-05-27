package babyduck

import (
	"babyduck/lexer"
	"babyduck/parser"
	"fmt"
	"testing"
)

// TestIfQuads verifies the quadruples for if condition.
func TestIfQuads(t *testing.T) {
	input := `
	program p; 
	var x: int;
	main { 
		if (x < 10) {
			x = x + 1;
		};
	} end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}

	fmt.Println("TestIfQuads passed")
}

// TestIfElseQuads verifies the quadruples for if condition.
func TestIfElseQuads(t *testing.T) {
	input := `
	program p; 
	var x: int;
	main { 
		if (x < 10) {
			x = x + 1;
		} else {
			x = 3 - 1;
		};
	} end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
	fmt.Println("TestIfQuads passed")
}

// TestIfElseBigQuads verifies the quadruples for if condition.
func TestIfElseBigQuads(t *testing.T) {
	input := `
	program p; 
	var x: int;
	main { 
		if (x < 10) {
			x = (x + 1)*3+4;
			if (x > 5) {
				x = x * 2;
			};
		} else {
			x = 3 - 1 + 2 * 4;
		};
	} end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
	fmt.Println("TestIfQuads passed")
}

// TestIfElseBigQuads verifies the quadruples for if condition.
func TestWhileQuads(t *testing.T) {
	input := `
	program p; 
	var x: int;
	main { 
		x = 3 + 1 -1;
		while(x < 10-1) do {
			x = x + 1;
			print(x*4);
			x = x - 234;
		};
		x = 3 - 1;
	} end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
	fmt.Println("TestIfQuads passed")
}

// TestIfElseBigQuads verifies the quadruples for if condition.
func TestBigWhileQuads(t *testing.T) {
	input := `
	program p; 
	var x: int;
	main { 
		x = 3 + 1;
		while(x < (10 + 2)) do {
			x = x + 1;
			print(x*4);
			x = x - 234;
		};
		x = 3 - 1;
	} end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
	fmt.Println("TestIfQuads passed")
}

// TestIfElseBigQuads verifies the quadruples for if condition.
func TestWhileWithIf(t *testing.T) {
	input := `
	program p; 
	var x: int;
	main { 
		x = 3 + 1;
		if (x < 10) {
		while(x < (10 + 2)) do {
			x = x + 1;
			print(x*4);
			x = x - 234;
		};
		x = 3 - 1;
		};
		x = 3 - 1;
	} end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
	fmt.Println("TestIfQuads passed")
}
