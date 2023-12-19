package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const MaxImagesToLoad = 10

// 画像とそのインデックスを保持するための構造体
type indexedImage struct {
	index int
	image *canvas.Image
}

var (
	currentImage int
	cImages      []*canvas.Image
	window       fyne.Window
)

func main() {
	fapp := app.New()
	window = fapp.NewWindow("viewer")

	images, err := os.ReadDir("./assets")
	if err != nil {
		log.Fatal(err)
	}

	loadImages(images)

	window.SetContent(container.NewVBox(
		cImages[0],
	))

	window.Resize(fyne.NewSize(600, 600))
	window.Canvas().SetOnTypedKey(handleKeys)
	window.ShowAndRun()
}

func loadImages(images []os.DirEntry) {
	loadNum := min(len(images), MaxImagesToLoad)
	imageCh := make(chan indexedImage, loadNum)
	for i, image := range images {
		log.Printf("Loading image %d: %s", i, image.Name())
		if i >= MaxImagesToLoad {
			break
		}
		if image.IsDir() {
			continue
		}
		go loadImage(i, image.Name(), imageCh)
	}

	cImages = make([]*canvas.Image, loadNum)
	for i := 0; i < loadNum; i++ {
		img := <-imageCh
		if img.image == nil {
			log.Printf("Error loading image at index %d", img.index)
			continue
		}
		cImages[img.index] = img.image
	}

	if len(cImages) == 0 {
		log.Fatal("No images found in assets directory")
	}
}

func loadImage(index int, name string, imageCh chan<- indexedImage) {
	cImage := canvas.NewImageFromFile("./assets/" + name)
	if cImage == nil {
		imageCh <- indexedImage{index, nil}
		return
	}
	cImage.FillMode = canvas.ImageFillOriginal
	imageCh <- indexedImage{index, cImage}
}

func handleKeys(e *fyne.KeyEvent) {
	switch e.Name {
	case fyne.KeyUp, fyne.KeyRight:
		currentImage = (currentImage + 1) % len(cImages)
	case fyne.KeyDown, fyne.KeyLeft:
		currentImage = (currentImage - 1 + len(cImages)) % len(cImages)
	}

	if currentImage < 0 {
		currentImage = 0
	} else if currentImage >= len(cImages) {
		currentImage = len(cImages) - 1
	}

	updateImage()
}

func updateImage() {
	if currentImage < 0 || currentImage >= len(cImages) || cImages[currentImage] == nil {
		log.Printf("No image to display at index %d", currentImage)
	}
	window.SetContent(container.NewVBox(
		cImages[currentImage],
	))
	window.Canvas().Refresh(window.Content().(*fyne.Container))
}
