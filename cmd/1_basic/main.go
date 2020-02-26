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
	var emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)
	mg := marchingsquares.NewMapGenerator(emptyImage, 51, 100, 100)

	g := &game{mg}

	if err := ebiten.Run(g.update, screenWidth, screenHeight, 1, "Marching Squares"); err != nil {
		log.Fatal(err)
	}
}


type game struct {
	mg *marchingsquares.MapGenerator
}

func (g *game) update(scr *ebiten.Image) error {
	if err := g.mg.GenerateMap(); err != nil {
		return errors.New(fmt.Sprint("error while computing:", err))
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if err := g.mg.Render(scr); err != nil {
		return errors.New(fmt.Sprint("error while rendering:", err))
	}
	return nil
}
