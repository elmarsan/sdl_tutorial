package main

import (
	"image/color"
	"log"
	"sdl_tutorial/tutorial/game/internal/texture"

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

	colorKey := &color.RGBA{
		R: 0,
		G: 0xff,
		B: 0xff,
	}

	t, err := texture.New(renderer, "../../asset/circle_sprites.png", colorKey)
	if err != nil {
		log.Fatal(err)
	}

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

				// Render red circle
				redCircle := &sdl.Rect{
					X: 0,
					Y: 0,
					W: 100,
					H: 100,
				}

				renderer.Copy(t.SDLTexture, redCircle, &sdl.Rect{
					W: t.W,
					H: t.H,
				})

				// Render blue circle
				blueCircle := &sdl.Rect{
					X: 100,
					Y: 100,
					W: 100,
					H: 100,
				}

				renderer.Copy(t.SDLTexture, blueCircle, &sdl.Rect{
					W: t.W,
					H: t.H,
					X: width - t.W,
					Y: height - t.H,
				})

				// Render green circle
				greenCircle := &sdl.Rect{
					X: 100,
					Y: 0,
					W: 100,
					H: 100,
				}

				renderer.Copy(t.SDLTexture, greenCircle, &sdl.Rect{
					W: t.W,
					H: t.H,
					X: width - t.W,
					Y: 0,
				})

				// Render yellow circle
				yellowCircle := &sdl.Rect{
					X: 0,
					Y: 100,
					W: 100,
					H: 100,
				}

				renderer.Copy(t.SDLTexture, yellowCircle, &sdl.Rect{
					X: 0,
					Y: height - t.H,
					W: t.W,
					H: t.H,
				})

				renderer.Present()
			}
		}
	}

	defer func() {
		t.Destroy()
		renderer.Destroy()
		win.Destroy()
	}()
}
