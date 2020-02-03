package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gocs/marchingsquares/internal/static"
	"github.com/hajimehoshi/ebiten"
)

var game *Game

func init() {
	rand.Seed(time.Now().UnixNano())
	game = New()
}

func main() {
	if err := ebiten.Run(update, static.ScreenWidth, static.ScreenHeight, 1, "Marching Squares"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	if err := game.Update(); err != nil {
		return err
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return game.Render(screen)
}
