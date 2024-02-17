package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	screenWidth, screenHeight         int
	cellSize, fieldWidth, fieldHeight int
	field                             [][]int
	nextGeneration                    [][]int
	timer, interval                   int
	offcetX, offcetY                  int
	scene                             int
	buttons                           []*Button
}

type Button struct {
	x, y, width, height int
	text                string
	onClick             func(g *Game, me *Button)
}

func (button *Button) HandleClick(g *Game, x, y int) {
	if button.x <= x && x <= button.x+button.width && button.y <= y && y <= button.y+button.height {
		button.onClick(g, button)
	}
}

func NewGame(cellSize, fieldWidth, fieldHeight int) *Game {
	field := make([][]int, fieldHeight)
	nextGeneration := make([][]int, fieldHeight)
	for y := 0; y < fieldHeight; y++ {
		field[y] = make([]int, fieldWidth)
		nextGeneration[y] = make([]int, fieldWidth)
	}
	buttons := []*Button{
		{
			text: "START",
			x:    720, y: 10, width: 70, height: 30,
			onClick: startOrStop,
		},
		{
			text: "NEXT GEN",
			x:    720, y: 50, width: 70, height: 30,
			onClick: next,
		},
		{
			text: "RANDOM",
			x:    720, y: 90, width: 70, height: 30,
			onClick: random,
		},
		{
			text: "CLEAR",
			x:    720, y: 130, width: 70, height: 30,
			onClick: clear,
		},
	}
	//0.25秒毎に世代交代する
	interval := ebiten.TPS() / 4
	return &Game{
		screenWidth:    800,
		screenHeight:   600,
		cellSize:       cellSize,
		fieldWidth:     fieldWidth,
		fieldHeight:    fieldHeight,
		field:          field,
		nextGeneration: nextGeneration,
		timer:          0,
		interval:       interval,
		offcetX:        cellSize,
		offcetY:        cellSize,
		scene:          0,
		buttons:        buttons,
	}
}

func startOrStop(g *Game, me *Button) {
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

func next(g *Game, me *Button) {
	g.next()
}

func random(g *Game, me *Button) {
	for y := 0; y < g.fieldHeight; y++ {
		for x := 0; x < g.fieldWidth; x++ {
			if rand.Intn(100) < 30 {
				g.field[y][x] = 1
			} else {
				g.field[y][x] = 0
			}
		}
	}
}

func clear(g *Game, me *Button) {
	for y := 0; y < g.fieldHeight; y++ {
		for x := 0; x < g.fieldWidth; x++ {
			g.field[y][x] = 0
		}
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

func (g *Game) next() {
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

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		px, py := ebiten.CursorPosition()
		for _, button := range g.buttons {
			button.HandleClick(g, px, py)
		}
		for y := 0; y < g.fieldHeight; y++ {
			for x := 0; x < g.fieldWidth; x++ {
				if x*g.cellSize+g.offcetX <= px && px < (x+1)*g.cellSize+g.offcetX && y*g.cellSize+g.offcetY <= py && py < (y+1)*g.cellSize+g.offcetY {
					if g.field[y][x] == 0 {
						g.field[y][x] = 1
					} else {
						g.field[y][x] = 0
					}
				}
			}
		}
	}
	switch g.scene {
	case 0:
	case 1:
		g.timer++
		if g.timer%g.interval == 0 {
			g.next()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, float32(g.screenWidth), float32(g.screenHeight), color.White, false)
	clr := color.Gray{0xC0}
	for y, x0, x1 := 0, float32(g.offcetX), float32(g.cellSize*g.fieldWidth+g.offcetX); y <= g.fieldHeight; y++ {
		y0 := float32(g.cellSize*y + g.offcetY)
		y1 := y0
		vector.StrokeLine(screen, x0, y0, x1, y1, 1, clr, false)
	}
	for x, y0, y1 := 0, float32(g.offcetY), float32(g.cellSize*g.fieldHeight+g.offcetY); x <= g.fieldWidth; x++ {
		x0 := float32(g.cellSize*x + g.offcetX)
		x1 := x0
		vector.StrokeLine(screen, x0, y0, x1, y1, 1, clr, false)
	}
	for y := 0; y < g.fieldHeight; y++ {
		for x := 0; x < g.fieldWidth; x++ {
			switch g.field[y][x] {
			case 1:
				vector.DrawFilledRect(screen, float32(x*g.cellSize+g.offcetX), float32(y*g.cellSize+g.offcetY), float32(g.cellSize), float32(g.cellSize), color.Black, false)
			}
		}
	}

	for _, button := range g.buttons {
		drawButton(screen, button.text, button.x, button.y, button.width, button.height)
	}
}

func drawButton(screen *ebiten.Image, s string, x, y, width, height int) {
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), color.Gray{0x60}, false)
	b, _ := font.BoundString(basicfont.Face7x13, s)
	text.Draw(screen, s, basicfont.Face7x13, x+(width/2)-((b.Max.X-b.Min.X).Ceil()/2), y+(height/2)+((b.Max.Y-b.Min.Y).Ceil()/2), color.White)
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
