package main

import (
	"testing"
)

func Test_NewMeshGenerator(t *testing.T) {
	mg := NewMeshGenerator()
	t.Logf("mg: %#+v\n", mg)
}

func Test_GenerateMesh(t *testing.T) {
	mg := NewMeshGenerator()
	atlas := [][]int{
		{1, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
	}
	mg.GenerateMesh(atlas, 10, 10, 10)
	for _, squares := range mg.squareGrid.squares {
		for _, square := range squares {
			t.Logf("mg square pos: %#+v\n", square.topLeft.node.position)
			t.Logf("mg square pos: %#+v\n", square.topRight.node.position)
			t.Logf("mg square pos: %#+v\n", square.bottomLeft.node.position)
			t.Logf("mg square pos: %#+v\n", square.bottomRight.node.position)
		}
	}
	t.Logf("mg triangles: %#+v\n", mg.triangles)
	t.Logf("mg verteces: %#+v\n", mg.verteces)
}

func Test_TriangulateSquare(t *testing.T) {
	mg := NewMeshGenerator()

	tl := NewControlNode(Vector2{1, 1}, true, 10)
	tr := NewControlNode(Vector2{2, 1}, false, 10)
	bl := NewControlNode(Vector2{1, 2}, false, 10)
	br := NewControlNode(Vector2{2, 2}, false, 10)
	s := NewSquare(tl, tr, br, bl)
	mg.TriangulateSquare(*s)

	t.Logf("mg triangles: %#+v\n", mg.triangles)
	t.Logf("mg verteces: %#+v\n", mg.verteces)
}

func Test_AssignVertices(t *testing.T) {
	nodes := append([]Node{},
		*NewNode(Vector2{1, 1}),
		*NewNode(Vector2{1, 2}),
		*NewNode(Vector2{2, 1}))
	mg := NewMeshGenerator()
	points := mg.AssignVertices(nodes)
	for _, v := range mg.verteces {
		t.Logf("v: %#+v\n", v)
	}
	for _, p := range points {
		t.Logf("vertexIndex: %#+v\n", p.vertexIndex)
	}
	points = mg.AssignVertices(nodes)
	for _, v := range mg.verteces {
		t.Logf("v: %#+v\n", v)
	}
	for _, p := range points {
		t.Logf("vertexIndex: %#+v\n", p.vertexIndex)
	}
}

func Test_CreateTriangle(t *testing.T) {
	node := *NewNode(Vector2{1, 2})
	node.vertexIndex = 1
	mg := NewMeshGenerator()
	mg.triangles = CreateTriangle(mg.triangles,
		node,
		*NewNode(Vector2{1, 2}),
		*NewNode(Vector2{1, 2}))
	t.Logf("mg: %#+v\n", mg.triangles)
}
