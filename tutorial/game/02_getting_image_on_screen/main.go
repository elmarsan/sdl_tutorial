package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	defer window.Destroy()

	screenSurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	defer screenSurface.Free()

	snailBmp, err := sdl.LoadBMP("../../asset/snail.bmp")
	if err != nil {
		panic(err)
	}

	defer snailBmp.Free()

	err = snailBmp.Blit(nil, screenSurface, nil)
	if err != nil {
		panic(err)
	}

	window.UpdateSurface()
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
}
