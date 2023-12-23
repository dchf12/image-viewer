package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

const MaxImagesToLoad = 10

var (
	imageManager *ImageManager
	window       fyne.Window
)

func main() {
	fapp := app.New()
	window = fapp.NewWindow("viewer")

	imageManager = NewImageManager(MaxImagesToLoad)
	if err := imageManager.Load("./assets"); err != nil {
		log.Fatal(err)
	}

	window.SetContent(container.NewVBox(
		imageManager.Current(),
	))
	window.Resize(fyne.NewSize(600, 600))
	window.Canvas().SetOnTypedKey(handleKeys)
	window.ShowAndRun()
}

func handleKeys(e *fyne.KeyEvent) {
	switch e.Name {
	case fyne.KeyUp, fyne.KeyRight:
		imageManager.Next()
	case fyne.KeyDown, fyne.KeyLeft:
		imageManager.Prev()
	case fyne.KeyO:
		openDirectory()
	}
	updateImage()
}

func updateImage() {
	window.SetContent(container.NewVBox(
		imageManager.Current(),
	))
	window.Canvas().Refresh(window.Content().(*fyne.Container))
}

func openDirectory() {
	dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
		if uri != nil || err != nil {
			if err := imageManager.Load(uri.Path()); err != nil {
				log.Fatal(err)
			}
			updateImage()
		}
	}, window)
}
