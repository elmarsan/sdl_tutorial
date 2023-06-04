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
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	_, renderer, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatalf("Unable to get window surface %s", err)
	}

	img.Init(img.INIT_PNG)
	surface, err := img.Load("../pikachu.png")
	if err != nil {
		log.Fatalf("Unable to load surface: %s", err)
	}

	// Replace yellows by purple
	colorKey := sdl.MapRGB(surface.Format, 244, 212, 88)
	surface.SetColorKey(true, colorKey)

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatalf("Unable to create texture: %s", err)
	}
	w := surface.W
	h := surface.H

	surface.Free()

	renderer.SetDrawColor(0xff, 0xff, 0xff, 0xff)
	renderer.Clear()
	renderer.Copy(texture, nil, &sdl.Rect{
		W: w,
		H: h,
	})
	renderer.Present()

	time.Sleep(time.Second * 5)
}
