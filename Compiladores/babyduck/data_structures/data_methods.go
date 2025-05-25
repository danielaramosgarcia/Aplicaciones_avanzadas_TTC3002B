package data_structures

import (
	"fmt"
)

// METODOS DE TABLA DE VARIABLES

// NewVarTable crea una nueva tabla de variables, opcionalmente enlazada a un padre.
func NewVarTable(parent *VarTable) *VarTable {
	return &VarTable{
		vars:   make(map[int]*VarEntry),
		parent: parent,
	}
}

// Add inserta una nueva variable; asigna una dirección según el tipo y ámbito y devuelve la dirección.
// Error si ya existe una variable con el mismo nombre en este ámbito.
func (vt *VarTable) Add(name string, typ int) error {
	// Validar duplicado por nombre
	for _, entry := range vt.vars {
		if entry.Name == name {
			return fmt.Errorf("variable %q ya declarada en este ámbito", name)
		}
	}
	// Asignar dirección según segmento
	var dir int
	if vt.parent == nil {
		// global
		if typ == 0 {
			if nextGlobalIntAddr >= GlobalFloatBase {
				return fmt.Errorf("overflow de direcciones globales int")
			}
			dir = nextGlobalIntAddr
			nextGlobalIntAddr++
		} else {
			if nextGlobalFloatAddr >= GlobalLimit {
				return fmt.Errorf("overflow de direcciones globales float")
			}
			dir = nextGlobalFloatAddr
			nextGlobalFloatAddr++
		}
	} else {
		// local
		if typ == 0 {
			if nextLocalIntAddr >= LocalFloatBase {
				return fmt.Errorf("overflow de direcciones locales int")
			}
			dir = nextLocalIntAddr
			nextLocalIntAddr++
		} else {
			if nextLocalFloatAddr >= LocalLimit {
				return fmt.Errorf("overflow de direcciones locales float")
			}
			dir = nextLocalFloatAddr
			nextLocalFloatAddr++
		}
	}
	vt.vars[dir] = &VarEntry{Name: name, Type: typ, DirInt: dir}
	return nil
}

// AddTemp inserta una variable temporal sin nombre, asignando dirección según tipo.
// Devuelve la dirección o error en caso de overflow.
func (vt *VarTable) AddTemp(typ Tipo) (int, error) {
	var dir int
	switch typ {
	case 0:
		if nextTempIntAddr > TempFloatBase-1 {
			return 0, fmt.Errorf("overflow de direcciones temporales int")
		}
		dir = nextTempIntAddr
		nextTempIntAddr++
	case 1:
		if nextTempFloatAddr > TempBoolBase-1 {
			return 0, fmt.Errorf("overflow de direcciones temporales float")
		}
		dir = nextTempFloatAddr
		nextTempFloatAddr++
	case 2:
		if nextTempBoolAddr > TempLimit {
			return 0, fmt.Errorf("overflow de direcciones temporales bool")
		}
		dir = nextTempBoolAddr
		nextTempBoolAddr++
	default:
		return 0, fmt.Errorf("tipo %v no soportado en temporales", typ)
	}
	vt.vars[dir] = &VarEntry{Name: "", Type: 0, DirInt: dir}
	return dir, nil
}

// AddTemp inserta una variable temporal sin nombre, asignando dirección según tipo.
// Devuelve la dirección o error en caso de overflow.
func (vt *VarTable) AddConst(typ int) (int, error) {
	var dir int
	switch typ {
	case 0:
		if nextConstIntAddr > ConstFloatBase-1 {
			return 0, fmt.Errorf("overflow de direcciones constantes int")
		}
		dir = nextConstIntAddr
		nextConstIntAddr++
	case 1:
		if nextConstFloatAddr > ConstStringBase-1 {
			return 0, fmt.Errorf("overflow de direcciones constantes float")
		}
		dir = nextConstFloatAddr
		nextConstFloatAddr++
	case 4:
		if nextConstStringAddr > ConstLimit {
			return 0, fmt.Errorf("overflow de direcciones constantes string")
		}
		dir = nextConstStringAddr
		nextConstStringAddr++
	default:
		return 0, fmt.Errorf("tipo %v no soportado en constantes", typ)
	}
	vt.vars[dir] = &VarEntry{Name: "", Type: typ, DirInt: dir}
	return dir, nil
}

// Get busca una variable por dirección; devuelve la entrada y true si existe.
func (vt *VarTable) Get(dir int) (*VarEntry, bool) {
	if e, ok := vt.vars[dir]; ok {
		return e, true
	}
	if vt.parent != nil {
		return vt.parent.Get(dir)
	}
	return nil, false
}

// Exists retorna true si la variable con esa dirección está en este ámbito o en sus padres.
func (vt *VarTable) Exists(dir int) bool {
	_, ok := vt.Get(dir)
	return ok
}

// List devuelve todas las variables del ámbito actual.
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
