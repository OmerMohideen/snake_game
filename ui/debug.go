package ui

import (
	"fmt"
	"snake_game/entity"
	"snake_game/utils"
	"snake_game/utils/color"

	"github.com/veandco/go-sdl2/sdl"
)

func Debug(snake *entity.Snake, renderer *sdl.Renderer) {
	speedText := fmt.Sprintf("speed: %.2f", float64(snake.Speed)/60)
	textRect := sdl.Rect{X: 5, Y: 5, W: 70, H: 20}
	utils.DrawText(speedText, 18, color.White, textRect, renderer)
	renderer.Present()
}
