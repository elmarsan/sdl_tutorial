package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 600
	height = 480
)

func main() {
	window, renderer, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	err = renderer.SetDrawColor(0xff, 0xff, 0xff, 0xff)
	if err != nil {
		fmt.Println("Unable to set renderer color")
		log.Fatal(err)
	}

	texture := loadTexture(renderer)
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			default:
				renderer.Clear()

				renderer.Copy(texture, nil, nil)

				renderer.Present()
			}

		}
	}

	defer func() {
		renderer.Destroy()
		window.Destroy()
		texture.Destroy()
	}()
}

func loadTexture(renderer *sdl.Renderer) *sdl.Texture {
	img.Init(img.INIT_JPG)

	texture, err := img.LoadTexture(renderer, "../tile.jpg")
	if err != nil {
		log.Fatalf("Unable to load tile surface: %s", err)
	}

	// texture, err := renderer.CreateTextureFromSurface(surface)
	// if err != nil {
	// 	log.Fatalf("Unable to create texture from surface: %s", err)
	// }

	// surface.Free()

	return texture
}
