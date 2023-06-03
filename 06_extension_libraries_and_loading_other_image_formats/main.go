package main

import (
	"fmt"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/veandco/go-sdl2/img"
)

const (
	width  = 600
	height = 480
)

func main() {
	win, _ := createWindow(width, height)
	surface, _ := win.GetSurface()

	bulbasaurSurface, err := loadBulbasaurSurface(surface.Format)
	if err != nil {
		log.Fatal(err)
	}

	rect := &sdl.Rect{
		W: width,
		H: height,
	}

	err = bulbasaurSurface.BlitScaled(nil, surface, rect)
	if err != nil {
		fmt.Println("Unable to apply bulbarsaur surface to screen surface")
		log.Fatal(err)
	}
	bulbasaurSurface.Free()

	err = win.UpdateSurface()
	if err != nil {
		fmt.Println("Unable to apply bulbarsaur surface to screen surface")
		log.Fatal(err)
	}

	time.Sleep(time.Second * 5)

	defer func() {
		surface.Free()
		win.Destroy()
	}()
}

func loadBulbasaurSurface(format *sdl.PixelFormat) (*sdl.Surface, error) {
	// prepare img loading
	err := img.Init(img.INIT_PNG)
	if err != nil {
		fmt.Println("Unable to init img")
		return nil, err
	}

	bulbasaurSurface, err := img.Load("../bulbasaur.png")
	if err != nil {
		fmt.Println("Unable to load bulbasaur surface")
		return nil, err
	}

	return bulbasaurSurface.Convert(format, 0)
}

func createWindow(w, h int32) (*sdl.Window, error) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, err
	}

	return sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		w, h, sdl.WINDOW_SHOWN)
}
