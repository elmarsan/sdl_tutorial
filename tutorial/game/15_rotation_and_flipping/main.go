package main

import (
	"log"
	"sdl_tutorial/tutorial/game/internal/texture"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

func main() {
	win, err := sdl.CreateWindow("14 - Animated sprites", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	renderer, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatal(err)
	}

	texture, _, _ := texture.Load(renderer, "../../asset/arrow.png")
	degrees := 0

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break

			case *sdl.KeyboardEvent:
				degrees += 60

			default:
				renderer.SetDrawColor(0xff, 0xff, 0xff, 0)
				renderer.Clear()
				renderer.CopyEx(texture, nil, nil, float64(degrees), nil, sdl.FLIP_HORIZONTAL)
				renderer.Present()
			}
		}
	}
}
