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

	isNextPressed := g.isPause && rl.IsKeyPressed(rl.KeyR)

	if !g.isPause || isNextPressed {
		if g.frames == 5 || isNextPressed {
			g.board.Update()
			g.frames = 0
		}
		g.frames += 1
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	g.board.Draw()

	if g.isPause {
		rl.DrawText(
			config.PAUSE_MESSAGE,
			config.WINDOW_WIDTH-rl.MeasureText(config.PAUSE_MESSAGE, 20)-25,
			25,
			20,
			rl.Red,
		)
	}

	rl.DrawRectangle(
		25,
		config.WINDOW_HEIGHT-60,
		rl.MeasureText(config.CONTROL_KEYS_MESSAGE, 14)+10, 50,
		rl.NewColor(0, 0, 0, 100),
	)
	rl.DrawText(config.CONTROL_KEYS_MESSAGE, 30, config.WINDOW_HEIGHT-55, 14, rl.White)

	rl.EndDrawing()
}
