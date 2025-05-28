package babyduck

import (
	"babyduck/lexer"
	"babyduck/parser"
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
}

// TestIfElseBigQuads verifies the quadruples for if condition.
func TestExampleDrive(t *testing.T) {
	input := `
	program p; 
	var a, b, c, d: int;
	main { 
		if (a + b > d) {
			if (a < b) {
				a = 0;
				b = b + d;
			} else {
				c = a + b;
			};
		} else {
			a = b + c;
		};
		d = b + a * c;
	} end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
}

// TestIfElseBigQuads verifies the quadruples for if condition.
func TestAssign(t *testing.T) {
	input := `
	program p; 
	var a, b, c, d: int;
	main { 
		a = 1 + 2;
		b = a + 3;
		c = a + b;
		d = a + b + c;
	} end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
}

// TestIfElseBigQuads verifies the quadruples for if condition.
func TestPrint(t *testing.T) {
	input := `
	program p; 
	var a, b, c, d: int;
	main { 
		while(a > b * c) do {
			a = a - d;
			print(a);
			**print(b + c);**
		};
		b = c + a;
	} end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
}

func TestPrintExp(t *testing.T) {
	input := `
	program p; 
	var a, b, c, d: int;
	main { 
		while(a > b * c) do {
			a = a - d;
			print(a);
			print(b + c);
		};
		b = c + a;
	} end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
}

func TestPrintSeveral(t *testing.T) {
	input := `
	program p; 
	var a, b, c, d: int;
	main 
	{ 
		while(a > b * c) do {
			a = a - d;
			print(a);
			print(b + c);
			print(a , b, c);
			print((a + b) , d, ((a + b) * c));
		};
		b = c + a;
	} 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
}

// TestFunctionWithLocalVariables verifies parsing of a function with local variables.
func TestFunctionCall(t *testing.T) {
	input := `
	program p; 
	var 
	x: int; 
	y: float;
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
	main { 
		z(3, y);
	} 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Function declaration should parse, got: %v", err)
	}
}

// TestSevFuncCall verifies parsing of a function with local variables.
func TestSevFuncCall(t *testing.T) {
	input := `
	program p; 
	var 
	x: int; 
	y: float;
	void f(a: int, r: float) 
	[ 
		var b, c: float; 
		{ }
	]; 
	void z(s: int,w: float, q: int) 
	[ 
		var f, t: float; 
		{ }
	]; 
	void u(q: int) 
	[ 
		var b, c: int; 
		{ }
	]; 
	main { 
		z(3, y, x);
		f(2, 3.5);
		u(x);
		x = 5 + 2;
	} 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Function declaration should parse, got: %v", err)
	}
}

// TooMuchArguments verifies parsing of a function with local variables.
func TestTooMuchArgumentsFail(t *testing.T) {
	input := `
	program p; 
	var 
	x: int; 
	y: float;
	void f(a: int, r: float) 
	[ 
		var b, c: float; 
		{ }
	]; 
	void z(s: int,w: float, q: int) 
	[ 
		var f, t: float; 
		{ }
	]; 
	void u(q: int) 
	[ 
		var b, c: int; 
		{ }
	]; 
	main { 
		z(3, y, x + 1, x);
		f(2, 3.5);
		u(x);
		x = 5 + 2;
	} 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err == nil {
		t.Errorf("Function deberia fallar, no fallo: %v", err)
	}
}

// TestMissingParams verifies parsing of a function with local variables.
func TestMissingParams(t *testing.T) {
	input := `
	program p; 
	var 
	x: int; 
	y: float;
	void f(a: int, r: float) 
	[ 
		var b, c: float; 
		{ }
	]; 
	void z(s: int,w: float, q: int) 
	[ 
		var f, t: float; 
		{ }
	]; 
	void u(q: int) 
	[ 
		var b, c: int; 
		{ }
	]; 
	main { 
		z(3, y);
		f(2, 3.5);
		u(x);
		x = 5 + 2;
	} 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err == nil {
		t.Errorf("Function deberia fallar por faltar parametros, no fallo: %v", err)
	}
}

// TestEndQuad verifies that the end of a program is parsed correctly.
func TestEndQuad(t *testing.T) {
	input := `
	program p; 
	var a, b, c, d: int;
	main 
	{ 
	} 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
}

// TestGoSubQuad verifies the gosub quadruples for function calls.
func TestGoSubQuad(t *testing.T) {
	input := `
	program p; 
	var 
	x: int; 
	y: float;
	void f(a: int, r: float) 
	[ 
		var b, c: float; 
		{
			x = a + 2;
		 }
	]; 
	void z(s:int) 
	[ 
		var f, t: float; 
		{
			t = s / 2;
		 }
	]; 
	void u(q: int) 
	[ 
		var b, c: int; 
		{ 
			b = q + 1;
		}
	]; 
	main { 
		z(1);
		f(2 + 4, 3.5 + 1);
		u(x);
		x = 5 + 2;
	} 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Function deberia fallar, no fallo: %v", err)
	}
}

// TestEjercicioCompletoDrive checa que los quads sean iuguales al ejercicio completo de la guia de ejercicios.
func TestEjercicioCompletoDrive(t *testing.T) {
	input := `
	program patito; 

	var 
		i, j, k: int; 
		f: float;

	void uno(a: int, b: int) 
	[ 
		{
			if (a > 0) 
			{
				i= a + b * j + i; 
				print ( i + j );
				uno ( a - i, i);
			} 
			else 
			{
				print ( a + b );
			};
		}
	]; 

	void dos(a:int, g:float) 
	[ 
		var i: int; 
		{
			i=a;
			while (a>0) do
			{ 
				a = a - k * j;
				uno( a * 2, a + k );
				g = g * j - k;
			};
		 }
	]; 

	main { 
		i= 2; 
		k= i + 1;
		f= 3.14;
		while (i > 0) do
		{ 
			print( (i + k / f * 3) + 3 ); 
			print(i , j * 2, f * 2 + 1.5); 
			i = i - k ;
		};
	} 

	end
	`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("ERR : %v", err)
	}
}

func TestPrintCte(t *testing.T) {
	input := `
	program p; 
	var a, b, c, d: int;
	main 
	{ 
		print('hola');
	} 
	end`
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	if _, err := p.Parse(l); err != nil {
		t.Errorf("Empty program should parse without error, got: %v", err)
	}
}
