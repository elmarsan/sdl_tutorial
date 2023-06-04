package font

import (
	"log"

	"github.com/veandco/go-sdl2/ttf"
)

func Load(size int) *ttf.Font {
	if err := ttf.Init(); err != nil {
		log.Fatal(err)
	}

	font, err := ttf.OpenFont("../../asset/minecraft.ttf", size)
	if err != nil {
		log.Fatal(err)
	}

	return font
}
