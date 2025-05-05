package data_structures

import "fmt"

// METODOS DE TABLA DE VARIABLES

// NewVarTable crea una nueva tabla de variables, opcionalmente enlazada a un padre.
func NewVarTable(parent *VarTable) *VarTable {
	return &VarTable{vars: make(map[string]*VarEntry), parent: parent}
}

// Add inserta una nueva variable; error si ya existe en este ámbito.
func (vt *VarTable) Add(name string, typ Tipo) error {
	if _, exists := vt.vars[name]; exists {
		return fmt.Errorf("variable %q ya declarada en este ámbito", name)
	}
	vt.vars[name] = &VarEntry{Name: name, Type: typ}
	return nil
}

// Get busca una variable en este ámbito o en los padres; devuelve entrada y true si existe.
func (vt *VarTable) Get(name string) (*VarEntry, bool) {
	if e, ok := vt.vars[name]; ok {
		return e, true
	}
	if vt.parent != nil {
		return vt.parent.Get(name)
	}
	return nil, false
}

// Exists retorna true si la variable está declarada en este ámbito o en sus padres.
func (vt *VarTable) Exists(name string) bool {
	_, ok := vt.Get(name)
	return ok
}

// List devuelve todas las variables locales de este ámbito.
func (vt *VarTable) List() []*VarEntry {
	list := make([]*VarEntry, 0, len(vt.vars))
	for _, entry := range vt.vars {
		list = append(list, entry)
	}
	return list
}

// METODOS DE DIRECTORIO DE FUNCIONES

// NewFuncDir inicializa un nuevo directorio de funciones.
func NewFuncDir() *FuncDir {
	return &FuncDir{funcs: make(map[string]*FuncEntry)}
}

// Add inserta una nueva función; error si ya existe.
func (fd *FuncDir) Add(f *FuncEntry) error {
	if _, exists := fd.funcs[f.Name]; exists {
		return fmt.Errorf("función %q ya definida", f.Name)
	}
	fd.funcs[f.Name] = f
	return nil
}

// Get devuelve la función registrada y true si existe.
func (fd *FuncDir) Get(name string) (*FuncEntry, bool) {
	f, ok := fd.funcs[name]
	return f, ok
}

// Exists retorna true si la función está en el directorio.
func (fd *FuncDir) Exists(name string) bool {
	_, ok := fd.funcs[name]
	return ok
}

// List devuelve todas las funciones registradas.
func (fd *FuncDir) List() []*FuncEntry {
	list := make([]*FuncEntry, 0, len(fd.funcs))
	for _, entry := range fd.funcs {
		list = append(list, entry)
	}
	return list
}
