package main

import (
	"log"
	"sdl_tutorial/tutorial/game/internal/dot"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
	fps          = 60
)

func main() {
	win, renderer, err := sdl.CreateWindowAndRenderer(screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	dot, err := dot.New(renderer, 2, screenWidth, screenHeight)
	if err != nil {
		log.Fatal(err)
	}

	wall := &sdl.Rect{
		W: 30,
		H: 250,
		X: screenWidth / 2,
		Y: 0,
	}

	running := true
	for running {
		start := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break

			default:
				dot.HandleEvent(event)
			}
		}

		dot.Move(wall)

		renderer.SetDrawColor(0xff, 0xff, 0xff, 0)
		renderer.Clear()

		dot.Render()

		renderer.SetDrawColor(0, 0, 0xff, 0)
		renderer.FillRect(wall)

		renderer.Present()

		elapsedTime := time.Since(start)
		frameDelay := time.Second/time.Duration(fps) - elapsedTime

		win.SetTitle(frameDelay.String())
		if frameDelay > 0 {

			sdl.Delay(uint32(frameDelay.Milliseconds()))
		}
	}
}
