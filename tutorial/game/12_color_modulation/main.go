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

	err = img.Init(img.INIT_JPG)
	if err != nil {
		log.Fatal(err)
	}

	texture, err := img.LoadTexture(renderer, "../../asset/green-bg.jpg")
	if err != nil {
		log.Fatalf("Unable to create texture: %s", err)
	}

	var r, g, b uint8 = 0xff, 0xff, 0xff

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break

			case *sdl.KeyboardEvent:
				g += 10

			default:
				renderer.Clear()
				texture.SetColorMod(r, g, b)
				renderer.Copy(texture, nil, nil)
				renderer.Present()
			}
		}
	}

	defer func() {
		win.Destroy()
		renderer.Destroy()
		sdl.Quit()
	}()
}
