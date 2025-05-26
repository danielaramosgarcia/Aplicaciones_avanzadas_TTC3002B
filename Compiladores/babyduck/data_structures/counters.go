package data_structures

// Simulación de direcciones de memoria (stack) por segmentos:
const (
	// Variables globales [0,500]
	GlobalIntBase   = 1
	GlobalFloatBase = 250
	GlobalLimit     = 500

	// Variables locales [501,1000]
	LocalIntBase   = 501
	LocalFloatBase = 751
	LocalLimit     = 1000

	// Variables temporales [1001,1600]
	TempIntBase   = 1001
	TempFloatBase = 1201
	TempBoolBase  = 1401
	TempLimit     = 1600

	// Constantes [1601,2200]
	ConstIntBase    = 1601
	ConstFloatBase  = 1801
	ConstStringBase = 2001
	ConstLimit      = 2200
)

// Contadores para asignar la siguiente dirección disponible.
var (
	nextGlobalIntAddr   = GlobalIntBase
	nextGlobalFloatAddr = GlobalFloatBase

	nextLocalIntAddr   = LocalIntBase
	nextLocalFloatAddr = LocalFloatBase

	nextTempIntAddr   = TempIntBase
	nextTempFloatAddr = TempFloatBase
	nextTempBoolAddr  = TempBoolBase

	nextConstIntAddr    = ConstIntBase
	nextConstFloatAddr  = ConstFloatBase
	nextConstStringAddr = ConstStringBase
)

// Referencia a tabla global de variables.
var programName = ""
