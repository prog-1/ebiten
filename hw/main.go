package main

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth  = 640
	screenHeight = 480

	dpi = 72
)

type figureType int

const (
	Rectangle figureType = iota
	Circle
	FigureCount
)

type Drawable interface {
	Draw(*ebiten.Image)
}

type coloredRect struct {
	image.Rectangle
	color.RGBA
}

func (r coloredRect) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, float64(r.Min.X), float64(r.Min.Y), float64(r.Dx()), float64(r.Dy()), r.RGBA)
}

type coloredCircle struct {
	cx, cy, r float64
	color.RGBA
}

func (c coloredCircle) Draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, c.cx, c.cy, c.r, c.RGBA)
}

type Game struct {
	width, height int
	font          font.Face

	figures []*Drawable
	last    time.Time
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
	g.last = t
	tmp := randomFigure(g.width, g.height)
	g.figures = append(g.figures, &tmp)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, f := range g.figures {
		(*f).Draw(screen)

	}
	const msg = "Press Q to quit...\nClick or wait for new rectangles(or circles)."
	r := text.BoundString(g.font, msg)
	text.Draw(screen, msg, g.font, (screen.Bounds().Dx()-r.Dx())/2, screen.Bounds().Dy()/2, color.White)
}

func NewGame(width, height int, f font.Face) *Game {
	return &Game{width: width, height: height, font: f}
}

func randomFigure(width, height int) Drawable {
	var figType figureType = figureType(rand.Intn(int(FigureCount)))
	switch figType {
	case Rectangle:
		x0, y0 := rand.Intn(width), rand.Intn(height)
		x1, y1 := rand.Intn(width-x0)+x0, rand.Intn(height-y0)+y0
		rect := image.Rect(x0, y0, x1, y1)
		return &coloredRect{rect, randColor()}

	case Circle:
		cx, cy := rand.Intn(width), rand.Intn(height)
		const min, max = 2, 300
		r := rand.Intn(max-min+1) + min
		return &coloredCircle{float64(cx), float64(cy), float64(r), randColor()}

	default:
		panic("must not happen")
	}
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
