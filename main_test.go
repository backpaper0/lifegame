package main

import (
	"testing"
)

func TestSurroundingSurvival(t *testing.T) {
	t.Run("左上", func(t *testing.T) {
		g := NewGame(10, 3, 3)
		g.field[0][1] = 1
		g.field[1][0] = 1
		g.field[1][1] = 1
		expected := 3
		if actual := g.surroundingSurvival(0, 0); actual != expected {
			t.Fatalf("Expected is %v but actual is %v", expected, actual)
		}
	})
	t.Run("右上", func(t *testing.T) {
		g := NewGame(10, 3, 3)
		g.field[0][1] = 1
		g.field[1][2] = 1
		// g.field[1][1] = 1
		expected := 2
		if actual := g.surroundingSurvival(2, 0); actual != expected {
			t.Fatalf("Expected is %v but actual is %v", expected, actual)
		}
	})
	t.Run("左下", func(t *testing.T) {
		g := NewGame(10, 3, 3)
		g.field[1][0] = 1
		g.field[1][1] = 1
		g.field[2][1] = 1
		expected := 3
		if actual := g.surroundingSurvival(0, 2); actual != expected {
			t.Fatalf("Expected is %v but actual is %v", expected, actual)
		}
	})
	t.Run("右下", func(t *testing.T) {
		g := NewGame(10, 3, 3)
		g.field[1][2] = 1
		// g.field[1][1] = 1
		g.field[2][1] = 1
		expected := 2
		if actual := g.surroundingSurvival(2, 2); actual != expected {
			t.Fatalf("Expected is %v but actual is %v", expected, actual)
		}
	})
	t.Run("中央1", func(t *testing.T) {
		g := NewGame(10, 3, 3)
		g.field[0][0] = 1
		g.field[0][1] = 1
		g.field[0][2] = 1
		g.field[1][0] = 1
		g.field[1][2] = 1
		g.field[2][0] = 1
		g.field[2][1] = 1
		g.field[2][2] = 1
		expected := 8
		if actual := g.surroundingSurvival(1, 1); actual != expected {
			t.Fatalf("Expected is %v but actual is %v", expected, actual)
		}
	})
	t.Run("中央2", func(t *testing.T) {
		g := NewGame(10, 3, 3)
		expected := 0
		if actual := g.surroundingSurvival(1, 1); actual != expected {
			t.Fatalf("Expected is %v but actual is %v", expected, actual)
		}
	})
}
