package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type game struct{}

func (*game) Update() error                             { return nil }
func (*game) Draw(screen *ebiten.Image)                 {}
func (*game) Layout(outWidth, outHeight int) (w, h int) { return 320, 200 }

func main() {
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
