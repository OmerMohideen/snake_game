package entity

import (
	"math/rand"
	"snake_game/utils"
	"snake_game/utils/color"
	"snake_game/world"
	"time"

	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/sdl"
)

type Apple struct {
	X, Y        int
	BoundingBox *resolv.Object
}

func GenerateRandomApple(snake *Snake) *Apple {
	rand.Seed(time.Now().UnixNano())
	var apple Apple

	maxX := (750 - 1) / 32
	maxY := (550 - 1) / 32

	for {
		apple = Apple{
			X: (rand.Intn(maxX) + 1) * 32,
			Y: (rand.Intn(maxY) + 1) * 32,
		}

		if !isSnakeAtPosition(snake, apple) {
			break
		}
	}
	bb := resolv.NewObject(float64(apple.X), float64(apple.Y), 32, 32, "Apple")
	bb.SetShape(resolv.NewRectangle(0, 0, 32, 32))
	world.Space.Add(bb)
	apple.BoundingBox = bb
	return &apple
}

func isSnakeAtPosition(snake *Snake, position Apple) bool {
	current := snake.Head
	for current != nil {
		if current.X == position.X && current.Y == position.Y {
			return true
		}
		current = current.Next
	}
	return false
}

func (apple *Apple) Update(render *sdl.Renderer) {
	utils.Draw(color.Green, int32(apple.X), int32(apple.Y), render)
}

func (apple *Apple) Debug(render *sdl.Renderer) {
	utils.DrawBoundingBox(color.Purple, apple.BoundingBox.X, apple.BoundingBox.Y, apple.BoundingBox.W, render)
}
