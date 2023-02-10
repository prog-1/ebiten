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
	offsetX, offsetY int
	rotation         float64
	speed            float64
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
	for _, r := range g.rect {
		r.rotation += r.speed
	}
	if t.Sub(g.last).Milliseconds() < 500 && !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	g.last = t
	g.rect = append(g.rect, randomRect(g.width, g.height))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, r := range g.rect {
		w, h := r.Size()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Rotate(float64(int(r.rotation)%360) * 2 * math.Pi / 360)
		op.GeoM.Translate(float64(r.offsetX), float64(r.offsetY))
		screen.DrawImage(r.Image, op)
	}
	const msg = "Press Q to quit...\nClick or wait for new rectangles."
	r := text.BoundString(g.font, msg)
	text.Draw(screen, msg, g.font, (screen.Bounds().Dx()-r.Dx())/2, screen.Bounds().Dy()/2, color.White)
}

func NewGame(width, height int, f font.Face) *Game {
	return &Game{width: width, height: height, font: f}
}

func randomRect(width, height int) *coloredRect {
	x0, y0 := rand.Intn(width)+1, rand.Intn(height)+1
	x1, y1 := rand.Intn(width-x0)+x0+1, rand.Intn(height-y0)+y0+1

	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	a := uint8(rand.Intn(255))
	col := color.RGBA{R: r, G: g, B: b, A: a}
	rect := ebiten.NewImage(x1-x0, y1-y0)
	rect.Fill(col)
	return &coloredRect{rect, x0, y0, 0, float64(rand.Intn(49)+1) / 10}
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
