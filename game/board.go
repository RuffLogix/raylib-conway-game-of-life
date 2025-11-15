package game

import (
	"fmt"
	"go-raylib/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var xDirection = []int8{-1, 0, 1}
var yDirection = []int8{-1, 0, 1}

type Board struct {
	cols  int32
	rows  int32
	cells [][]int8

	population int32
}

func (b *Board) Init() {
	b.rows = config.WINDOW_WIDTH / config.TILE_SIZE
	b.cols = config.WINDOW_HEIGHT / config.TILE_SIZE

	cells := make([][]int8, b.rows)
	for i := range cells {
		cells[i] = make([]int8, b.cols)
	}

	b.cells = cells
}

func (b *Board) Update() {
	b.population = 0

	tempCells := make([][]int8, b.rows)
	for i := range tempCells {
		tempCells[i] = make([]int8, b.cols)
		copy(tempCells[i], b.cells[i])
	}

	for i := range b.rows {
		for j := range b.cols {
			b.cells[i][j] = b.isAlive(i, j, &tempCells)
			b.population += int32(b.cells[i][j])
		}
	}
}

func (b *Board) isCellValid(i int32, j int32) bool {
	if i < 0 || j < 0 ||
		i >= b.rows ||
		j >= b.cols {
		return false
	}

	return true
}

func (b *Board) countNeighbors(i int32, j int32, cells *[][]int8) int8 {
	neighbors := int8(0)
	for xIndex := range xDirection {
		for yIndex := range yDirection {
			if b.isCellValid(
				i+int32(xDirection[xIndex]),
				j+int32(yDirection[yIndex]),
			) && (xIndex != 1 || yIndex != 1) {
				neighbors += (*cells)[i+int32(xDirection[xIndex])][j+int32(yDirection[yIndex])]
			}
		}
	}

	return neighbors
}

func (b *Board) isAlive(i int32, j int32, cells *[][]int8) int8 {
	neightbors := b.countNeighbors(i, j, cells)

	if neightbors == 3 && (*cells)[i][j] == 0 {
		return 1
	}
	if neightbors < 2 || neightbors > 3 || (*cells)[i][j] == 0 {
		return 0
	}

	return 1
}

func (b *Board) Draw() {
	for i := range b.rows {
		for j := range b.cols {
			mouseCol := rl.GetMouseX() / config.TILE_SIZE
			mouseRow := rl.GetMouseY() / config.TILE_SIZE

			if i == mouseCol && j == mouseRow {
				rl.DrawRectangle(
					i*config.TILE_SIZE,
					j*config.TILE_SIZE,
					config.TILE_SIZE,
					config.TILE_SIZE,
					rl.Red,
				)

				if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					if b.cells[i][j] == 1 {
						b.population -= 1
					} else {
						b.population += 1
					}
					b.cells[i][j] = 1 - b.cells[i][j]
				}
			} else {
				if b.cells[i][j] == 1 {
					rl.DrawRectangle(
						i*config.TILE_SIZE,
						j*config.TILE_SIZE,
						config.TILE_SIZE,
						config.TILE_SIZE,
						rl.Blue,
					)
				}
			}
			rl.DrawRectangleLines(
				i*config.TILE_SIZE,
				j*config.TILE_SIZE,
				config.TILE_SIZE,
				config.TILE_SIZE,
				rl.LightGray,
			)
		}
	}

	populationText := fmt.Sprintf("Population: %d", b.population)

	rl.DrawRectangle(25, 25, rl.MeasureText(populationText, 14)+10, 20, rl.NewColor(0, 0, 0, 100))
	rl.DrawText(populationText, 30, 30, 14, rl.White)
}
