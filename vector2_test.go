package main

import (
	"testing"
)

var (
	v1 = Vector2{x: 8, y: 6}
	v2 = Vector2{x: 2, y: 3}
)

func Test_Add(t *testing.T) {
	t.Log("v1 + v2:", Vector2{x: v1.x + v2.x, y: v1.y + v2.y})
}