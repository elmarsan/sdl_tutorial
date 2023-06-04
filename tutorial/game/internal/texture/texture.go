package texture

import (
	"fmt"
	"image/color"
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Texture struct {
	SDLTexture *sdl.Texture
	W, H       int32
}

func New(renderer *sdl.Renderer, path string, colorKey *color.RGBA) (*Texture, error) {
	err := img.Init(img.INIT_PNG)
	if err != nil {
		return nil, err
	}

	surface, err := img.Load(path)
	if err != nil {
		return nil, err
	}

	if colorKey != nil {
		fmt.Println("Setting color key")

		key := sdl.MapRGB(surface.Format, colorKey.R, colorKey.G, colorKey.B)
		surface.SetColorKey(true, key)
	}

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}

	t := &Texture{
		SDLTexture: texture,
		W:          surface.W,
		H:          surface.H,
	}

	surface.Free()

	return t, nil
}

func (t *Texture) SetColor(r, g, b uint8) error {
	return t.SDLTexture.SetColorMod(r, g, b)
}

func (t *Texture) Destroy() error {
	return t.SDLTexture.Destroy()
}

func Load(r *sdl.Renderer, path string) (*sdl.Texture, int32, int32) {
	err := img.Init(img.INIT_PNG)
	if err != nil {
		log.Fatal(err)
	}

	surface, err := img.Load(path)
	if err != nil {
		log.Fatal(err)
	}

	texture, err := r.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatal(err)
	}

	w, h := surface.W, surface.H
	surface.Free()

	return texture, w, h
}
