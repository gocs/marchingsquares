package marchingsquares

// Triangle ...
type Triangle struct {
	vertexIndexA, vertexIndexB, vertexIndexC int
	verteces                                 []int
}

// NewTriangle generates new Triangle
func NewTriangle(a, b, c int) *Triangle {
	verteces := []int{}
	verteces = append(verteces, a, b, c)
	return &Triangle{
		vertexIndexA: a,
		vertexIndexB: b,
		vertexIndexC: c,
		verteces:     verteces,
	}
}

// At accepts vertex index returns vertex number
func (t *Triangle) At(i int) int {
	return t.verteces[i]
}

func (t *Triangle) contains(vertexIndex int) bool {
	return vertexIndex == t.vertexIndexA || vertexIndex == t.vertexIndexB || vertexIndex == t.vertexIndexC
}