package main

import (
	"image"
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
	image.Rectangle
	//ebiten.Image
	color.RGBA
	op     ebiten.DrawImageOptions
	rad    float64
	offset image.Point
}

func (r coloredRect) Draw(screen *ebiten.Image) {
	//ebitenutil.DrawRect(screen, float64(r.Min.X), float64(r.Min.Y), float64(r.Dx()), float64(r.Dy()), r.RGBA)
	image := ebiten.NewImage(r.Dx(), r.Dy())
	image.Fill(r.RGBA)
	screen.DrawImage(image, &r.op)
}

type coloredCircle struct {
	cx, cy, r float64
	color.RGBA
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

	t := time.Now()
	if t.Sub(g.last).Milliseconds() < 500 && !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	for _, r := range g.rects {
		w, h := r.Size().X, r.Size().Y
		r.op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		r.op.GeoM.Rotate(math.Pi / 6)
		r.op.GeoM.Translate(float64(r.offset.X), float64(r.offset.Y))
		//r.rad += math.Pi / 12
	}
	g.last = t
	g.rects = append(g.rects, randomRectangle(g.width, g.height))
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

func randomRectangle(width, height int) *coloredRect {
	x0, y0 := rand.Intn(width), rand.Intn(height)
	x1, y1 := rand.Intn(width-x0)+x0+1, rand.Intn(height-y0)+y0+1
	minRadians, maxRadians := -math.Pi, math.Pi
	radians := ((maxRadians-minRadians)*rand.Float64() + minRadians)
	cx, cy := x1-x0, y1-y0
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cx), float64(cy))
	return &coloredRect{image.Rect(x0, y0, x1, y1), randColor(), *op, radians, image.Point{x0, y0}}
	//rect := coloredRect{*ebiten.NewImage(x1-x0, y1-y0), ebiten.DrawImageOptions{}, radians}
	//rect.Fill(randColor())
	//return &rect
	//// NOTE: I end up copying image, which is forbidden
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

// How to rotate?
// We store matrix and angle inside coloredRectangle
// In Update() rotate each of rectangle's matrices
// In Draw() we draw rectangle using it's own matrix
