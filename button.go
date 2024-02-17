package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type Button struct {
	x, y, width, height int
	text                string
	onClick             func(me *Button)
}

func NewStartButton(g *Game) *Button {
	b := &Button{
		text: "START",
		x:    720, y: 10, width: 70, height: 30,
	}
	b.onClick = func(me *Button) {
		switch g.scene {
		case 0:
			g.timer = 0
			g.scene = 1
			me.text = "STOP"
		case 1:
			g.scene = 0
			me.text = "START"
		}
	}
	return b
}

func NewNextGenButton(g *Game) *Button {
	b := &Button{
		text: "NEXT GEN",
		x:    720, y: 50, width: 70, height: 30,
	}
	b.onClick = func(me *Button) {
		g.field.Next()
	}
	return b
}

func NewRandomButton(g *Game) *Button {
	b := &Button{
		text: "RANDOM",
		x:    720, y: 90, width: 70, height: 30,
	}
	b.onClick = func(me *Button) {
		g.field.Random()
	}
	return b
}

func NewClearButton(g *Game) *Button {
	b := &Button{
		text: "CLEAR",
		x:    720, y: 130, width: 70, height: 30,
	}
	b.onClick = func(me *Button) {
		g.field.Clear()
	}
	return b
}

func (b *Button) HandleClick(px, py int) {
	if Contains(px, py, b.x, b.y, b.width, b.height) {
		b.onClick(b)
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.x), float32(b.y), float32(b.width), float32(b.height), color.Gray{0x60}, false)
	DrawTextCenter(screen, b.text, basicfont.Face7x13, b.x+(b.width/2), b.y+(b.height/2), color.White)
}
