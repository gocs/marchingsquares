package main

import "testing"

func Test_NewTriangle(t *testing.T) {
	tri := NewTriangle(1,2,3)
	t.Logf("tri: %#+v\n", tri)
}

func Test_At(t *testing.T) {
	tri := NewTriangle(1,2,3)
	at := tri.At(1)
	t.Logf("at: %#+v\n", at)
}

func Test_contains(t *testing.T) {
	tri := NewTriangle(1,2,3)
	containing := tri.contains(2)
	t.Logf("containing: %#+v\n", containing)
}
