package main

import (
	"snake_game/entity"
	"snake_game/math"
	"snake_game/ui"
	"snake_game/world"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Snake Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	world.Init()

	snake := entity.SpawnSnake()
	snake.LastApple = entity.GenerateRandomApple(snake)

	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch eventType := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if eventType.State == sdl.PRESSED {
					switch eventType.Keysym.Scancode {
					case sdl.SCANCODE_LEFT, sdl.SCANCODE_A:
						if snake.LastDirection != math.Left && snake.LastDirection != math.Right && !world.GameOver {
							snake.Move(math.Left)
						}
					case sdl.SCANCODE_RIGHT, sdl.SCANCODE_D:
						if snake.LastDirection != math.Right && snake.LastDirection != math.Left && !world.GameOver {
							snake.Move(math.Right)
						}
					case sdl.SCANCODE_UP, sdl.SCANCODE_W:
						if snake.LastDirection != math.Up && snake.LastDirection != math.Down && !world.GameOver {

							snake.Move(math.Up)
						}
					case sdl.SCANCODE_DOWN, sdl.SCANCODE_S:
						if snake.LastDirection != math.Down && snake.LastDirection != math.Up && !world.GameOver {
							snake.Move(math.Down)
						}
					case sdl.SCANCODE_TAB:
						world.Debug = !world.Debug
					case sdl.SCANCODE_PERIOD:
						snake.Speed += 0.1 * 60
					case sdl.SCANCODE_COMMA:
						snake.Speed -= 0.1 * 60
					case sdl.SCANCODE_P:
						snake.LastDirection = math.Idle
					}
				}
			}
		}

		snake.Update(renderer)
		snake.LastApple.Update(renderer)

		if world.Debug {
			snake.Debug(renderer)
			snake.LastApple.Debug(renderer)
			ui.Debug(snake, renderer)
		}

		if world.GameOver {
			snake.LastDirection = math.Idle
			ui.GameOver(snake, renderer)
		}

		renderer.Present()
		sdl.Delay(snake.Speed)
	}
}
