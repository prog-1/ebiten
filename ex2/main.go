package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/prog-1/ebiten-school/ex2/game"
)

const (
	screenWidth  = 320
	screenHeight = 200

	minRect = 10
)

func main() {
	rand.Seed(time.Now().UnixNano())
	g := game.New(screenWidth, screenHeight, minRect)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
