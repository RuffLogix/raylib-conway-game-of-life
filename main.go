package main

import (
	"go-raylib/config"
	"go-raylib/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	game := game.Game{}
	game.Init()

	rl.InitWindow(config.WINDOW_WIDTH, config.WINDOW_HEIGHT, config.WINDOW_TITLE)
	defer rl.CloseWindow()

	rl.SetTargetFPS(config.TARGET_FPS)

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}
}
