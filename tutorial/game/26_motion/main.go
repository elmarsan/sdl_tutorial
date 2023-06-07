package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
	defaultSpeed = 2
	fps          = 60
)

func main() {
	win, renderer, err := sdl.CreateWindowAndRenderer(screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	dot, err := newDot(renderer, "../../asset/dot.bmp")
	if err != nil {
		log.Fatal(err)
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
				dot.handleEvent(event)
			}
		}

		dot.move()
		dot.render()

		elapsedTime := time.Since(start)
		frameDelay := time.Second/time.Duration(fps) - elapsedTime

		win.SetTitle(frameDelay.String())
		if frameDelay > 0 {

			sdl.Delay(uint32(frameDelay.Milliseconds()))
		}
	}

	defer func() {
		dot.texture.Destroy()
		renderer.Destroy()
		sdl.Quit()
	}()
}

type dot struct {
	velX   int32
	velY   int32
	width  int32
	height int32
	xPos   int32
	yPos   int32
	speed  int32

	texture  *sdl.Texture
	renderer *sdl.Renderer
}

func newDot(r *sdl.Renderer, imgPath string) (*dot, error) {
	surface, err := sdl.LoadBMP(imgPath)
	if err != nil {
		return nil, err
	}

	texture, err := r.CreateTextureFromSurface(surface)

	if err != nil {
		return nil, err
	}

	dot := &dot{
		width:    surface.W,
		height:   surface.H,
		yPos:     0,
		xPos:     0,
		speed:    defaultSpeed,
		texture:  texture,
		renderer: r,
	}

	surface.Free()

	return dot, nil
}

func (d *dot) handleEvent(e sdl.Event) {
	switch t := e.(type) {
	case *sdl.KeyboardEvent:
		keyCode := t.Keysym.Scancode

		if t.Repeat == 0 {
			switch keyCode {
			case sdl.SCANCODE_UP:
				d.velY -= d.speed
			case sdl.SCANCODE_DOWN:
				d.velY += d.speed
			case sdl.SCANCODE_LEFT:
				d.velX -= d.speed
			case sdl.SCANCODE_RIGHT:
				d.velX += d.speed
			}
		} else {
			switch keyCode {
			case sdl.SCANCODE_UP:
				d.velY += d.speed
			case sdl.SCANCODE_DOWN:
				d.velY -= d.speed
			case sdl.SCANCODE_LEFT:
				d.velX += d.speed
			case sdl.SCANCODE_RIGHT:
				d.velX -= d.speed
			}
		}
	}
}

func (d *dot) move() {
	d.yPos += d.velY

	if d.yPos < 0 || d.yPos+d.height > screenHeight {
		d.yPos -= d.velY
	}

	d.xPos += d.velX
	if d.xPos < 0 || d.xPos+d.width > screenWidth {
		d.xPos -= d.velX
	}
}

func (d *dot) render() {
	d.renderer.SetDrawColor(0xff, 0xff, 0xff, 0)
	d.renderer.Clear()

	rc := &sdl.Rect{
		X: d.xPos,
		Y: d.yPos,
		W: d.width,
		H: d.height,
	}

	d.renderer.Copy(d.texture, nil, rc)
	d.renderer.Present()
}
