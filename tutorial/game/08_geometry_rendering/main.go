package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 640
	height = 480
)

func main() {
	_, renderer, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
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
				renderer.SetDrawColor(0, 0, 0, 0xff)
				renderer.Clear()

				rect := &sdl.Rect{
					X: width / 4,
					Y: height / 4,
					W: width / 2,
					H: height / 2,
				}

				renderer.SetDrawColor(0, 0xff, 0, 0xff)
				renderer.FillRect(rect)

				renderer.SetDrawColor(0xff, 0, 0, 0)

				// 2/3 Off total screen
				renderer.DrawRect(&sdl.Rect{
					X: width / 6,
					Y: height / 6,
					W: width * 2 / 3,
					H: height * 2 / 3,
				})

				renderer.SetDrawColor(0, 0, 0xff, 0xff)
				// horizontal blue line
				renderer.DrawLine(0, height/2, width, height/2)

				// vertical yellow dots
				renderer.SetDrawColor(0xff, 0xff, 0, 0xff)
				for i := int32(0); i < height; i += 4 {
					renderer.DrawPoint(width/2, i)
				}

				renderer.Present()
			}
		}
	}

}
