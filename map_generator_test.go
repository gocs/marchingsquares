package main

import (
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten"
)

/**
 * Guys i'm new to testing
 */

func Test_NewMapGenerator(t *testing.T) {
	var emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)
	mg := NewMapGenerator(emptyImage, 51, 128, 96)
	t.Logf("mg: %#+v\n", mg)
}

// // / not working aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
// func Test_Update(t *testing.T) {
// 	// var emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
// 	// emptyImage.Fill(color.White)
// 	// mg := NewMapGenerator(emptyImage, 51, 128, 96)
// 	// if err := ebiten.Run(mg.Update, 640, 480, 1, "title"); err != nil {
// 	if err := ebiten.Run(func (screen *ebiten.Image) error {
// 		return nil
// 	}, 640, 480, 1, "title"); err != nil {
// 		t.Error("run exited:", err)
// 	}
// }

// // /// not working
// func Test_OnDraw(t *testing.T) {
// }

func Test_RandomFillMap(t *testing.T) {
	var emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)
	mg := NewMapGenerator(emptyImage, 51, 8, 6)

	mg.atlas = RandomFillMap(mg.width, mg.height, mg.randomFillPercent)
	t.Logf("mg: %#+v\n", mg)
}

func Test_GenerateMap(t *testing.T) {
	var emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)
	mg := NewMapGenerator(emptyImage, 51, 20, 10)
	mg.GenerateMap()
	for _, squares := range mg.sq.squares {
		for _, square := range squares {
			t.Logf("square: %#+v\n", square)
		}
	}
}
