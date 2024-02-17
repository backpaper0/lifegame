package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/internal/ui"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	screenWidth, screenHeight int
	field                     *Field
	timer, interval           int
	scene                     int
	buttons                   []*Button
	touchIDs                  []ui.TouchID
}

func NewGame(cellSize, fieldWidth, fieldHeight int) *Game {
	field := NewField(cellSize, fieldWidth, fieldHeight)

	//0.25秒毎に世代交代する
	interval := ebiten.TPS() / 4

	g := &Game{
		screenWidth:  800,
		screenHeight: 600,
		field:        field,
		timer:        0,
		interval:     interval,
		scene:        0,
		touchIDs:     make([]ui.TouchID, 0),
	}

	g.buttons = []*Button{
		NewStartButton(g),
		NewNextGenButton(g),
		NewRandomButton(g),
		NewClearButton(g),
		NewFullscreenButton(g),
	}
	return g
}

func (g *Game) Update() error {

	//スマートフォン対応
	g.touchIDs = inpututil.AppendJustPressedTouchIDs(g.touchIDs[:0])
	for _, touchID := range g.touchIDs {
		px, py := ebiten.TouchPosition(touchID)
		for _, button := range g.buttons {
			button.HandleClick(px, py)
		}
		g.field.Toggle(px, py)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		px, py := ebiten.CursorPosition()
		for _, button := range g.buttons {
			button.HandleClick(px, py)
		}
		g.field.Toggle(px, py)
	}
	switch g.scene {
	case 0:
	case 1:
		g.timer++
		if g.timer%g.interval == 0 {
			g.field.Next()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, float32(g.screenWidth), float32(g.screenHeight), color.White, false)

	g.field.Draw(screen)

	for _, button := range g.buttons {
		button.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}

func main() {
	g := NewGame(10, 70, 50)
	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle("Life Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
