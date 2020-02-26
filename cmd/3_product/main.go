package main

import (
	"errors"
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/gocs/marchingsquares"
	"github.com/gocs/viewdrag"
	"github.com/hajimehoshi/ebiten"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	mg := marchingsquares.NewMapGenerator(51, 100, 100)

	emptyImage, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)
	w, h := emptyImage.Size()

	// another dependencies
	v := viewdrag.NewViewWithMesh(
		emptyImage,
		rand.Intn(screenWidth-w),
		rand.Intn(screenHeight-h),
		screenWidth,
		screenHeight,
		ebiten.MouseButtonMiddle,
	)

	g := &game{v, mg, emptyImage}

	if err := ebiten.Run(g.update, screenWidth, screenHeight, 1, "Camera Drag"); err != nil {
		log.Fatal("error while running:", err)
	}
}

type game struct {
	v   *viewdrag.View
	mg  *marchingsquares.MapGenerator
	img *ebiten.Image
}

func (g *game) update(scr *ebiten.Image) error {
	if err := g.mg.GenerateMap(); err != nil {
		return errors.New(fmt.Sprint("error while computing:", err))
	}

	v, i := g.mg.GetTriangles()
	if err := g.v.SetMesh(v, i); err != nil {
		return errors.New(fmt.Sprint("error while SetMesh: ", err))
	}
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
