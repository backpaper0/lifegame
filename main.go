package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	screenWidth, screenHeight         int
	cellSize, fieldWidth, fieldHeight int
	field                             [][]int
	nextGeneration                    [][]int
	timer, interval                   int
}

func NewGame(cellSize, fieldWidth, fieldHeight int, random bool) *Game {
	field := make([][]int, fieldHeight)
	nextGeneration := make([][]int, fieldHeight)
	for y := 0; y < fieldHeight; y++ {
		field[y] = make([]int, fieldWidth)
		nextGeneration[y] = make([]int, fieldWidth)
		if random {
			for x := 0; x < fieldWidth; x++ {
				if rand.Intn(100) < 30 {
					field[y][x] = 1
				}
			}
		}
	}
	//0.25秒毎に世代交代する
	interval := ebiten.TPS() / 4
	return &Game{
		screenWidth:    cellSize * fieldWidth,
		screenHeight:   cellSize * fieldHeight,
		cellSize:       cellSize,
		fieldWidth:     fieldWidth,
		fieldHeight:    fieldHeight,
		field:          field,
		nextGeneration: nextGeneration,
		timer:          0,
		interval:       interval,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func (g *Game) surroundingSurvival(x, y int) int {
	alive := 0
	for i := max(y-1, 0); i < min(y+2, g.fieldHeight); i++ {
		for j := max(x-1, 0); j < min(x+2, g.fieldWidth); j++ {
			if i == y && j == x {
				continue
			}
			if g.field[i][j] == 1 {
				alive++
			}
		}
	}
	return alive
}

func (g *Game) Update() error {
	g.timer++
	// if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
	if g.timer%g.interval == 0 {
		for y := 0; y < g.fieldHeight; y++ {
			for x := 0; x < g.fieldWidth; x++ {
				surroundingSurvival := g.surroundingSurvival(x, y)
				g.nextGeneration[y][x] = 0
				switch g.field[y][x] {
				case 0:
					switch surroundingSurvival {
					case 3:
						g.nextGeneration[y][x] = 1
					}
				case 1:
					switch surroundingSurvival {
					case 2, 3:
						g.nextGeneration[y][x] = 1
					}
				}
			}
		}
		g.field, g.nextGeneration = g.nextGeneration, g.field
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, float32(g.screenWidth), float32(g.screenHeight), color.White, false)
	for y := 0; y < g.fieldHeight; y++ {
		for x := 0; x < g.fieldWidth; x++ {
			switch g.field[y][x] {
			case 1:
				vector.DrawFilledRect(screen, float32(x*g.cellSize), float32(y*g.cellSize), float32(g.cellSize), float32(g.cellSize), color.Black, false)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}

func main() {
	g := NewGame(10, 40, 30, true)
	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle("Life Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
