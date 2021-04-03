package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Master Academy!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi! This is first desktop application", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
