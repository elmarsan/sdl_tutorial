package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	quit := false

	for !quit {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quit = false
				break
			}
		}
	}
}
