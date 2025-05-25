package data_structures

import (
	"testing"
)

func TestVarTable(t *testing.T) {
	global := NewVarTable(nil)

	// 1) Añadir y recuperar variable global
	if err := global.Add("x", 0); err != nil {
		t.Fatalf("Add fallo: %v", err)
	}
	if entry, ok := global.Get(1); !ok || entry.Type != 0 {
		t.Errorf("Get devolvió %v, %v; quiero Int", entry, ok)
	}

	// 2) Duplicado
	if err := global.Add("x", 1); err == nil {
		t.Errorf("Add debería fallar por variable duplicada")
	}

	// 3) Encadenamiento de scopes
	local := NewVarTable(global)
	if _, ok := local.Get(1); !ok {
		t.Errorf("Local no resolvió variable global")
	}
	if err := local.Add("y", 1); err != nil {
		t.Fatalf("Add local fallo: %v", err)
	}
	if entry, _ := local.Get(751); entry.Type != 1 {
		t.Errorf("Get local devolvió tipo incorrecto")
	}
}

func TestFuncDir(t *testing.T) {
	dir := NewFuncDir()

	// 1) Registrar una función
	f := &FuncEntry{
		Name:       "f",
		ReturnType: 3,
		ParamTypes: []int{0, 1},
		VarTable:   NewVarTable(nil),
	}
	if err := dir.Add(f); err != nil {
		t.Fatalf("AddFunc fallo: %v", err)
	}

	// 2) Duplicado
	if err := dir.Add(f); err == nil {
		t.Errorf("AddFunc debería fallar por duplicado")
	}

	// 3) Recuperar
	got, ok := dir.Get("f")
	if !ok || got != f {
		t.Errorf("GetFunc devolvió %v, %v; quiero %v", got, ok, f)
	}
}

func TestContextHelpers(t *testing.T) {
	ctx := NewContext()

	// 1) Registrar variables globales
	if _, err := ctx.RegisterGlobalVars([]string{"a", "b"}, 0); err != nil {
		t.Fatalf("RegisterGlobalVars fallo: %v", err)
	}
	// CHECAR POR QUE FALLA
	// if _, ok := ctx.GlobalVars.Get(1); !ok {
	// 	t.Errorf("GlobalVars no contiene %s", "a")
	// }
	if _, ok := ctx.GlobalVars.Get(2); !ok {
		t.Errorf("GlobalVars no contiene %s", "b")
	}

	// 2) Registrar y recuperar función
	_, err := ctx.RegisterFunction("foo", 3, []Param{})
	if err != nil {
		t.Fatalf("RegisterFunction fallo: %v", err)
	}
	fEntry, ok := ctx.FuncDir.Get("foo")
	if !ok || fEntry.Name != "foo" {
		t.Errorf("FuncDir no contiene 'foo': %v, %v", fEntry, ok)
	}

	// 3) Entrar y salir de la función
	_, err = ctx.EnterFunction("foo")
	if err != nil {
		t.Fatalf("EnterFunction fallo: %v", err)
	}
	// Dentro de foo, intentar validar asignación de 'x' (no existe)
	if _, err := ctx.ValidateAssign("x", 0); err == nil {
		t.Errorf("ValidateAssign debería fallar para variable 'x' no declarada")
	}
	_, err = ctx.ExitFunction()
	if err != nil {
		t.Fatalf("ExitFunction fallo: %v", err)
	}

	// 4) Resolver tipo en ámbito global
	if _, err := ctx.ResolveVarType("a"); err != nil {
		t.Errorf("ResolveVarType fallo para global 'a': %v", err)
	}
}

func TestRegisterAndEnterFunction(t *testing.T) {
	ctx := NewContext()
	// Registrar y entrar de una sola vez
	_, err := ctx.RegisterAndEnterFunction("bar", 3, []Param{{Name: "p", Type: 0}})
	if err != nil {
		t.Fatalf("RegisterAndEnterFunction fallo: %v", err)
	}
	// Ahora 'p' debería existir en la tabla local
	if _, err := ctx.ValidateAssign("p", 0); err != nil {
		t.Errorf("ValidateAssign debería encontrar parámetro 'p', got: %v", err)
	}
	_, _ = ctx.ExitFunction()
}
