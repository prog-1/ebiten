package main

import (
	"image/color"
	"log"
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

type Game struct {
	width, height int
	font          font.Face
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (*Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	red := uint8(rand.Intn(255))
	green := uint8(rand.Intn(255))
	blue := uint8(rand.Intn(255))
	alpha := uint8(rand.Intn(255))
	screen.Fill(color.RGBA{R: red, G: green, B: blue, A: alpha})

	const msg = "Press Q to quit..."
	r := text.BoundString(g.font, msg)
	text.Draw(screen, msg, g.font, (screen.Bounds().Dx()-r.Dx())/2, screen.Bounds().Dy()/2, color.White)
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
		Size:    48,
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
