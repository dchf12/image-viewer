package main

import (
	"log"
	"os"
	"path"
	"strings"

	"fyne.io/fyne/v2/canvas"
)

// 画像とそのインデックスを保持するための構造体
type indexedImage struct {
	index int
	image *canvas.Image
}

type ImageManager struct {
	images    []*canvas.Image
	current   int
	MaxToLoad int
}

func NewImageManager(maxImages int) *ImageManager {
	return &ImageManager{
		MaxToLoad: maxImages,
	}
}

func (m *ImageManager) Load(dir string) error {
	images, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	imageCh := make(chan indexedImage, m.MaxToLoad)
	for i, image := range images {
		log.Printf("Loading image %d: %s", i, image.Name())
		if i >= m.MaxToLoad {
			break
		}
		if image.IsDir() {
			continue
		}
		if strings.HasPrefix(image.Name(), ".") {
			continue
		}
		go m.load(i, path.Join(dir, image.Name()), imageCh)
	}

	for i := 0; i < min(len(images), m.MaxToLoad); i++ {
		img := <-imageCh
		if img.image == nil {
			log.Printf("Error loading image at index %d", img.index)
			continue
		}
		m.images = append(m.images, img.image)
	}

	if len(m.images) == 0 {
		log.Fatal("No images found in assets directory")
	}
	return nil
}

func (m *ImageManager) load(index int, filePath string, imageCh chan<- indexedImage) {
	cImage := canvas.NewImageFromFile(filePath)
	if cImage == nil {
		imageCh <- indexedImage{index, nil}
		return
	}
	cImage.FillMode = canvas.ImageFillOriginal
	imageCh <- indexedImage{index, cImage}
}

func (m *ImageManager) Current() *canvas.Image {
	if m.current < 0 || m.current >= len(m.images) {
		log.Printf("No image to display at index %d", m.current)
		return nil
	}
	return m.images[m.current]
}

func (m *ImageManager) Next() {
	m.current = (m.current + 1) % len(m.images)
}

func (m *ImageManager) Prev() {
	m.current = (m.current - 1 + len(m.images)) % len(m.images)
}
