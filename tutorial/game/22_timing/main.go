package main

import (
	"fmt"
	"log"
	"sdl_tutorial/tutorial/game/internal/font"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

func main() {
	win, err := sdl.CreateWindow("22 - timming", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	renderer, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal(err)
	}

	font := font.Load(16)

	color := sdl.Color{
		R: 0xff,
		G: 0xff,
		B: 0xff,
	}

	start := 0

	helpText, err := font.RenderUTF8Solid("Press any key to reset timer", color)
	if err != nil {
		log.Fatal(err)
	}

	helpTexture, _ := renderer.CreateTextureFromSurface(helpText)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break

			case *sdl.KeyboardEvent:
				start = int(sdl.GetTicks64()) / 1000

			default:
				renderer.Clear()
				renderer.Copy(helpTexture, nil, &sdl.Rect{
					Y: 20,
					W: width,
					H: 20,
				})

				elapsedS := sdl.GetTicks64() / 1000

				textSurface, err := font.RenderUTF8Blended(fmt.Sprintf("Timer: %d", int(elapsedS)-start), color)
				if err != nil {
					log.Fatal(err)
				}

				texture, err := renderer.CreateTextureFromSurface(textSurface)
				if err != nil {
					log.Fatal(err)
				}

				renderer.Copy(texture, nil, &sdl.Rect{
					Y: height / 2,
					W: width / 2,
					H: 20,
				})

				renderer.Present()
			}
		}
	}
}
