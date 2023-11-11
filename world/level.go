package world

import (
	"github.com/solarlune/resolv"
)

var Space *resolv.Space
var Debug bool = false
var GameOver bool = false

func Init() {
	Space = resolv.NewSpace(800, 600, 32, 32)
}
