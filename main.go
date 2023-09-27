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
	image := canvas.NewImageFromResource(nil)
	image.FillMode = canvas.ImageFillOriginal

	content := container.NewMax(bg, image)
	cv.SetContent(content)
	window.Resize(fyne.NewSize(800, 600))

	window.ShowAndRun()
}
