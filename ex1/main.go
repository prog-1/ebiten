package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type game struct{}

func (*game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (*game) Update() error                             { return nil }
func (*game) Draw(screen *ebiten.Image)                 {}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
