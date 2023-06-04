package main

import (
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}

	win, renderer, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	_ = img.Init(img.INIT_PNG)
	texture, _ := img.LoadTexture(renderer, "../../asset/house.png")

	var alpha uint8 = 0xff

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break

			case *sdl.KeyboardEvent:
				alpha += 0xa

			default:
				renderer.Clear()
				texture.SetAlphaMod(123)
				renderer.Copy(texture, nil, nil)
				renderer.Present()
			}
		}
	}

	defer func() {
		renderer.Destroy()
		win.Destroy()
		sdl.Quit()
	}()
}
