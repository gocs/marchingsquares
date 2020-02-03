package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/gocs/marchingsquares/internal/static"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// New creates new instance of the game struct
func New() *Game {
	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.
	img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Println("err:", err)
		return nil
	}

	// Initialize the sprites.
	w, h := ebitenImage.Size()
	sprite := &Sprite{
		image: ebitenImage,
		x:     rand.Intn(static.ScreenWidth - w),
		y:     rand.Intn(static.ScreenHeight - h),
	}

	return &Game{
		strokes: map[*Stroke]struct{}{},
		sprite:  sprite,
	}
}

// Game handles the states
type Game struct {
	strokes map[*Stroke]struct{}
	sprite  *Sprite
}

func (g *Game) spriteAt(x, y int) *Sprite {
	// As the sprites are ordered from back to front,
	// search the clicked/touched sprite in reverse order.
	s := g.sprite
	if s.In(x, y) {
		return s
	}
	return nil
}

func (g *Game) updateStroke(stroke *Stroke) {
	stroke.Update()
	if !stroke.IsReleased() {
		return
	}

	s := stroke.DraggingObject().(*Sprite)
	if s == nil {
		return
	}

	s.MoveBy(stroke.PositionDiff())

	// Move the dragged sprite to the front.
	g.sprite = s

	stroke.SetDraggingObject(nil)
}

func (g *Game) getMousePress() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s := NewStroke(&MouseStrokeSource{})
		s.SetDraggingObject(g.spriteAt(s.Position()))
		g.strokes[s] = struct{}{}
	}
}

func (g *Game) getTouchPress() {
	for _, id := range inpututil.JustPressedTouchIDs() {
		s := NewStroke(&TouchStrokeSource{id})
		s.SetDraggingObject(g.spriteAt(s.Position()))
		g.strokes[s] = struct{}{}
	}
}

func (g *Game) setReleaseStrokes() {
	for s := range g.strokes {
		g.updateStroke(s)
		if s.IsReleased() {
			delete(g.strokes, s)
		}
	}
}

func (g *Game) setDropped(screen *ebiten.Image) {
	draggingSprites := map[*Sprite]struct{}{}
	for s := range g.strokes {
		if sprite := s.DraggingObject().(*Sprite); sprite != nil {
			draggingSprites[sprite] = struct{}{}
		}
	}
	if _, ok := draggingSprites[g.sprite]; !ok {
		g.sprite.Draw(screen, 0, 0, 1)
	}
}

func (g *Game) setDragged(screen *ebiten.Image) {
	for s := range g.strokes {
		dx, dy := s.PositionDiff()
		if sprite := s.DraggingObject().(*Sprite); sprite != nil {
			sprite.Draw(screen, dx, dy, 0.5)
		}
	}
}

// Update calculates most of the logic
func (g *Game) Update() error {
	g.getMousePress()
	g.getTouchPress()
	g.setReleaseStrokes()
	return nil
}

// Render renders the state of the game
func (g *Game) Render(screen *ebiten.Image) error {
	g.setDropped(screen)
	// transparent sprites
	g.setDragged(screen)
	return nil
}
