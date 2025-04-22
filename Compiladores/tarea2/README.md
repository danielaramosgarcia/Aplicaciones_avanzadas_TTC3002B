# Tarea 2 – Evaluador de Expresiones Aritméticas con Goyacc y Lexer en Go

Este proyecto implementa un evaluador de expresiones aritméticas utilizando únicamente herramientas nativas del lenguaje **Go**:

- `goyacc`: para generar el **parser** (analizador sintáctico).
- `text/scanner`: para construir el **lexer** (analizador léxico).

---

## 🧠 ¿Qué hace?

Este programa analiza y evalúa expresiones matemáticas con operaciones básicas como:

- Suma (`+`)
- Resta (`-`)
- Multiplicación (`*`)
- División (`/`)
- Uso de paréntesis

**Ejemplo de entrada:**

```
3 + 4 * (2 + 1)
```

**Salida esperada:**

```
Resultado: 15
```

---

## 🚀 Cómo ejecutarlo

### 1. Instala Go (si aún no lo tienes)

[https://go.dev/doc/install](https://go.dev/doc/install)

---

### 2. Instala `goyacc` (una sola vez)

```bash
go install golang.org/x/tools/cmd/goyacc@latest
```

Asegúrate de tener `$GOPATH/bin` en tu `PATH`. Puedes hacerlo así:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Agrega esa línea a tu archivo `~/.zshrc` o `~/.bashrc` para que se cargue automáticamente en futuras sesiones.

---

### 3. Genera el parser

Desde el directorio del proyecto (`tarea2/`):

```bash
goyacc -o parser.go parser.y
```

Este comando genera el archivo `parser.go` a partir de tu gramática.

---

### 4. Ejecuta el programa

```bash
go run .
```

Puedes modificar la expresión de entrada directamente en el archivo `main.go` para probar diferentes operaciones.

---

## 📌 Notas

- Toda la lógica está escrita en Go puro.
- El resultado de la expresión se guarda en una variable global `result` y se imprime desde `main.go`.

---

## 👩‍💻 Autor

**Daniela Ramos García**  
*Tarea 2 – Evaluador de expresiones usando goyacc + lexer en Go*
