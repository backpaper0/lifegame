package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Field struct {
	cellSize, fieldWidth, fieldHeight int
	field                             [][]int
	nextGeneration                    [][]int
	offcetX, offcetY                  int
}

func NewField(cellSize, fieldWidth, fieldHeight int) *Field {
	field := make([][]int, fieldHeight)
	nextGeneration := make([][]int, fieldHeight)
	for y := 0; y < fieldHeight; y++ {
		field[y] = make([]int, fieldWidth)
		nextGeneration[y] = make([]int, fieldWidth)
	}
	f := &Field{
		cellSize:       cellSize,
		fieldWidth:     fieldWidth,
		fieldHeight:    fieldHeight,
		field:          field,
		nextGeneration: nextGeneration,
		offcetX:        cellSize,
		offcetY:        cellSize,
	}
	return f
}

func (f *Field) surroundingSurvival(x, y int) int {
	alive := 0
	for i := Max(y-1, 0); i < Min(y+2, f.fieldHeight); i++ {
		for j := Max(x-1, 0); j < Min(x+2, f.fieldWidth); j++ {
			if i == y && j == x {
				continue
			}
			if f.field[i][j] == 1 {
				alive++
			}
		}
	}
	return alive
}

func (f *Field) Next() {
	for y := 0; y < f.fieldHeight; y++ {
		for x := 0; x < f.fieldWidth; x++ {
			surroundingSurvival := f.surroundingSurvival(x, y)
			f.nextGeneration[y][x] = 0
			switch f.field[y][x] {
			case 0:
				switch surroundingSurvival {
				case 3:
					f.nextGeneration[y][x] = 1
				}
			case 1:
				switch surroundingSurvival {
				case 2, 3:
					f.nextGeneration[y][x] = 1
				}
			}
		}
	}
	f.field, f.nextGeneration = f.nextGeneration, f.field
}

func (f *Field) Random() {
	for y := 0; y < f.fieldHeight; y++ {
		for x := 0; x < f.fieldWidth; x++ {
			if rand.Intn(100) < 30 {
				f.field[y][x] = 1
			} else {
				f.field[y][x] = 0
			}
		}
	}
}

func (f *Field) Clear() {
	for y := 0; y < f.fieldHeight; y++ {
		for x := 0; x < f.fieldWidth; x++ {
			f.field[y][x] = 0
		}
	}
}

func (f *Field) Toggle(px, py int) {
	for y := 0; y < f.fieldHeight; y++ {
		for x := 0; x < f.fieldWidth; x++ {
			if Contains(px, py, x*f.cellSize+f.offcetX, y*f.cellSize+f.offcetY, f.cellSize, f.cellSize) {
				if f.field[y][x] == 0 {
					f.field[y][x] = 1
				} else {
					f.field[y][x] = 0
				}
			}
		}
	}
}

func (f *Field) Draw(screen *ebiten.Image) {

	//格子を描く
	gridColor := color.Gray{0xC0}
	for y, x0, x1 := 0, float32(f.offcetX), float32(f.cellSize*f.fieldWidth+f.offcetX); y <= f.fieldHeight; y++ {
		y0 := float32(f.cellSize*y + f.offcetY)
		y1 := y0
		vector.StrokeLine(screen, x0, y0, x1, y1, 1, gridColor, false)
	}
	for x, y0, y1 := 0, float32(f.offcetY), float32(f.cellSize*f.fieldHeight+f.offcetY); x <= f.fieldWidth; x++ {
		x0 := float32(f.cellSize*x + f.offcetX)
		x1 := x0
		vector.StrokeLine(screen, x0, y0, x1, y1, 1, gridColor, false)
	}

	//生存セルを描く
	cellColor := color.Gray{0x60}
	for y := 0; y < f.fieldHeight; y++ {
		for x := 0; x < f.fieldWidth; x++ {
			switch f.field[y][x] {
			case 1:
				vector.DrawFilledRect(screen, float32(x*f.cellSize+f.offcetX), float32(y*f.cellSize+f.offcetY), float32(f.cellSize), float32(f.cellSize), cellColor, false)
			}
		}
	}
}
