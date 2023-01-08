package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/prog-1/ebiten-school/ex2/game"
)

const (
	screenWidth  = 320
	screenHeight = 200

	minRect = 10
)

type Game struct{ game.Game }

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	return g.Game.Update()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := &Game{*game.New(screenWidth, screenHeight, minRect)}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
