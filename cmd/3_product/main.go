package main

import (
	"errors"
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/gocs/marchingsquares"
	"github.com/gocs/viewdrag"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 1080
	screenHeight = 768
)

func main() {
	s := 127 // safest
	mg := marchingsquares.NewMapGenerator(51, s, s, 10)
	if err := mg.GenerateMap(); err != nil {
		log.Fatalln("error while computing:", err)
	}
	vx, ix := mg.GetTriangles()

	emptyImage, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)

	// another dependencies
	v := viewdrag.NewViewWithMesh(
		emptyImage,
		vx, ix,
		0, 0,
		screenWidth,
		screenHeight,
		ebiten.MouseButtonMiddle,
	)

	g := &game{v}

	if err := ebiten.Run(g.update, screenWidth, screenHeight, 1, "Camera Drag"); err != nil {
		log.Fatal("error while running:", err)
	}
}

type game struct {
	v *viewdrag.View
}

func (g *game) update(scr *ebiten.Image) error {
	if err := g.v.Compute(scr); err != nil {
		return errors.New(fmt.Sprint("error while computing:", err))
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if err := g.v.Render(scr); err != nil {
		return errors.New(fmt.Sprint("error while rendering:", err))
	}
	return nil
}
