package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth  = 640
	screenHeight = 480

	dpi = 72
)

type coloredRect struct {
	*ebiten.Image
	op      *ebiten.DrawImageOptions
	degrees float64
}

func (r coloredRect) Draw(screen *ebiten.Image) {
	screen.DrawImage(r.Image, r.op)
}

type Game struct {
	width, height int
	font          font.Face

	rects []*coloredRect
	last  time.Time
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	for _, r := range g.rects {
		w, h := r.Size()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Rotate(r.degrees * math.Pi / 180)
		op.GeoM.Translate(float64(r.Bounds().Max.X), float64(r.Bounds().Max.Y))
		r.op = op
	}

	t := time.Now()
	if t.Sub(g.last).Milliseconds() < 500 && !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	g.last = t
	g.rects = append(g.rects, randomRect(g.width, g.height))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, f := range g.rects {
		(*f).Draw(screen)
	}
	const msg = "Press Q to quit...\nClick or wait for new rectangles(or circles)."
	r := text.BoundString(g.font, msg)
	text.Draw(screen, msg, g.font, (screen.Bounds().Dx()-r.Dx())/2, screen.Bounds().Dy()/2, color.White)
}

func NewGame(width, height int, f font.Face) *Game {
	return &Game{width: width, height: height, font: f}
}

func randomRect(width, height int) *coloredRect {
	x0, y0 := rand.Intn(width), rand.Intn(height)
	x1, y1 := rand.Intn(width-x0)+x0, rand.Intn(height-y0)+y0

	cx, cy := x1-x0, y1-y0
	image := ebiten.NewImage(cx+1, cy+1)
	image.Fill(randColor())
	return &coloredRect{image, nil, 0.001}
}

func randColor() color.RGBA {
	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	a := uint8(rand.Intn(255))

	return color.RGBA{R: r, G: g, B: b, A: a}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	g := NewGame(screenWidth, screenHeight, mplusNormalFont)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
