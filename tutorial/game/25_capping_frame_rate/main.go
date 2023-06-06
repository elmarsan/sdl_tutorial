package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

func main() {
	win, renderer, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	var (
		size int32 = 50
		x    int32 = 0
		y    int32 = 0
		fps        = 30
	)

	running := true
	for running {
		start := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		rect := &sdl.Rect{
			X: x,
			Y: y,
			W: size,
			H: size,
		}

		x += size

		if x >= width {
			y += size
			x = 0
		}

		if y > height {
			x = 0
			y = 0
		}

		renderer.SetDrawColor(0, 0xff, 0, 0)
		renderer.FillRect(rect)
		renderer.Present()

		sdl.Delay(16)

		elapsedTime := time.Since(start)
		frameDelay := time.Second/time.Duration(fps) - elapsedTime

		win.SetTitle(frameDelay.String())
		if frameDelay > 0 {
			time.Sleep(frameDelay)
		}
	}
}
