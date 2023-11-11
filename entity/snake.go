package entity

import (
	"snake_game/math"
	"snake_game/utils"
	"snake_game/utils/color"
	"snake_game/world"

	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/sdl"
)

type SnakeSegment struct {
	X, Y        int
	Next        *SnakeSegment
	BoundingBox *resolv.Object
}

type Snake struct {
	Head, Tail    *SnakeSegment
	LastDirection math.Direction
	LastApple     *Apple
	Speed         uint32
	Length        int
}

func SpawnSnake() *Snake {
	head := &SnakeSegment{X: 96, Y: 5}
	head.BoundingBox = resolv.NewObject(float64(head.X), float64(head.Y), 32, 32, "Head")
	head.BoundingBox.SetShape(resolv.NewRectangle(0, 0, 32, 32))
	world.Space.Add(head.BoundingBox)
	tail := head
	for i := 1; i < 3; i++ {
		tail.Next = &SnakeSegment{X: 96 - 32*i, Y: 5}
		tail.Next.BoundingBox = resolv.NewObject(float64(tail.Next.X), float64(tail.Next.Y), 32, 32, "Tail")
		tail.Next.BoundingBox.SetShape(resolv.NewRectangle(0, 0, 32, 32))
		world.Space.Add(tail.Next.BoundingBox)
		tail = tail.Next
	}
	return &Snake{Head: head, Tail: tail, Speed: uint32(150), Length: 2}
}

func (snake *Snake) Move(direction math.Direction) {
	newHead := &SnakeSegment{
		X: snake.Head.X + direction.X,
		Y: snake.Head.Y + direction.Y,
	}
	newHead.BoundingBox = resolv.NewObject(float64(newHead.X), float64(newHead.Y), 32, 32, "Head")
	newHead.BoundingBox.SetShape(resolv.NewRectangle(0, 0, 32, 32))
	world.Space.Add(newHead.BoundingBox)

	teleport(newHead)

	newHead.Next = snake.Head
	snake.Head = newHead

	snake.Head.BoundingBox.X = float64(snake.Head.Next.X)
	snake.Head.BoundingBox.Y = float64(snake.Head.Next.Y)
	snake.Head.BoundingBox.Update()

	if snake.collusionCheckBody() && snake.LastDirection != math.Idle {
		world.GameOver = true
	}

	snake.LastDirection = direction

	if snake.collusionCheckApple() {
		snake.Head = newHead
		world.Space.Remove(snake.LastApple.BoundingBox)
		snake.LastApple = GenerateRandomApple(snake)
		snake.Length++

		if snake.Length > 6 && snake.Length%3 == 0 && snake.Speed > 60 {
			snake.Speed -= 0.1 * 60
		}
	} else {
		current := snake.Head
		for current.Next != snake.Tail {
			current = current.Next
		}
		current.Next = nil
		snake.Tail = current
	}
}

func teleport(head *SnakeSegment) {
	if head.X > 800 {
		head.X = 1
	}

	if head.X < 0 {
		head.X = 799
	}

	if head.Y > 600 {
		head.Y = 1
	}

	if head.Y < 0 {
		head.Y = 599
	}
}

func (snake *Snake) collusionCheckApple() bool {
	collision := snake.Head.BoundingBox.Check(float64(snake.Head.BoundingBox.X)-float64(snake.LastApple.BoundingBox.X), float64(snake.Head.BoundingBox.Y)-float64(snake.LastApple.BoundingBox.Y), "Apple")
	return collision != nil
}

func (snake *Snake) collusionCheckBody() bool {
	currentSegment := snake.Head.Next

	headX := snake.Head.BoundingBox.X
	headY := snake.Head.BoundingBox.Y
	headW := snake.Head.BoundingBox.W
	headH := snake.Head.BoundingBox.H

	for currentSegment != nil {
		segmentX := currentSegment.BoundingBox.X
		segmentY := currentSegment.BoundingBox.Y
		segmentW := currentSegment.BoundingBox.W
		segmentH := currentSegment.BoundingBox.H

		if headX < segmentX+segmentW && headX+headW > segmentX &&
			headY < segmentY+segmentH && headY+headH > segmentY {
			return true 
		}

		currentSegment = currentSegment.Next
	}

	return false
}

func (snake *Snake) Update(render *sdl.Renderer) {
	render.SetDrawColor(0, 0, 0, 0)
	render.Clear()

	current := snake.Head
	for current != nil {
		utils.Draw(color.Red, int32(current.X), int32(current.Y), render)
		current = current.Next
	}

	if snake.LastDirection != math.Idle {
		snake.Move(snake.LastDirection)
	}
}

func (snake *Snake) Debug(render *sdl.Renderer) {
	utils.DrawBoundingBox(color.Purple, snake.Head.BoundingBox.X, snake.Head.BoundingBox.Y, snake.Head.BoundingBox.W, render)
	nextSegment := snake.Head.Next
	for i := 0; i < snake.Length; i++ {
		utils.DrawBoundingBox(color.Purple, nextSegment.BoundingBox.X, nextSegment.BoundingBox.Y, nextSegment.BoundingBox.W, render)
		nextSegment = nextSegment.Next
	}
}
