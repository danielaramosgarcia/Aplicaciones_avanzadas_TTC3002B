%{
package main
var result int  // Variable global para almacenar el resultado final
%}

%union {
    num int
}

%token <num> NUMBER
%left '+' '-'
%left '*' '/'

%type <num> expr

%%

/* Símbolo inicial */
input:
      expr { result = $1 }   /* Aquí se asigna el resultado final */
    ;

expr:
      expr '+' expr { $$ = $1 + $3 }
    | expr '-' expr { $$ = $1 - $3 }
    | expr '*' expr { $$ = $1 * $3 }
    | expr '/' expr { $$ = $1 / $3 }
    | '(' expr ')'  { $$ = $2 }
    | NUMBER        { $$ = $1 }
    ;
%%
