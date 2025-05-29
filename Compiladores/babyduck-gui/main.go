package main

import (
	"babyduck/VM"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Go Compiler IDE")
	w.Padded()
	w.SetPadded(true)

	// --- 1) Creamos los widgets centrales primero ---

	// Editor de código
	codeEntry := widget.NewMultiLineEntry()
	codeEntry.SetText(`program p;
var x: int;
main {
  x = 2 + 2;
  print(x);
}
end`)
	codeEntry.Wrapping = fyne.TextWrapOff

	// Output (etiqueta + scroll)
	outputLabel := widget.NewLabelWithStyle(
		`Click "Compile" to see output here...`,
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)
	outputScroll := container.NewScroll(outputLabel)

	// --- 2) Botón Compile que usa esos widgets ---
	// --- 2) Botón Compile de color verde claro ---

	compileBtn := widget.NewButton("Compile ❤️", func() {
		result := VM.CodeInput(codeEntry.Text)
		outputLabel.SetText(result)
	})
	// Creamos un rectángulo verde claro como fondo
	greenBg := canvas.NewRectangle(color.NRGBA{R: 200, G: 255, B: 200, A: 255})
	// Superponemos el botón sobre él
	btnContainer := container.NewStack(greenBg, compileBtn)

	// --- 3) Cabecera con título y botón alineados a extremos ---

	title := widget.NewLabelWithStyle(
		"Go Compiler IDE",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)
	headerRow := container.NewHBox(
		title,
		layout.NewSpacer(),
		btnContainer,
	)

	// --- 4) Paneles izquierdo/derecho con su encabezado ---

	left := container.NewBorder(
		widget.NewLabelWithStyle(
			"Source Code",
			fyne.TextAlignLeading,
			fyne.TextStyle{Bold: true},
		),
		nil, nil, nil,
		codeEntry,
	)
	right := container.NewBorder(
		widget.NewLabelWithStyle(
			"Output",
			fyne.TextAlignLeading,
			fyne.TextStyle{Bold: true},
		),
		nil, nil, nil,
		outputScroll,
	)

	// --- 5) Split horizontal que reparte 50/50 y se expande ---

	split := container.NewHSplit(left, right)
	split.Offset = 0.5

	// --- 6) Composición final: header arriba, split rellena el resto ---

	content := container.NewBorder(
		headerRow, // top
		nil,       // bottom
		nil,       // left
		nil,       // right
		split,     // center
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(900, 600))
	w.ShowAndRun()
}
