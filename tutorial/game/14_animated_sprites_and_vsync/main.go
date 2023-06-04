package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480

	frames = 4
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

	spriteClips := [frames]sdl.Rect{
		{
			W: 64,
			H: 205,
		},
		{
			X: 64,
			W: 64,
			H: 205,
		},
		{
			X: 128,
			W: 64,
			H: 205,
		},
		{
			X: 192,
			W: 64,
			H: 205,
		},
	}

	texture, w, h := loadTexture(renderer, "../../asset/stick_walking.png")
	frame := 0

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break

			default:
				renderer.SetDrawColor(0xff, 0xff, 0xff, 0)
				renderer.Clear()
				renderer.Copy(texture, &spriteClips[frame], &sdl.Rect{
					W: w,
					H: h,
				})
				renderer.Present()

				frame++
				if frame >= 4 {
					frame = 0
				}

				time.Sleep(time.Second / 8)
			}
		}
	}

}

func loadTexture(r *sdl.Renderer, path string) (*sdl.Texture, int32, int32) {
	err := img.Init(img.INIT_PNG)
	if err != nil {
		log.Fatal(err)
	}

	surface, err := img.Load(path)
	if err != nil {
		log.Fatal(err)
	}

	colorKey := sdl.MapRGB(surface.Format, 0, 0xff, 0xff)
	surface.SetColorKey(true, colorKey)

	texture, err := r.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatal(err)
	}

	w, h := surface.W, surface.H
	surface.Free()

	return texture, w, h
}
