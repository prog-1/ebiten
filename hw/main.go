package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 640
	screenHeight = 480
	dpi          = 72
)

type game struct {
	width, height int
	font          font.Face
	count         int
	circel        []*circel
	rect          []*coloredRect
	last          time.Time
}

type circel struct {
	cx int
	cy int
	r  int
	color.RGBA
}

type coloredRect struct {
	*ebiten.Image
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return g.width, g.height }
func (g *game) Update() error {
	g.count++
	t := time.Now()
	if t.Sub(g.last).Milliseconds() < 500 {
		return nil
	}
	g.last = t
	g.circel = append(g.circel, randomCircel(g.width, g.height))
	g.rect = append(g.rect, randomRect(g.width, g.height))
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	for _, r := range g.circel {
		ebitenutil.DrawCircle(screen, float64(r.cx), float64(r.cy), float64(r.r), r.RGBA)
		//drawCircel(screen, r.cx, r.cy, r.r, r.RGBA)
	}
	for _, rot := range g.rect {
		w, h := rot.Size()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Rotate(float64(rand.Intn(g.count)%360) * 2 * math.Pi / 360)
		op.GeoM.Translate(float64(screenWidth/2), screenHeight/2)
		screen.DrawImage(rot.Image, op)
	}
}

func randomCircel(width, height int) *circel {
	cx, cy, rad := rand.Intn(width), rand.Intn(height), rand.Intn((height+width)/4)

	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	a := uint8(rand.Intn(255))
	col := color.RGBA{R: r, G: g, B: b, A: a}

	return &circel{cx, cy, rad, col}
}
func randomRect(width, height int) *coloredRect {
	x0, y0 := rand.Intn(width), rand.Intn(height)
	x1, y1 := rand.Intn(width-x0)+x0, rand.Intn(height-y0)+y0
	rect := ebiten.NewImage(x1-x0+1, y1-y0+1)

	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	a := uint8(rand.Intn(255))
	col := color.RGBA{R: r, G: g, B: b, A: a}
	rect.Fill(col)
	return &coloredRect{rect}
}

func NewGame(width, height int, f font.Face) *game {
	return &game{width: width, height: height, font: f}
}

func main() {
	//ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Homework")
	//ebiten.SetWindowIcon()
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

/*func drawCircel(screen *ebiten.Image, x, y, rad int, clr color.Color) {
	rad64 := float64(rad)
	minAngle := math.Acos(1 - 1/rad64)
	for angle := float64(0); angle <= 360; angle += minAngle {
		xd := rad64 * math.Cos(angle)
		yd := rad64 * math.Sin(angle)
		x1 := int(math.Round(float64(x) + xd))
		y1 := int(math.Round(float64(y) + yd))
		screen.Set(x1, y1, clr)

	}
}*/ //this func draws just outer lines of circles
