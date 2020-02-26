package marchingsquares

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
	t.Logf("atlas: %#+v\n", mg.atlas)
}

func Test_SmoothMap(t *testing.T) {
	var emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)
	mg := NewMapGenerator(emptyImage, 51, 8, 6)

	mg.atlas = RandomFillMap(mg.width, mg.height, mg.randomFillPercent)
	for i := 0; i < 4; i++ {
		mg.atlas = SmoothMap(mg.atlas, mg.width, mg.height)
	}
	t.Logf("atlas: %#+v\n", mg.atlas)
}

func Test_GenerateMap(t *testing.T) {
	var emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)
	mg := NewMapGenerator(emptyImage, 51, 20, 10)
	mg.GenerateMap()
	if len(mg.atlas) < 1 {
		t.Error("empty squares")
	}
	// call second time let guard pass
	mg.GenerateMap()
	for _, v := range mg.atlas {
		for _, square := range v {
			t.Logf("square: %#+v\n", square)
		}
	}
}
