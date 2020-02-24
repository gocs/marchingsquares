package main

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// MapGenerator contains map generator state
type MapGenerator struct {
	emptyImage    *ebiten.Image
	width, height int
	seed          string
	useRandomSeed bool

	randomFillPercent int
	atlas             [][]int

	mg *MeshGenerator
}

// NewMapGenerator instantiates a map
func NewMapGenerator(emptyImage *ebiten.Image, randomFillPercent, width, height int) *MapGenerator {
	return &MapGenerator{emptyImage: emptyImage, randomFillPercent: randomFillPercent, width: width, height: height}
}

// Update implements ebiten run
func (mg *MapGenerator) Update(screen *ebiten.Image) error {
	if err := mg.GenerateMap(); err != nil {
		return err
	}

	if err := mg.OnDraw(screen); err != nil {
		return err
	}

	return nil
}

// OnDraw ...
func (mg *MapGenerator) OnDraw(screen *ebiten.Image) error {
	if len(mg.atlas) < 1 {
		return errors.New("ondraw mg atlas empty")
	}

	mg.mg.Update(screen)
	return nil
}

// GenerateMap generates map by filling, smoothing and generating mesh
func (mg *MapGenerator) GenerateMap() error {
	if len(mg.atlas) > 0 {
		return nil
	}
	mg.atlas = RandomFillMap(mg.width, mg.height, mg.randomFillPercent)
	for i := 0; i < 4; i++ {
		mg.atlas = SmoothMap(mg.atlas, mg.width, mg.height)
	}
	mg.atlas = InvertMap(mg.atlas, mg.width, mg.height)
	log.Println("atlas", mg.atlas)

	mg.mg = NewMeshGenerator()
	mg.mg.GenerateMesh(mg.atlas, 10, 200, 100)
	return nil
}

// RandomFillMap fills map with random values if using random seed
func RandomFillMap(w, h, percent int) (atlas [][]int) {

	atlas = make([][]int, w)
	for x := 0; x < w; x++ {
		atlas[x] = make([]int, h)
		for y := 0; y < h; y++ {
			if rand.Intn(100) < percent {
				atlas[x][y] = 1
			}
		}
	}
	return
}

// SmoothMap smoothes map to look like an actual map
func SmoothMap(atlas [][]int, width, height int) [][]int {
	neighborWallTiles := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			neighborWallTiles = GetSurroundingWallCount(atlas, width, height, x, y)

			if neighborWallTiles > 4 {
				atlas[x][y] = 1
			} else if neighborWallTiles < 4 {
				atlas[x][y] = 0
			}
		}
	}
	return atlas
}

// GetSurroundingWallCount ...
func GetSurroundingWallCount(atlas [][]int, width, height int, gridX, gridY int) int {
	wallCount := 0
	for neighborX := gridX - 1; neighborX <= gridX+1; neighborX++ {
		for neighborY := gridY - 1; neighborY <= gridY+1; neighborY++ {
			if neighborX >= 0 && neighborX < width && neighborY >= 0 && neighborY < height {
				if neighborX != gridX || neighborY != gridY {
					wallCount += atlas[neighborX][neighborY]
				}
			} else {
				wallCount++
			}
		}
	}
	return wallCount
}

// InvertMap inverts map
func InvertMap(atlas [][]int, width, height int) [][]int {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if atlas[x][y] == 0 {
				atlas[x][y] = 1
			} else {
				atlas[x][y] = 0
			}
		}
	}
	return atlas
}
