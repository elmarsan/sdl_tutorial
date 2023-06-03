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
	_, renderer, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	bulbasaur := loadTexture("../bulbasaur.png", renderer)
	squirtle := loadTexture("../squirtle.png", renderer)

	topLeftViewport := &sdl.Rect{
		W: width / 2,
		H: height / 2,
	}
	renderer.SetViewport(topLeftViewport)
	renderer.Copy(bulbasaur, nil, nil)

	topRightViewport := &sdl.Rect{
		X: width / 2,
		W: width / 2,
		H: height / 2,
	}
	renderer.SetViewport(topRightViewport)
	renderer.Copy(bulbasaur, nil, nil)

	bottomViewport := &sdl.Rect{
		Y: height / 2,
		W: width,
		H: height / 2,
	}
	renderer.SetViewport(bottomViewport)
	renderer.Copy(squirtle, nil, nil)

	renderer.Present()
	time.Sleep(time.Second * 5)
}

func loadTexture(path string, renderer *sdl.Renderer) *sdl.Texture {
	img.Init(img.INIT_PNG)

	texture, err := img.LoadTexture(renderer, path)
	if err != nil {
		log.Fatalf("Unable to load bulbasaur texture: %s", err)
	}

	return texture
}
