package main

import (
	"image/color"
	"log"

	"github.com/gocs/marchingsquares"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
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

var locX float32 = -100
var locY float32 = -100

var initX float32
var initY float32

var diffX float32
var diffY float32

func update(screen *ebiten.Image) error {

	currXint, currYint := ebiten.CursorPosition()
	currX, currY := float32(currXint), float32(currYint)
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		initX = currX
		initY = currY
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		diffX = currX - initX
		diffY = currY - initY
		locX += diffX
		locY += diffY
	}

	atlas := [][]int{
		{15, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 4, 12, 8, 0, 0, 0, 0, 0, 0},
		{0, 6, 0, 9, 0, 0, 0, 0, 0, 0},
		{0, 2, 3, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	vertices, indexes := marchingsquares.GenerateSquares(atlas, locX, locY, 100, 100)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawTrianglesOptions{}
	for _, triangle := range vertices {
		screen.DrawTriangles(triangle, indexes, emptyImage, op)
	}

	return nil
}
