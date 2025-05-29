package data_structures

import (
	"fmt"
	"strconv"
)

// METODOS DE TABLA DE VARIABLES
// NewVarTable crea una nueva tabla de variables, opcionalmente enlazada a un padre.
func NewVarTable(Parent *VarTable) *VarTable {
	return &VarTable{
		Vars:   make(map[int]*VarEntry),
		Parent: Parent,
	}
}

// Add inserta una nueva variable; asigna una dirección según el tipo y ámbito y devuelve la dirección.
// Error si ya existe una variable con el mismo nombre en este ámbito.
func (vt *VarTable) Add(name string, typ int) error {
	// Validar duplicado por nombre
	for _, entry := range vt.Vars {
		if entry.Name == name {
			return fmt.Errorf("variable %q ya declarada en este ámbito", name)
		}
	}
	// Asignar dirección según segmento
	var dir int
	if vt.Parent == nil {
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
	vt.Vars[dir] = &VarEntry{Name: name, Type: typ, DirInt: dir}
	return nil
}

// AddTemp inserta una variable temporal sin nombre, asignando dirección según tipo.
// Devuelve la dirección o error en caso de overflow.
func (vt *VarTable) AddTemp(typ int) (int, error) {
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
	vt.Vars[dir] = &VarEntry{Name: "", Type: typ, DirInt: dir}
	return dir, nil
}

// AddTemp inserta una variable temporal sin nombre, asignando dirección según tipo.
// Devuelve la dirección o error en caso de overflow.
func (ctx *Context) AddConst(typ int, val string) (int, error) {
	var dir int
	found := false
	for _, v := range ctx.AddedConst {
		if v == val {
			found = true
			break
		}
	}

	if !found {
		switch typ {
		case 0:
			if nextConstIntAddr > ConstFloatBase-1 {
				return 0, fmt.Errorf("overflow de direcciones constantes int")
			}

			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Error al convertir:", err)
			}
			dir = nextConstIntAddr
			nextConstIntAddr++
			ctx.ConstTable.Num[dir] = num
		case 1:
			if nextConstFloatAddr > ConstStringBase-1 {
				return 0, fmt.Errorf("overflow de direcciones constantes float")
			}
			num, err := strconv.ParseFloat(val, 64)
			if err != nil {
				// Manejo de error si la cadena no es un float válido
				fmt.Printf("Error convirtiendo '%s' a float64: %v\n", val, err)
				return 0, fmt.Errorf("ERR al convertir float en add const")
			}
			dir = nextConstFloatAddr
			nextConstFloatAddr++
			ctx.ConstTable.Float[dir] = num

		case 4:
			if nextConstStringAddr > ConstLimit {
				return 0, fmt.Errorf("overflow de direcciones constantes string")
			}

			dir = nextConstStringAddr
			nextConstStringAddr++
			ctx.ConstTable.Str[dir] = val

		default:
			return 0, fmt.Errorf("tipo %v no soportado en constantes", typ)
		}
		ctx.AddedConst = append(ctx.AddedConst, val)

	} else {
		fmt.Println("Constante ya existe, buscando dirección...")
		// Si ya existe, buscar la dirección
		switch typ {
		case 0:
			for dir, num := range ctx.ConstTable.Num {
				if strconv.Itoa(num) == val {
					return dir, nil
				}
			}
		case 1:
			for dir, num := range ctx.ConstTable.Float {
				if fmt.Sprintf("%f", num) == val {
					return dir, nil
				}
			}
		case 4:
			for dir, str := range ctx.ConstTable.Str {
				if str == val {
					return dir, nil
				}
			}
		default:
			return 0, fmt.Errorf("tipo %v no soportado en constantes", typ)
		}
	}

	return dir, nil
}

// Get busca una variable por dirección; devuelve la entrada y true si existe.
func (vt *VarTable) Get(dir int) (*VarEntry, bool) {
	if e, ok := vt.Vars[dir]; ok {
		return e, true
	}
	if vt.Parent != nil {
		return vt.Parent.Get(dir)
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
	list := make([]*VarEntry, 0, len(vt.Vars))
	for _, entry := range vt.Vars {
		list = append(list, entry)
	}
	return list
}

// METODOS DE DIRECTORIO DE FUNCIONES
// NewFuncDir inicializa un nuevo directorio de funciones.
func NewFuncDir() *FuncDir {
	return &FuncDir{Funcs: make(map[string]*FuncEntry)}
}

// Add inserta una nueva función; error si ya existe.
func (fd *FuncDir) Add(f *FuncEntry) error {
	if _, exists := fd.Funcs[f.Name]; exists {
		return fmt.Errorf("función %q ya definida", f.Name)
	}
	fd.Funcs[f.Name] = f
	return nil
}

// Get devuelve la función registrada y true si existe.
func (fd *FuncDir) Get(name string) (*FuncEntry, bool) {
	f, ok := fd.Funcs[name]
	return f, ok
}

// Exists retorna true si la función está en el directorio.
func (fd *FuncDir) Exists(name string) bool {
	_, ok := fd.Funcs[name]
	return ok
}

// List devuelve todas las funciones registradas.
func (fd *FuncDir) List() []*FuncEntry {
	list := make([]*FuncEntry, 0, len(fd.Funcs))
	for _, entry := range fd.Funcs {
		list = append(list, entry)
	}
	return list
}
