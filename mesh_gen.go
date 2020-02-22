package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// MeshGenerator ...
type MeshGenerator struct {
	squareGrid     *SquareGrid
	vertecesEbiten []ebiten.Vertex
	verteces       []Vector2
	triangles      []uint16
}

// NewMeshGenerator ...
func NewMeshGenerator() *MeshGenerator {
	return &MeshGenerator{}
}

var emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)

func init() {
	rand.Seed(time.Now().Unix())
	emptyImage.Fill(color.White)
}

// Update implements the ebiten update
func (mg *MeshGenerator) Update(screen *ebiten.Image) error {

	screen.DrawTriangles(mg.vertecesEbiten, mg.triangles, emptyImage, nil)
	return nil
}

// GenerateMesh ...
func (mg *MeshGenerator) GenerateMesh(atlas [][]int, squareSize, offsetX, offsetY float32) {
	mg.squareGrid = NewSquareGrid(atlas, squareSize, offsetX, offsetY)

	for _, squares := range mg.squareGrid.squares {
		for _, square := range squares {
			mg.TriangulateSquare(*square)
		}
	}

	mg.vertecesEbiten = []ebiten.Vertex{}
	for _, vertex := range mg.verteces {
		mg.vertecesEbiten = append(mg.vertecesEbiten,
			ebiten.Vertex{
				DstX: vertex.x, DstY: vertex.y,
				SrcX: 0, SrcY: 0,
				ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1})
	}
}

// TriangulateSquare ...
func (mg *MeshGenerator) TriangulateSquare(square Square) {
	switch square.config {
	case 0:
	// 1 point
	case 1:
		mg.MeshFromPoints(*square.topLeft.node, *square.topMid, *square.midLeft)
	case 2:
		mg.MeshFromPoints(*square.topMid, *square.topRight.node, *square.midRight)
	case 4:
		mg.MeshFromPoints(*square.midRight, *square.bottomRight.node, *square.bottomMid)
	case 8:
		mg.MeshFromPoints(*square.midLeft, *square.bottomMid, *square.bottomLeft.node)

	// 2 points
	case 3:
		mg.MeshFromPoints(*square.topLeft.node, *square.topRight.node, *square.midRight, *square.midLeft)
	case 6:
		mg.MeshFromPoints(*square.topMid, *square.topRight.node, *square.bottomRight.node, *square.bottomMid)
	case 9:
		mg.MeshFromPoints(*square.topLeft.node, *square.topMid, *square.bottomMid, *square.bottomLeft.node)
	case 12:
		mg.MeshFromPoints(*square.midLeft, *square.midRight, *square.bottomRight.node, *square.bottomLeft.node)
	case 5:
		mg.MeshFromPoints(*square.topLeft.node, *square.topMid, *square.midRight, *square.bottomRight.node, *square.bottomMid, *square.midLeft)
	case 10:
		mg.MeshFromPoints(*square.topMid, *square.topRight.node, *square.midRight, *square.bottomMid, *square.bottomLeft.node, *square.midLeft)

	// 3 points
	case 7:
		mg.MeshFromPoints(*square.topLeft.node, *square.topRight.node, *square.bottomRight.node, *square.bottomMid, *square.midLeft)
	case 11:
		mg.MeshFromPoints(*square.topLeft.node, *square.topRight.node, *square.midRight, *square.bottomMid, *square.bottomLeft.node)
	case 13:
		mg.MeshFromPoints(*square.topLeft.node, *square.topMid, *square.midRight, *square.bottomRight.node, *square.bottomLeft.node)
	case 14:
		mg.MeshFromPoints(*square.topMid, *square.topRight.node, *square.midRight, *square.bottomRight.node, *square.bottomLeft.node, *square.midLeft)

	// 4 points
	case 15:
		mg.MeshFromPoints(*square.topLeft.node, *square.topRight.node, *square.bottomRight.node, *square.bottomLeft.node)
	}
}

// MeshFromPoints ...
func (mg *MeshGenerator) MeshFromPoints(points ...Node) {
	points = mg.AssignVertices(points)

	if len(points) >= 3 {
		mg.triangles = CreateTriangle(mg.triangles, points[0], points[1], points[2])
	}
	if len(points) >= 4 {
		mg.triangles = CreateTriangle(mg.triangles, points[0], points[2], points[3])
	}
	if len(points) >= 5 {
		mg.triangles = CreateTriangle(mg.triangles, points[0], points[3], points[4])
	}
	if len(points) >= 6 {
		mg.triangles = CreateTriangle(mg.triangles, points[0], points[4], points[5])
	}
}

// AssignVertices ...
func (mg *MeshGenerator) AssignVertices(points []Node) []Node {
	for i, point := range points {
		// if unassigned
		if point.IsConsumed() {
			continue
		}
		points[i].vertexIndex = uint16(len(mg.verteces))
		point.Consume()
		mg.verteces = append(mg.verteces, point.position)
	}
	return points
}

// CreateTriangle adds an aggregate indexes of three nodes/vertices
func CreateTriangle(triangles []uint16, a, b, c Node) []uint16 {
	return append(triangles, a.vertexIndex, b.vertexIndex, c.vertexIndex)
}
