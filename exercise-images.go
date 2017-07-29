package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width  int
	Height int
}

func (image Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{21, 21, 21, 255}
}

func main() {
	m := Image{12, 12}
	pic.ShowImage(m)
}
