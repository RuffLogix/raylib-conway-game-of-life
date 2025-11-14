package game

import (
	"go-raylib/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	board   Board
	isPause bool
	frames  int8
}

func (g *Game) Init() {
	g.board.Init()

	g.isPause = true
}

func (g *Game) Update() {
	if rl.IsKeyPressed(rl.KeySpace) {
		g.isPause = !g.isPause
	}

	if !g.isPause {
		if g.frames == 5 {
			g.board.Update()
			g.frames = 0
		}
		g.frames += 1
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	for i := range config.WINDOW_WIDTH / config.TILE_SIZE {
		for j := range config.WINDOW_HEIGHT / config.TILE_SIZE {
			isMouseHover := rl.CheckCollisionPointRec(
				rl.GetMousePosition(),
				rl.NewRectangle(
					float32(i*config.TILE_SIZE),
					float32(j*config.TILE_SIZE),
					float32(config.TILE_SIZE),
					float32(config.TILE_SIZE),
				),
			)
			if isMouseHover {
				rl.DrawRectangle(
					i*config.TILE_SIZE,
					j*config.TILE_SIZE,
					config.TILE_SIZE,
					config.TILE_SIZE,
					rl.Red,
				)

				if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					g.board.cells[i][j] = 1 - g.board.cells[i][j]
				}
			} else {
				if g.board.cells[i][j] == 1 {
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

	// if g.isPause {
	// rl.DrawRectangle(0, 0, config.WINDOW_WIDTH, config.WINDOW_HEIGHT, rl.NewColor(0, 0, 0, 200))
	// rl.DrawText("Game Paused", config.WINDOW_WIDTH/2-rl.MeasureText("Game Paused", 40)/2, config.WINDOW_HEIGHT/2-40, 40, rl.Black)
	// }

	rl.EndDrawing()
}
