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
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth  = 640
	screenHeight = 480

	dpi = 72
)

type coloredRect struct {
	image.Rectangle
	angle float64
	color.RGBA
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
	for _, r := range g.rect {
		r.angle -= math.Pi / 90 // constantly updating rotation
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	t := time.Now()
	if t.Sub(g.last).Milliseconds() < 500 {
		return nil
	}
	g.last = t
	g.rect = append(g.rect, randomRect(g.width, g.height))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, r := range g.rect {
		rectImg := ebiten.NewImage(r.Dx(), r.Dy())                                          //converting rect to image
		rectImg.Fill(r.RGBA)                                                                //filling image with rect color
		g := ebiten.GeoM{}                                                                  //declaring geometry operations
		g.Translate(-float64(r.Dx())/2, -float64(r.Dy())/2)                                 //translating rect to rotate it
		g.Rotate(r.angle)                                                                   //rotating rect
		g.Translate(float64(r.Min.X)+float64(r.Dx())/2, float64(r.Min.Y)+float64(r.Dy())/2) //moving it back on it's location
		screen.DrawImage(rectImg, &ebiten.DrawImageOptions{                                 //drawing the rectangle
			GeoM: g,
		})
	}
	const msg = "Wait for new rectangles\nPress Q to quit"
	r := text.BoundString(g.font, msg)
	text.Draw(screen, msg, g.font, (screen.Bounds().Dx()-r.Dx())/2, screen.Bounds().Dy()/2, color.White)
}

func NewGame(width, height int, f font.Face) *Game {
	return &Game{width: width, height: height, font: f}
}

func randomRect(width, height int) *coloredRect {
	x0, y0 := rand.Intn(width), rand.Intn(height)
	x1, y1 := rand.Intn(width-x0)+x0, rand.Intn(height-y0)+y0
	rect := image.Rect(x0, y0, x1, y1)

	angle := rand.Float64() * 2 * math.Pi //choosing random starting angle

	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	a := uint8(rand.Intn(255))
	col := color.RGBA{R: r, G: g, B: b, A: a}

	return &coloredRect{rect, angle, col}
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
