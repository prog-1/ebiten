package game

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type imgFunc func(*ebiten.Image, image.Point)

type Game struct {
	imgs       []*ebiten.Image
	pos        []image.Point
	lastUpdate time.Time

	screenWidth  int
	screenHeight int

	addRect func() bool
	newRect func() (*ebiten.Image, image.Point)
}

func New(scrW, scrH, minRect int, opts ...OptionFunc) *Game {
	o := Options{
		ScreenWidth:  scrW,
		ScreenHeight: scrH,
		MinRect:      minRect,
	}
	for _, f := range opts {
		f(&o)
	}
	if o.AddRect == nil {
		EverySecond()(&o)
	}
	if o.NewRect == nil {
		RandomRect()(&o)
	}
	return &Game{
		screenWidth:  o.ScreenWidth,
		screenHeight: o.ScreenHeight,
		addRect:      o.AddRect,
		newRect:      o.NewRect,
	}
}

func (g *Game) Update() error {
	if g.addRect() {
		img, pos := g.newRect()
		g.imgs = append(g.imgs, img)
		g.pos = append(g.pos, pos)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, img := range g.imgs {
		var opt ebiten.DrawImageOptions
		opt.GeoM.Translate(float64(g.pos[i].X), float64(g.pos[i].Y))
		screen.DrawImage(img, &opt)
	}
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) { return g.screenWidth, g.screenHeight }
