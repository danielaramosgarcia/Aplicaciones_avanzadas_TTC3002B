package main

import (
	// ajusta al import de tu VM
	"babyduck/VM"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Babyduck IDE")

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

	codeEntry.SetMinRowsVisible(20) // Ajusta la altura mínima visible del editor

	codeEntry.SetPlaceHolder("Write your BabyDuck code here...")
	// Output
	outputLabel := widget.NewLabelWithStyle(
		`Click "Compile" to see output here...`,
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)

	outputScroll := container.NewScroll(outputLabel)

	// Botón Compile
	compileBtn := widget.NewButton("Compile ▶", func() {
		result := VM.CodeInput(codeEntry.Text)
		outputLabel.SetText(result)
	})

	// Contenedor para el editor de código con su título
	// Usamos container.NewBorder para añadir el título y luego el widget de entrada.
	// La altura mínima se gestionará por el HSplit y el tamaño de la ventana.
	codeContainer := container.NewBorder(
		widget.NewLabelWithStyle("Source Code", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		nil, nil, nil,
		codeEntry,
	)

	// Contenedor para el output con su título
	outputContainer := container.NewBorder(
		widget.NewLabelWithStyle("Output", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		nil, nil, nil,
		outputScroll,
	)

	// Los dos campos de texto (codeContainer y outputContainer) uno al lado del otro horizontalmente
	// NewHSplit se encargará de distribuir el espacio.
	textFields := container.NewHSplit(codeContainer, outputContainer)
	textFields.Offset = 0.5 // Puedes ajustar el offset para el tamaño inicial

	// Contenido principal: El botón arriba, y luego los campos de texto
	content := container.NewVBox(
		compileBtn,
		textFields,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600)) // Establece un tamaño inicial de ventana que acomode los elementos.
	w.ShowAndRun()
}
