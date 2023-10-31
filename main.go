package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	fapp := app.New()
	window := fapp.NewWindow("viewer")
	cv := window.Canvas()

	bg := canvas.NewRectangle(color.RGBA{0, 0, 0, 255})
	image := canvas.NewImageFromFile("./assets/blue.png")
	image.FillMode = canvas.ImageFillOriginal

	content := container.NewMax(bg, image)
	cv.SetContent(content)
	window.Resize(fyne.NewSize(800, 600))

	// keyboard event
	window.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		switch e.Name {
		case fyne.KeyUp:
			image.Move(fyne.NewPos(0, -10))
		case fyne.KeyDown:
			image.Move(fyne.NewPos(0, 10))
		case fyne.KeyLeft:
			image.Move(fyne.NewPos(-10, 0))
		case fyne.KeyRight:
			image.Move(fyne.NewPos(10, 0))
		}
	})

	window.ShowAndRun()
}
