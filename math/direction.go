package math

type Direction struct {
	X, Y int
}

var (
	Up    = Direction{X: 0, Y: -32}
	Down  = Direction{X: 0, Y: 32}
	Left  = Direction{X: -32, Y: 0}
	Right = Direction{X: 32, Y: 0}
	Idle  = Direction{X: 0, Y: 0}
	Pause = Direction{X: 1, Y: 1}
)
