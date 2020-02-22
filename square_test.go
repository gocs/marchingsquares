package main

import "testing"

func Test_NewSquare(t *testing.T) {
	tl := NewControlNode(Vector2{0, 0}, true, 10)
	tr := NewControlNode(Vector2{10, 0}, true, 10)
	bl := NewControlNode(Vector2{0, 10}, false, 10)
	br := NewControlNode(Vector2{10, 10}, true, 10)
	s := NewSquare(tl, tr, br, bl)
	t.Logf("config     : %v", s.config)
}

func Test_NewOneSquareGrid(t *testing.T) {
	// 1
	atlas := [][]int{
		{1, 0},
		{0, 0}}
	s := NewSquareGrid(atlas, 100, 100, 100).squares[0][0]
	t.Logf("config     : %v", s.config)
}

func Test_NewSquareGrid(t *testing.T) {
	// 0
	atlas := [][]int{
		{0, 0},
		{0, 0}}
	s := NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// 1
	atlas = [][]int{
		{1, 0},
		{0, 0}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// 2
	atlas = [][]int{
		{0, 0},
		{1, 0}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// 3
	atlas = [][]int{
		{1, 0},
		{1, 0}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// 4
	atlas = [][]int{
		{0, 1},
		{0, 0}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// 5
	atlas = [][]int{
		{1, 1},
		{0, 0}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// 6
	atlas = [][]int{
		{0, 1},
		{1, 0}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// 7
	atlas = [][]int{
		{1, 1},
		{1, 0}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	//8
	atlas = [][]int{
		{0, 0},
		{0, 1}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	//9
	atlas = [][]int{
		{1, 0},
		{0, 1}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// 10
	atlas = [][]int{
		{0, 0},
		{1, 1}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// 11
	atlas = [][]int{
		{1, 0},
		{1, 1}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	//12
	atlas = [][]int{
		{0, 1},
		{0, 1}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	//13
	atlas = [][]int{
		{1, 1},
		{0, 1}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// 14
	atlas = [][]int{
		{0, 1},
		{1, 1}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// 15
	atlas = [][]int{
		{1, 1},
		{1, 1}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

	// [10,7]
	atlas = [][]int{
		{0, 0},
		{1, 1},
		{1, 0}}
	s = NewSquareGrid(atlas, 10, 10, 10).squares[0][0]
	t.Logf("config     : %v", s.config)

}
