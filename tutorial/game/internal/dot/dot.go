package dot

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Dot struct {
	velX   int32
	velY   int32
	width  int32
	height int32
	xPos   int32
	yPos   int32
	speed  int32

	xLimit int32
	yLimit int32

	collider *sdl.Rect

	texture  *sdl.Texture
	renderer *sdl.Renderer
}

func New(r *sdl.Renderer, speed, xLimit, yLimit int32) (*Dot, error) {
	surface, err := sdl.LoadBMP("../../asset/dot.bmp")
	if err != nil {
		return nil, err
	}

	texture, err := r.CreateTextureFromSurface(surface)

	if err != nil {
		return nil, err
	}

	dot := &Dot{
		width:  surface.W,
		height: surface.H,
		yPos:   0,
		xPos:   0,
		speed:  speed,
		xLimit: xLimit,
		yLimit: yLimit,
		collider: &sdl.Rect{
			W: surface.W,
			H: surface.H,
		},
		texture:  texture,
		renderer: r,
	}

	surface.Free()

	return dot, nil
}

func (d *Dot) HandleEvent(e sdl.Event) {
	switch t := e.(type) {
	case *sdl.KeyboardEvent:
		keyCode := t.Keysym.Scancode

		fmt.Println(keyCode)

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

func (d *Dot) Move(rect *sdl.Rect) {
	d.yPos += d.velY
	d.collider.Y += d.velY

	if d.yPos < 0 || d.yPos+d.height > d.yLimit || d.checkCollision(rect) {
		d.yPos -= d.velY
		d.collider.Y = d.yPos
	}

	d.xPos += d.velX
	d.collider.X += d.velX
	if d.xPos < 0 || d.xPos+d.width > d.xLimit || d.checkCollision(rect) {
		d.xPos -= d.velX
		d.collider.X = d.xPos
	}
}

func (d *Dot) Render() {
	d.renderer.SetDrawColor(0xff, 0xff, 0xff, 0)

	rc := &sdl.Rect{
		X: d.xPos,
		Y: d.yPos,
		W: d.width,
		H: d.height,
	}

	d.renderer.Copy(d.texture, nil, rc)
	d.renderer.Present()
}

func (d *Dot) checkCollision(rect *sdl.Rect) bool {
	dotLeft := d.collider.X
	dotRight := d.collider.X + d.collider.W
	dotTop := d.collider.Y
	dotBottom := d.collider.Y + d.collider.H

	rectLeft := rect.X
	rectRight := rect.X + rect.W
	rectTop := rect.Y
	rectBottom := rect.Y + rect.H

	if dotBottom <= rectTop {
		return false
	}

	if dotTop >= rectBottom {
		return false
	}

	if dotRight <= rectLeft {
		return false
	}

	if dotLeft >= rectRight {
		return false
	}

	return true
}
