package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	var emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)
	mg := NewMapGenerator(emptyImage, 51, 100, 100)

	if err := ebiten.Run(mg.Update, screenWidth, screenHeight, 1, "Marching Squares"); err != nil {
		log.Fatal(err)
	}
}
