package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Rook 🐦‍⬛")
	w.SetContent(widget.NewLabel("Rook"))
	w.ShowAndRun()
}
