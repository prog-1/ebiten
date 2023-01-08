package main

import (
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/prog-1/ebiten-school/ex2/game"
)

const (
	screenWidth  = 320
	screenHeight = 200

	minRect = 10
)

func AtMouseCursor() game.OptionFunc {
	return func(o *game.Options) {
		o.AddRect = func() bool { return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) }
		o.NewRect = func() (*ebiten.Image, image.Point) {
			img := game.NewRect(o.ScreenWidth, o.ScreenHeight, o.MinRect)
			x, y := ebiten.CursorPosition()
			return img, image.Point{x, y}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := game.New(screenWidth, screenHeight, minRect, AtMouseCursor())
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
