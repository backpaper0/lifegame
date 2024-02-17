package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func Contains(px, py, x, y, width, height int) bool {
	return x <= px && px < x+width && y <= py && py < y+height
}

func DrawTextCenter(screen *ebiten.Image, s string, face font.Face, x, y int, clr color.Color) {
	bound, _ := font.BoundString(face, s)
	text.Draw(screen, s, face, x-((bound.Max.X-bound.Min.X).Ceil()/2), y+((bound.Max.Y-bound.Min.Y).Ceil()/2), color.White)
}
