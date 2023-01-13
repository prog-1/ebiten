package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480

	// These are angles per second.
	minW = math.Pi / 100
	maxW = math.Pi / 2
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(screenWidth, screenHeight)

	g := NewGame(screenWidth, screenHeight, minW, maxW)
	// g.AddRotatingImage()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
