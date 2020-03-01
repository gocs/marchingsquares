package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"

	"github.com/gocs/marchingsquares"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	mg := marchingsquares.NewMapGenerator(51, 100, 100, 10)

	emptyImage, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)
	g := &game{mg, emptyImage}

	if err := ebiten.Run(g.update, screenWidth, screenHeight, 1, "Marching Squares"); err != nil {
		log.Fatal(err)
	}
}

type game struct {
	mg  *marchingsquares.MapGenerator
	img *ebiten.Image
}

func (g *game) update(scr *ebiten.Image) error {
	if err := g.mg.GenerateMap(); err != nil {
		return errors.New(fmt.Sprint("error while computing:", err))
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawTrianglesOptions{}
	v, i := g.mg.GetTriangles()
	scr.DrawTriangles(v, i, g.img, op)

	return nil
}
