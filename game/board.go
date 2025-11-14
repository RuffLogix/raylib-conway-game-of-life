package game

import "go-raylib/config"

var xDirection = []int8{-1, 0, 1}
var yDirection = []int8{-1, 0, 1}

type Board struct {
	cells [][]int8
}

func (b *Board) Init() {
	cells := make([][]int8, config.WINDOW_WIDTH/config.TILE_SIZE)
	for i := range cells {
		cells[i] = make([]int8, config.WINDOW_HEIGHT/config.TILE_SIZE)
	}

	b.cells = cells
}

func (b *Board) Update() {
	tempCells := make([][]int8, config.WINDOW_WIDTH/config.TILE_SIZE)
	for i := range tempCells {
		tempCells[i] = make([]int8, config.WINDOW_HEIGHT/config.TILE_SIZE)
		copy(tempCells[i], b.cells[i])
	}

	for i := range config.WINDOW_WIDTH / config.TILE_SIZE {
		for j := range config.WINDOW_HEIGHT / config.TILE_SIZE {
			b.cells[i][j] = b.isAlive(i, j, &tempCells)
		}
	}
}

func (b *Board) isCellValid(i int32, j int32) bool {
	if i < 0 || j < 0 ||
		i >= config.WINDOW_WIDTH/config.TILE_SIZE ||
		j >= config.WINDOW_HEIGHT/config.TILE_SIZE {
		return false
	}

	return true
}

func (b *Board) countNeighbors(i int32, j int32, cells *[][]int8) int8 {
	neighbors := int8(0)
	for xIndex := range 3 {
		for yIndex := range 3 {
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
