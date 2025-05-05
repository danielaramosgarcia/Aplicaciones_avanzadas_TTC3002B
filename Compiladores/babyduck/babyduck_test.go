package babyduck

import (
	"babyduck/lexer"
	"babyduck/parser"
	"testing"
)

// testData maps input programs to a boolean indicating whether parsing should succeed.
var testData = map[string]bool{
	// Test case 1: testing empty program
	"program p; main { } end": true,
	// Test case 2: testing variable declaration and assignment
	"program p; var x: int; main { x = 5; } end": true,
	// Test case 3: testing float
	"** sample ** program p; var x: float; main { x = 3.14; } end": true,
	// Test case 4: testing if
	"program p; var x: int; main { if (x < 10) { x = x + 1; }; } end": true,
	// Test case 5: testing comments
	"** Testing comments ** program p; var x: int; main { x = 5; } end": true,
	// Test case 6: testing print
	"program p; var x: int; main { x = 5; print(x); } end": true,
	// Test case 7: testing whitespaces
	`** Testing new line whitespace ** program p; 
	main { } 
	end`: true,
	// Test case 8: testing while
	`program p; var x: int; main 
	{ while (x < 10) do { x = x + 1; }; } 
	end`: true,
	// Test case 9: missing semicolon after identifier
	"program p main { } end": false,
	// Test case 10: missing 'end'
	"program p; main { }": false,
	// Test case 11: missing semicolon after assignment
	"program p; var x: int; main { x = 5 } end": false,
	// Test case 12: extra text after 'end'
	"program p; var x: int; main { x = 5; } end extra": false,
	// Test case 13: testing function declaration
	"program p; var x: int; void f(a: int) [{ b = a + 2; }]; main { } end": true,
}

func TestParse(t *testing.T) {
	i := 1
	for input, ok := range testData {
		// Log the test input and output
		t.Logf("=== Parsing Test #%d", i)
		l := lexer.NewLexer([]byte(input))
		p := parser.NewParser()
		_, err := p.Parse(l)

		// Check expectation
		if (err == nil) != ok {
			if ok {
				t.Errorf("unexpected error parsing valid input:\n%s\nerror: %v", input, err)
				t.Logf("Parse error: %v", err)
			} else {
				t.Errorf("expected error parsing invalid input, but got none:\n%s", input)
				t.Logf("Parse error: %v", err)
			}
		}
		i++
	}
}
