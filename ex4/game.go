package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	width, height int
	minW, maxW    float64
	imgs          []*RotatingImage
	lastUpd       time.Time
	lastAdd       time.Time
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	t := time.Now()
	dt := float64(t.Sub(g.lastUpd).Milliseconds())
	g.lastUpd = t

	for _, img := range g.imgs {
		if err := img.Update(dt); err != nil {
			return err
		}
	}

	if t.Sub(g.lastAdd).Milliseconds() >= 1500 {
		g.lastAdd = t
		g.AddRotatingImage()
	}

	return nil
}

func (g *Game) AddRotatingImage() {
	img := NewRectImage(g.width, g.height, g.minW, g.maxW)
	g.imgs = append(g.imgs, img)
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, img := range g.imgs {
		img.Draw(screen)
	}
}

func NewGame(width, height int, minW, maxW float64) *Game {
	return &Game{width: width, height: height, minW: minW, maxW: maxW}
}
