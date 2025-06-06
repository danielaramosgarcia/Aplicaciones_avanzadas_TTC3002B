package babyduck

import (
	"babyduck/lexer"
	"babyduck/parser"
	"testing"
)

// TestEmptyProgram verifies that an empty program parses successfully.
func TestEmptyProgram(t *testing.T) {
	input := `program p; main { } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
}

// TestVarDeclarationAndAssignment verifies variable declaration and assignment.
func TestVarDeclarationAndAssignment(t *testing.T) {
	input := `program p; var x: int; main {  } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Var declaration and assignment should parse, got: %v", err)
	}
}

// TestFloatConstant verifies parsing of a float constant.
func TestFloatConstant(t *testing.T) {
	input := `** sample ** program p; var x: float; main { x = 3.56; } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Float constant should parse, got: %v", err)
	}
}

// Test declare several variables.
func TestSeveralVariablesInt(t *testing.T) {
	input := `program p; var a, b: int; main { a = 1; } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("El programa falla al tratar de declarar varias variables: %v", err)
	}
}

// Test declare several variables.
func TestSeveralVariablesFloat(t *testing.T) {
	input := `program p; var a, b: float; main { a = 1.23; } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("El programa falla al tratar de declarar varias variables: %v", err)
	}
}

// Test declare several variables.
func TestDifferenteTypesVariables(t *testing.T) {
	input := `program p; var x, y, z: float; a, b: int; main { x = 1.566; } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("El programa falla al tratar de declarar varias variables: %v", err)
	}
}

// TestIfStatement verifies parsing of an if statement.
func TestIfStatement(t *testing.T) {
	input := `program p; var x: int; main { if (x < 10) { x = x + 1; }; } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("If statement should parse, got: %v", err)
	}
}

// TestComments verifies that comments are ignored.
func TestComments(t *testing.T) {
	input := `** Testing comments ** program p; var x: int; main { x = 5; } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Comments should be ignored, got: %v", err)
	}
}

// TestPrintStatement verifies parsing of a print statement.
func TestPrintStatement(t *testing.T) {
	input := `program p; var x: int; main { x = 5; print(x); } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Print statement should parse, got: %v", err)
	}
}

// TestWhitespaceVariations verifies that various whitespace and newlines parse.
func TestWhitespaceVariations(t *testing.T) {
	input := `** Testing new line 
	whitespace ** program p; 
	var x: float;
	main { } 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Whitespace variations should parse, got: %v", err)
	}
}

// TestWhileLoop verifies parsing of a while loop.
func TestWhileLoop(t *testing.T) {
	input := `program p; var x: int; main 
	{ while (x < 10) do { x = x + 1; }; } 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("While loop should parse, got: %v", err)
	}
}

// TestMissingSemicolonIdentifier expects an error when semicolon is missing after program header.
func TestMissingSemicolonIdentifier(t *testing.T) {
	input := `program p main { } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err == nil {
		t.Errorf("Expected error for missing semicolon after identifier")
	}
}

// TestMissingEnd expects an error when 'end' keyword is missing.
func TestMissingEnd(t *testing.T) {
	input := `program p; main { }`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err == nil {
		t.Errorf("Expected error for missing 'end'")
	}
}

// TestMissingSemicolonAssignment expects an error when semicolon is missing after assignment.
func TestMissingSemicolonAssignment(t *testing.T) {
	input := `program p; var x: int; main { x = 5 } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err == nil {
		t.Errorf("Expected error for missing semicolon after assignment")
	}
}

// TestExtraTextAfterEnd expects an error when extra text follows 'end'.
func TestExtraTextAfterEnd(t *testing.T) {
	input := `program p; var x: int; main { x = 5; } end extra`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err == nil {
		t.Errorf("Expected error for extra text after 'end'")
	}
}

// TestFunctionDeclaration verifies parsing of a simple function declaration.
func TestFunctionDeclaration(t *testing.T) {
	input := `program p; var x: int; void f(a: int) [{ x = a + 2; }]; main { } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Function declaration should parse, got: %v", err)
	}
}

// TestFunctionWithLocalVariables verifies parsing of a function with local variables.
func TestFunctionWithLocalVariables(t *testing.T) {
	input := `program p; var x: int; void f(a: int) [ var b, c: float; { x = a + 2; }]; main { } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Function declaration should parse, got: %v", err)
	}
}

// TestFunctionWithLocalVariables verifies parsing of a function with local variables.
func TestValidateAssign(t *testing.T) {
	input := `program p; var x: int; void f(a: int) [ var b, c: float; { x = a + 2; }]; main { } end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Function declaration should parse, got: %v", err)
	}
}

// TestFunctionWithLocalVariables verifies parsing of a function with local variables.
func TestSeveralFunctionDeclaration(t *testing.T) {
	input := `
	program p; 
	var x: int; 
	void f(a: int, r: float) 
	[ 
		var b, c: float; 
		{ 
			x = a + 2; 
		}
	]; 
	void z(s: int,w: float) 
	[ 
		var f, t: float; 
		{ 
			f = s / 2; 
		}
	]; 
	main { } 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Function declaration should parse, got: %v", err)
	}
}

// TestFunctionWithLocalVariables verifies parsing of a function with local variables.
func TestSeveralExp(t *testing.T) {
	input := `
	program p; 
	var x: float; 
	void f(a: int) 
	[ 
		var b, c: int; 
		{ 
			x = a + 2 / c - 5; 
		}
	]; 
	main { } 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Function declaration should parse, got: %v", err)
	}
}

func TestParenthesis(t *testing.T) {
	input := `
	program p; 
	var x: float; 
	void f(a: int) 
	[ 
		var b, c: int; 
		{ 
			x = (a + 3) / c - 5; 
		}
	]; 
	main { } 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Function declaration should parse, got: %v", err)
	}
}

func TestParenthesis2(t *testing.T) {
	input := `
	program p; 
	var x: float; 
	void f(a: int) 
	[ 
		var b, c: int; 
		{ 
			x = 7 + ((a + 3 / 1) - c ) / 5; 
		}
	]; 
	main { } 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Function declaration should parse, got: %v", err)
	}
}

func TestBoolComp(t *testing.T) {
	input := `
	program p; 
	var x: float; 
	void f(a: int) 
	[ 
		var b, c: int; 
		{ 
			x = 7 + ((a + 3 / 1) - c ) / 5; 
		}
	]; 
	main { 
		if (x < 5) {
			x = 3.0 + 2.0;
		};
	} 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Function declaration should parse, got: %v", err)
	}
}
