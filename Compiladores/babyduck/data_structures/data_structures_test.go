package data_structures

import (
	"testing"
)

func TestVarTable(t *testing.T) {
	global := NewVarTable(nil)

	// 1) Añadir y recuperar variable global
	if err := global.Add("x", Int); err != nil {
		t.Fatalf("Add fallo: %v", err)
	}
	if entry, ok := global.Get("x"); !ok || entry.Type != Int {
		t.Errorf("Get devolvió %v, %v; quiero Int", entry, ok)
	}

	// 2) Duplicado
	if err := global.Add("x", Float); err == nil {
		t.Errorf("Add debería fallar por variable duplicada")
	}

	// 3) Encadenamiento de scopes
	local := NewVarTable(global)
	if _, ok := local.Get("x"); !ok {
		t.Errorf("Local no resolvió variable global")
	}
	if err := local.Add("y", Float); err != nil {
		t.Fatalf("Add local fallo: %v", err)
	}
	if entry, _ := local.Get("y"); entry.Type != Float {
		t.Errorf("Get local devolvió tipo incorrecto")
	}
}

func TestFuncDir(t *testing.T) {
	dir := NewFuncDir()

	// 1) Registrar una función
	f := &FuncEntry{Name: "f", ReturnType: Void, ParamTypes: []Tipo{Int, Float}, VarTable: NewVarTable(nil)}
	if err := dir.Add(f); err != nil {
		t.Fatalf("AddFunc fallo: %v", err)
	}

	// 2) Duplicado
	if err := dir.Add(f); err == nil {
		t.Errorf("AddFunc debería fallar por duplicado")
	}

	// 3) Recuperar
	if got, ok := dir.Get("f"); !ok || got != f {
		t.Errorf("GetFunc devolvió %v, %v; quiero %v", got, ok, f)
	}
}

func TestContextHelpers(t *testing.T) {
	ctx := NewContext()

	// Global vars
	if _, err := ctx.RegisterGlobalVars([]string{"a", "b"}, Int); err != nil {
		t.Fatalf("RegisterGlobalVars fallo: %v", err)
	}
	for _, name := range []string{"a", "b"} {
		if _, ok := ctx.GlobalVars.Get(name); !ok {
			t.Errorf("GlobalVars no contiene %s", name)
		}
	}

	// Función
	if _, err := ctx.RegisterFunction("foo", Void, []Tipo{Float}); err != nil {
		t.Fatalf("RegisterFunction fallo: %v", err)
	}
	if !ctx.FuncDir.Exists("foo") {
		t.Errorf("FuncDir no contiene 'foo'")
	}

	// Resolver variable dentro de función activa
	if err := ctx.EnterFunction("foo"); err != nil {
		t.Fatalf("EnterFunction fallo: %v", err)
	}
	if _, err := ctx.ValidateAssign("x", Int); err == nil {
		// x no existe aún
	}
	// if err := ctx.ExitFunction(); err != nil {
	// 	// ExitFunction no retorna error, solo limpia
	// }

	// ResolveVarType en global
	if _, err := ctx.ResolveVarType("a"); err != nil {
		t.Errorf("ResolveVarType fallo para global 'a': %v", err)
	}
}
