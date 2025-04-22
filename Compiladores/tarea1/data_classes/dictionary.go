package data_classes

import "fmt"

// Struck para la estructura de datos Dictionary
type Dictionary struct {
	Keys   []string
	Values []string
}

// Metodo para agregar un elemento al diccionario
func (d *Dictionary) Add(key, value string) {
	d.Keys = append(d.Keys, key)
	d.Values = append(d.Values, value)
}

// Metodo para obtener un elemento del diccionario
func (d *Dictionary) Get(key string) (string, bool) {
	for i, k := range d.Keys {
		if k == key {
			return d.Values[i], true
		}
	}
	return "", false
}

// Metodo para eliminar un elemento del diccionario
func (d *Dictionary) Remove(key string) bool {
	for i, k := range d.Keys {
		if k == key {
			d.Keys = append(d.Keys[:i], d.Keys[i+1:]...)
			d.Values = append(d.Values[:i], d.Values[i+1:]...)
			return true
		}
	}
	return false
}

// Metodo para imprimir el diccionario
func (d *Dictionary) Print() {
	for i, k := range d.Keys {
		fmt.Printf("%s - %s\n", k, d.Values[i])
	}
}

// Metodo para verificar si el diccionario está vacío
func (d *Dictionary) IsEmpty() bool {
	return len(d.Keys) == 0
}

// Metodo para obtener el tamaño del diccionario
func (d *Dictionary) Size() int {
	return len(d.Keys)
}

// Metodo para verificar si una clave existe en el diccionario
func (d *Dictionary) ContainsKey(key string) bool {
	for _, k := range d.Keys {
		if k == key {
			return true
		}
	}
	return false
}

// Metodo para verificar si un valor existe en el diccionario
func (d *Dictionary) ContainsValue(value string) bool {
	for _, v := range d.Values {
		if v == value {
			return true
		}
	}
	return false
}
