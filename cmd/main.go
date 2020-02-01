package main

import (
	"image/color"
	"log"

	"github.com/gocs/marchingsquares"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)

func init() {
	emptyImage.Fill(color.White)
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Polygons (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}

var locx float32
var locy float32

func update(screen *ebiten.Image) error {

	_, dy := ebiten.Wheel()
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		locx -= float32(dy)
	} else {
		locy -= float32(dy)
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	atlas := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 4, 12, 8, 0, 0, 0, 0, 0, 0},
		{0, 6, 0, 9, 0, 0, 0, 0, 0, 0},
		{0, 2, 3, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	vertices, indexes := marchingsquares.GenerateSquares(atlas, locx, locy, 100, 100)

	op := &ebiten.DrawTrianglesOptions{}
	for _, triangle := range vertices {
		screen.DrawTriangles(triangle, indexes, emptyImage, op)
	}

	return nil
}
