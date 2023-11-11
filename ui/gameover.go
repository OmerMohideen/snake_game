package ui

import (
	"fmt"
	"snake_game/entity"
	"snake_game/utils"
	"snake_game/utils/color"

	"github.com/veandco/go-sdl2/sdl"
)

func GameOver(snake *entity.Snake, renderer *sdl.Renderer) {
	textRect := sdl.Rect{X: 320, Y: 260, W: 140, H: 40}
	utils.DrawText("Game Over", 32, color.White, textRect, renderer)

	text := fmt.Sprintf("Score: %v", ((155-snake.Speed)*uint32(snake.Length))-((155-snake.Speed)*3))	
	textRect = sdl.Rect{X: 320, Y: 300, W: 140, H: 30}
	utils.DrawText(text, 32, color.White, textRect, renderer)
	renderer.Present()
}
