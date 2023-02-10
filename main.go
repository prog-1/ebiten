package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 1290
	screenHeight = 960
	dpi          = 72
)

type coloredRect struct {
	*ebiten.Image
	x, y int
	color.RGBA
	dir  int
	time time.Time
}

type Game struct {
	width, height int
	font          font.Face

	rect []*coloredRect
	last time.Time
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	t := time.Now()
	if t.Sub(g.last).Milliseconds() < 5000 && !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	g.last = t
	g.rect = append(g.rect, randomRect(g.width, g.height))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Rotate(screen)
	for _, r := range g.rect {
		t := time.Now()
		if t.Sub(r.time).Milliseconds() < 1 {
			g.Rotate(screen)
		}
	}
	const msg = "Press Q to quit..."
	r := text.BoundString(g.font, msg)
	text.Draw(screen, msg, g.font, (screen.Bounds().Dx()-r.Dx())/2, screen.Bounds().Dy()/2, color.White)
}

func (g *Game) Rotate(screen *ebiten.Image) {
	for _, r := range g.rect {
		w, h := r.Size()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		if r.dir == 0 {
			op.GeoM.Rotate(float64(-math.Pi / 4))
		} else {
			op.GeoM.Rotate(float64(math.Pi / 4))
		}
		op.GeoM.Translate(float64(r.x), float64(r.y))
		screen.DrawImage(r.Image, op)
	}
}

func randomRect(width, height int) *coloredRect {
	x0, y0 := rand.Intn(width), rand.Intn(height)
	x1, y1 := rand.Intn(width-x0)+x0+1, rand.Intn(height-y0)+y0+1
	rect := ebiten.NewImage(x1-x0, y1-y0)
	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	a := uint8(rand.Intn(255))
	col := color.RGBA{R: r, G: g, B: b, A: a}
	rect.Fill(col)

	t := time.Now()
	return &coloredRect{rect, x0, y0, col, rand.Intn(1), t}
}

func NewGame(width, height int, f font.Face) *Game {
	return &Game{width: width, height: height, font: f}
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
