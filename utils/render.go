package utils

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func Draw(color sdl.Color, x int32, y int32, render *sdl.Renderer) {
	rect := sdl.Rect{X: x, Y: y, W: 32, H: 32}
	render.SetDrawColor(color.R, color.G, color.B, color.A)
	err := render.FillRect(&rect)
	if err != nil {
		fmt.Println("unable to draw texture: ", err)
	}
}

func DrawBoundingBox(color sdl.Color, x float64, y float64, w float64, render *sdl.Renderer) {
	render.SetDrawColor(color.R, color.G, color.B, color.A)
	borderWidth := int32(3)
	render.DrawRect(&sdl.Rect{
		X: (int32(x) - borderWidth),
		Y: (int32(y) - borderWidth),
		W: (int32(w) + 2*borderWidth),
		H: (int32(w) + 2*borderWidth)})
}

func DrawText(text string, fontSize int, color sdl.Color, textRect sdl.Rect, render *sdl.Renderer) {
	if err := ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()

	font, err := ttf.OpenFont("../assets/test.ttf", fontSize)
	if err != nil {
		panic(err)
	}
	defer font.Close()

	surface, err := font.RenderUTF8Solid(text, color)
	if err != nil {
		panic(err)
	}
	defer surface.Free()

	texture, err := render.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	render.Copy(texture, nil, &textRect)
}
