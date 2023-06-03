package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	win, _ := createWindow(600, 480)
	screenSurface, _ := win.GetSurface()

	// Load image and generate an optimized surface
	snailSurface, _ := loadSurface("../snail.bmp")
	snailOptSurface, err := snailSurface.Convert(screenSurface.Format, 0)
	if err != nil {
		fmt.Println("Unable to optimize snail surface")
		log.Fatal(err)
	}

	// We no longer need the original surface, we can drop it
	snailSurface.Free()

	// Apply the image stretched
	rect := &sdl.Rect{
		W: 600,
		H: 480,
	}

	// Optimize image
	err = snailOptSurface.BlitScaled(nil, screenSurface, rect)
	if err != nil {
		log.Fatal(err)
	}

	win.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}

	defer func() {
		win.Destroy()
		screenSurface.Free()
		snailOptSurface.Free()
	}()
}

func createWindow(w, h int32) (*sdl.Window, error) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, err
	}

	return sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		w, h, sdl.WINDOW_SHOWN)
}

func loadSurface(path string) (*sdl.Surface, error) {
	return sdl.LoadBMP(path)
}
