package main

import (
	"babyduck/VM"
	"fmt"
)

func main() {
	output := ""
	input := `
	program primerintento; 
	
	var x: int;
	
	void uno(a: int) 
	[ 
		{
			while (a < 10) do
			{
				print('El valor de a es: ', a);
				a = a + 1;
			};
		}
	]; 

	void dos() 
	[ 
		var y: int;
		{
			y = 10;
			print(y);
			if (x > 5) 
			{
				print ( 'X es mayor que 5' );
			}
			else
			{
				print ( 'X es menor a 5' );
			};
		}
	]; 
	
	main { 
		x = 3;
		print('Hola mundo!');
		uno(x);
		dos();
		print(21/3);
		print('Fin del programa');
	} end`
	output = VM.CodeInput(input)
	fmt.Print("=================== OUTPUT DE LA VM ==========================\n")
	fmt.Print(output)
}
