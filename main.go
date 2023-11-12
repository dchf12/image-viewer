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

	// 画像とそのインデックスを保持するための構造体
	type indexedImage struct {
		index int
		image *canvas.Image
	}

	imageCh := make(chan indexedImage, len(images))
	for i, image := range images {
		if image.IsDir() {
			continue
		}
		go func(i int, name string) {
			cImage := canvas.NewImageFromFile("./assets/" + name)
			cImage.FillMode = canvas.ImageFillOriginal
			imageCh <- indexedImage{i, cImage}
		}(i, image.Name())
	}

	cImages := make([]*canvas.Image, len(images))
	for range images {
		image := <-imageCh
		cImages[image.index] = image.image
	}

	window.SetContent(container.NewVBox(
		cImages[0],
	))

	window.Resize(fyne.NewSize(800, 600))

	window.Canvas().SetOnTypedKey(handleKeys)

	window.ShowAndRun()
}

var currentImage int

func handleKeys(e *fyne.KeyEvent) {
	switch e.Name {
	case fyne.KeyUp:
		// TODO: move to next image
		currentImage++
	case fyne.KeyDown:
		// TODO: move to previous image
		currentImage--
	case fyne.KeyLeft:
		// TODO: move to previous image
		currentImage--
	case fyne.KeyRight:
		// TODO: move to next image
		currentImage++
	}
}
