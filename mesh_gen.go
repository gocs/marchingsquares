package marchingsquares

import "github.com/hajimehoshi/ebiten"

// MeshGenerator provides struct for DrawTriangle based on the atlas from map generator
type MeshGenerator struct {
	squareGrid     *SquareGrid
	vertecesEbiten []ebiten.Vertex
	verteces       []Vector2
	triangles      []uint16
}

// GetTriangles returns the verteces and the indices after generating the mesh
func (mg *MeshGenerator) GetTriangles() ([]ebiten.Vertex, []uint16) {
	return mg.vertecesEbiten, mg.triangles
}

// GenerateMesh generates mesh based on atlas,  squareSize, and displacement
func (mg *MeshGenerator) GenerateMesh(atlas [][]int, squareSize float32) {
	mg.squareGrid = NewSquareGrid(atlas, squareSize)

	for _, squares := range mg.squareGrid.squares {
		for _, square := range squares {
			mg.TriangulateSquare(*square)
		}
	}

	// offset to zero
	var minX, minY float32
	for _, vertex := range mg.verteces {
		if vertex.x < minX {
			minX = vertex.x
		}
		if vertex.y < minY {
			minY = vertex.y
		}
	}

	for _, vertex := range mg.verteces {
		mg.vertecesEbiten = append(mg.vertecesEbiten,
			ebiten.Vertex{
				DstX: vertex.x -minX, DstY: vertex.y -minY,
				SrcX: 0, SrcY: 0,
				ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1})
	}
}

// TriangulateSquare translate square configs to ebiten verteces and indeces
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

// MeshFromPoints sets triangle points to meshes
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

// AssignVertices connects triangles' verteces to another point
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
