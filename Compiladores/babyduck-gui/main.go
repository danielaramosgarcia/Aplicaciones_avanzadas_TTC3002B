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
	w := a.NewWindow("Hecho con amor 💖 ")
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
		`Da click al boton de compilar para ver el resultado aca 🥳`,
		fyne.TextAlignLeading,
		fyne.TextStyle{Italic: true},
	)
	outputScroll := container.NewScroll(outputLabel)

	// --- 2) Botón Compile que usa esos widgets ---
	// --- 2) Botón Compile de color verde claro ---

	compileBtn := widget.NewButton("❤️ Compilar ❤️", func() {
		result := VM.CodeInput(codeEntry.Text)
		outputLabel.SetText(result)
	})

	// Creamos un rectángulo verde claro como fondo
	greenBg := canvas.NewRectangle(color.NRGBA{R: 200, G: 255, B: 200, A: 255})
	// Superponemos el botón sobre él
	btnContainer := container.NewStack(greenBg, compileBtn)

	// --- 3) Cabecera con título y botón alineados a extremos ---

	title := widget.NewLabelWithStyle(
		"Compilador para BabyDuck",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)

	nombre := widget.NewLabelWithStyle(
		"Dani Ramos A01174259",
		fyne.TextAlignLeading,
		fyne.TextStyle{Italic: false},
	)

	headerRow := container.NewHBox(
		layout.NewSpacer(),
		nombre,
		layout.NewSpacer(),
		layout.NewSpacer(),
		title,
		layout.NewSpacer(),
		layout.NewSpacer(),
		btnContainer,
		layout.NewSpacer(),
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
