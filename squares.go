package marchingsquares

import "github.com/hajimehoshi/ebiten"

// Point contains x and y axis
type Point struct {
	X, Y float32
}

// triangle produces regular triangle. note that you will still depend on this []uint16{0, 1, 2, 1, 0, 2} indexes when DrawTriangles-ing
func triangle(p1, p2, p3 Point) []ebiten.Vertex {
	return []ebiten.Vertex{
		ebiten.Vertex{DstX: p1.X, DstY: p1.Y, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		ebiten.Vertex{DstX: p2.X, DstY: p2.Y, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		ebiten.Vertex{DstX: p3.X, DstY: p3.Y, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	}
}

// GenerateVertices generates map of vertices
func GenerateSquares(fieldmap [][]int, x, y, w, h float32) (vertices [][]ebiten.Vertex, triIndex []uint16) {
	for i, horizontal := range fieldmap {
		for j, vertical := range horizontal {
			vertices = append(vertices, marchingSquare(vertical, x+float32(j)*w, y+float32(i)*h, w, h)...)
		}
	}
	return vertices, []uint16{0, 1, 2, 1, 0, 2}
}

// marchingSquare produces a square piece. note that you will still depend on this []uint16{0, 1, 2, 1, 0, 2} indexes when DrawTriangles-ing
func marchingSquare(layout int, x, y, w, h float32) (layouts [][]ebiten.Vertex) {
	toplef := Point{X: x, Y: y}
	topmid := Point{X: x + w/2, Y: y}
	toprgt := Point{X: x + w, Y: y}
	ctrlef := Point{X: x, Y: y + h/2}
	ctrrgt := Point{X: x + w, Y: y + h/2}
	botlef := Point{X: x, Y: y + h}
	botmid := Point{X: x + w/2, Y: y + h}
	botrgt := Point{X: x + w, Y: y + h}

	switch layout {
	case 0:
	case 1:
		layouts = append(layouts, triangle(toplef, topmid, ctrlef))
	case 2:
		layouts = append(layouts, triangle(topmid, toprgt, ctrrgt))
	case 3:
		layouts = append(layouts, triangle(toplef, toprgt, ctrrgt))
		layouts = append(layouts, triangle(toplef, ctrrgt, ctrlef))
	case 4:
		layouts = append(layouts, triangle(ctrrgt, botrgt, botmid))
	case -5:
		layouts = append(layouts, triangle(toplef, topmid, ctrlef))
		layouts = append(layouts, triangle(ctrrgt, botrgt, botmid))
	case 5:
		layouts = append(layouts, triangle(topmid, toprgt, ctrrgt))
		layouts = append(layouts, triangle(topmid, ctrrgt, botmid))
		layouts = append(layouts, triangle(topmid, botmid, botlef))
		layouts = append(layouts, triangle(topmid, botlef, ctrlef))
	case 6:
		layouts = append(layouts, triangle(topmid, toprgt, botrgt))
		layouts = append(layouts, triangle(topmid, botrgt, botmid))
	case 7:
		layouts = append(layouts, triangle(toplef, toprgt, ctrrgt))
		layouts = append(layouts, triangle(toplef, ctrrgt, botrgt))
		layouts = append(layouts, triangle(toplef, botrgt, botmid))
		layouts = append(layouts, triangle(toplef, botmid, ctrlef))
	case 8:
		layouts = append(layouts, triangle(ctrlef, botmid, botlef))
	case 9:
		layouts = append(layouts, triangle(toplef, topmid, botmid))
		layouts = append(layouts, triangle(toplef, botmid, botlef))
	case -10:
		layouts = append(layouts, triangle(topmid, toprgt, ctrrgt))
		layouts = append(layouts, triangle(ctrlef, botmid, botlef))
	case 10:
		layouts = append(layouts, triangle(toplef, topmid, ctrrgt))
		layouts = append(layouts, triangle(toplef, ctrrgt, botrgt))
		layouts = append(layouts, triangle(toplef, botrgt, botmid))
		layouts = append(layouts, triangle(toplef, botmid, ctrlef))
	case 11:
		layouts = append(layouts, triangle(toplef, toprgt, ctrrgt))
		layouts = append(layouts, triangle(toplef, ctrrgt, botmid))
		layouts = append(layouts, triangle(toplef, botmid, botlef))
	case 12:
		layouts = append(layouts, triangle(ctrlef, ctrrgt, botrgt))
		layouts = append(layouts, triangle(ctrlef, botrgt, botlef))
	case 13:
		layouts = append(layouts, triangle(toplef, topmid, ctrrgt))
		layouts = append(layouts, triangle(toplef, ctrrgt, botrgt))
		layouts = append(layouts, triangle(toplef, botrgt, botlef))
	case 14:
		layouts = append(layouts, triangle(topmid, toprgt, botrgt))
		layouts = append(layouts, triangle(topmid, botrgt, botlef))
		layouts = append(layouts, triangle(topmid, botlef, ctrlef))
	case 15:
		layouts = append(layouts, triangle(toplef, toprgt, botrgt))
		layouts = append(layouts, triangle(toplef, botrgt, botlef))
	}
	return
}
