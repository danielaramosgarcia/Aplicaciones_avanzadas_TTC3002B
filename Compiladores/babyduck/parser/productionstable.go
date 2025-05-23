// Code generated by gocc; DO NOT EDIT.

package parser

import (
    "babyduck/token"
    "babyduck/data_structures"
)
// Creamos un contexto único para todo el programa:
var ctx = data_structures.NewContext()

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib, interface{}) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Start	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Start : Reset Programa	<< ctx.ReturnContext() >>`,
		Id:         "Start",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ctx.ReturnContext()
		},
	},
	ProdTabEntry{
		String: `Reset : empty	<< ctx.Reset() >>`,
		Id:         "Reset",
		NTType:     2,
		Index:      2,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ctx.Reset()
		},
	},
	ProdTabEntry{
		String: `Programa : "program" id ";" Vars Funcs "main" Body "end"	<<  >>`,
		Id:         "Programa",
		NTType:     3,
		Index:      3,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Vars : empty	<<  >>`,
		Id:         "Vars",
		NTType:     4,
		Index:      4,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Vars : "var" VarList ":" Type ";" Vars	<< ctx.RegisterGlobalVars(X[1].([]string), X[3].(data_structures.Tipo)) >>`,
		Id:         "Vars",
		NTType:     4,
		Index:      5,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ctx.RegisterGlobalVars(X[1].([]string), X[3].(data_structures.Tipo))
		},
	},
	ProdTabEntry{
		String: `VarList : id	<< data_structures.MakeVarList(string(X[0].(*token.Token).Lit)) >>`,
		Id:         "VarList",
		NTType:     5,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return data_structures.MakeVarList(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `VarList : id "," VarList	<< data_structures.ConcatVarList(string(X[0].(*token.Token).Lit), X[2].([]string)) >>`,
		Id:         "VarList",
		NTType:     5,
		Index:      7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return data_structures.ConcatVarList(string(X[0].(*token.Token).Lit), X[2].([]string))
		},
	},
	ProdTabEntry{
		String: `Type : "int"	<< data_structures.Int, nil >>`,
		Id:         "Type",
		NTType:     6,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return data_structures.Int, nil
		},
	},
	ProdTabEntry{
		String: `Type : "float"	<< data_structures.Float, nil >>`,
		Id:         "Type",
		NTType:     6,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return data_structures.Float, nil
		},
	},
	ProdTabEntry{
		String: `Funcs : empty	<<  >>`,
		Id:         "Funcs",
		NTType:     7,
		Index:      10,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Funcs : Func Funcs	<<  >>`,
		Id:         "Funcs",
		NTType:     7,
		Index:      11,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Func : FuncStart Vars Body "]" ";"	<< ctx.ExitFunction() >>`,
		Id:         "Func",
		NTType:     8,
		Index:      12,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ctx.ExitFunction()
		},
	},
	ProdTabEntry{
		String: `FuncStart : "void" id "(" ")" "["	<< ctx.RegisterAndEnterFunction(string(X[1].(*token.Token).Lit), data_structures.Void, []data_structures.Param{}) >>`,
		Id:         "FuncStart",
		NTType:     9,
		Index:      13,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ctx.RegisterAndEnterFunction(string(X[1].(*token.Token).Lit), data_structures.Void, []data_structures.Param{})
		},
	},
	ProdTabEntry{
		String: `FuncStart : "void" id "(" ParamList ")" "["	<< ctx.RegisterAndEnterFunction(string(X[1].(*token.Token).Lit), data_structures.Void, X[3].([]data_structures.Param)) >>`,
		Id:         "FuncStart",
		NTType:     9,
		Index:      14,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ctx.RegisterAndEnterFunction(string(X[1].(*token.Token).Lit), data_structures.Void, X[3].([]data_structures.Param))
		},
	},
	ProdTabEntry{
		String: `ParamList : id ":" Type	<< data_structures.MakeParam(string(X[0].(*token.Token).Lit),X[2].(data_structures.Tipo)) >>`,
		Id:         "ParamList",
		NTType:     10,
		Index:      15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return data_structures.MakeParam(string(X[0].(*token.Token).Lit),X[2].(data_structures.Tipo))
		},
	},
	ProdTabEntry{
		String: `ParamList : id ":" Type "," ParamList	<< data_structures.PrependParam(string(X[0].(*token.Token).Lit), X[2].(data_structures.Tipo), X[4].([]data_structures.Param)) >>`,
		Id:         "ParamList",
		NTType:     10,
		Index:      16,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return data_structures.PrependParam(string(X[0].(*token.Token).Lit), X[2].(data_structures.Tipo), X[4].([]data_structures.Param))
		},
	},
	ProdTabEntry{
		String: `Body : "{" StatementList "}"	<<  >>`,
		Id:         "Body",
		NTType:     11,
		Index:      17,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StatementList : empty	<<  >>`,
		Id:         "StatementList",
		NTType:     12,
		Index:      18,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `StatementList : Statement StatementList	<<  >>`,
		Id:         "StatementList",
		NTType:     12,
		Index:      19,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Assign	<<  >>`,
		Id:         "Statement",
		NTType:     13,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Condition	<<  >>`,
		Id:         "Statement",
		NTType:     13,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Cycle	<<  >>`,
		Id:         "Statement",
		NTType:     13,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : F_Call	<<  >>`,
		Id:         "Statement",
		NTType:     13,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : Print	<<  >>`,
		Id:         "Statement",
		NTType:     13,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Assign : id "=" Expression ";"	<< ctx.ValidateAssign(string(X[0].(*token.Token).Lit), X[2].(data_structures.Tipo)) >>`,
		Id:         "Assign",
		NTType:     14,
		Index:      25,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ctx.ValidateAssign(string(X[0].(*token.Token).Lit), X[2].(data_structures.Tipo))
		},
	},
	ProdTabEntry{
		String: `Expression : AddExpr RelExpr	<<  >>`,
		Id:         "Expression",
		NTType:     15,
		Index:      26,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelExpr : empty	<<  >>`,
		Id:         "RelExpr",
		NTType:     16,
		Index:      27,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `RelExpr : RelOp AddExpr	<<  >>`,
		Id:         "RelExpr",
		NTType:     16,
		Index:      28,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "<"	<<  >>`,
		Id:         "RelOp",
		NTType:     17,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : ">"	<<  >>`,
		Id:         "RelOp",
		NTType:     17,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "!="	<<  >>`,
		Id:         "RelOp",
		NTType:     17,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AddExpr : AddExpr "+" MulExpr	<<  >>`,
		Id:         "AddExpr",
		NTType:     18,
		Index:      32,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AddExpr : AddExpr "-" MulExpr	<<  >>`,
		Id:         "AddExpr",
		NTType:     18,
		Index:      33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AddExpr : MulExpr	<<  >>`,
		Id:         "AddExpr",
		NTType:     18,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MulExpr : MulExpr "*" Primary	<<  >>`,
		Id:         "MulExpr",
		NTType:     19,
		Index:      35,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MulExpr : MulExpr "/" Primary	<<  >>`,
		Id:         "MulExpr",
		NTType:     19,
		Index:      36,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MulExpr : Primary	<<  >>`,
		Id:         "MulExpr",
		NTType:     19,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Primary : "(" Expression ")"	<< data_structures.ReturnExpression(X[1].(data_structures.Tipo)) >>`,
		Id:         "Primary",
		NTType:     20,
		Index:      38,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return data_structures.ReturnExpression(X[1].(data_structures.Tipo))
		},
	},
	ProdTabEntry{
		String: `Primary : id	<< ctx.ResolveVarType(string(X[0].(*token.Token).Lit)) >>`,
		Id:         "Primary",
		NTType:     20,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ctx.ResolveVarType(string(X[0].(*token.Token).Lit))
		},
	},
	ProdTabEntry{
		String: `Primary : cte_int	<< data_structures.Int, nil >>`,
		Id:         "Primary",
		NTType:     20,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return data_structures.Int, nil
		},
	},
	ProdTabEntry{
		String: `Primary : cte_float	<< data_structures.Float, nil >>`,
		Id:         "Primary",
		NTType:     20,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return data_structures.Float, nil
		},
	},
	ProdTabEntry{
		String: `Print : "print" "(" ")" ";"	<<  >>`,
		Id:         "Print",
		NTType:     21,
		Index:      42,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Print : "print" "(" ArgList ")" ";"	<<  >>`,
		Id:         "Print",
		NTType:     21,
		Index:      43,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ArgList : Expression	<<  >>`,
		Id:         "ArgList",
		NTType:     22,
		Index:      44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ArgList : Expression "," ArgList	<<  >>`,
		Id:         "ArgList",
		NTType:     22,
		Index:      45,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ArgList : cte_string	<< data_structures.String, nil >>`,
		Id:         "ArgList",
		NTType:     22,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return data_structures.String, nil
		},
	},
	ProdTabEntry{
		String: `Cycle : "while" "(" Expression ")" "do" Body ";"	<<  >>`,
		Id:         "Cycle",
		NTType:     23,
		Index:      47,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Condition : "if" "(" Expression ")" Body "else" Body ";"	<<  >>`,
		Id:         "Condition",
		NTType:     24,
		Index:      48,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Condition : "if" "(" Expression ")" Body ";"	<<  >>`,
		Id:         "Condition",
		NTType:     24,
		Index:      49,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `F_Call : id "(" ")" ";"	<<  >>`,
		Id:         "F_Call",
		NTType:     25,
		Index:      50,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `F_Call : id "(" ArgList ")" ";"	<<  >>`,
		Id:         "F_Call",
		NTType:     25,
		Index:      51,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
}
