package main

import (
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	fapp := app.New()
	window := fapp.NewWindow("viewer")

	images, err := ioutil.ReadDir("./assets")
	if err != nil {
		log.Fatal(err)
	}
	cImages := make([]*canvas.Image, len(images))
	for i, image := range images {
		cImages[i] = canvas.NewImageFromFile("./assets/" + image.Name())
		cImages[i].FillMode = canvas.ImageFillOriginal
	}
	window.SetContent(container.NewVBox(
		cImages[0],
	))

	window.Resize(fyne.NewSize(800, 600))

	window.Canvas().SetOnTypedKey(handleKeys)

	window.ShowAndRun()
}

func handleKeys(e *fyne.KeyEvent) {
	switch e.Name {
	case fyne.KeyUp:
		// TODO: move to next image
	case fyne.KeyDown:
		// TODO: move to previous image
	case fyne.KeyLeft:
		// TODO: move to previous image
	case fyne.KeyRight:
		// TODO: move to next image
	}
}
