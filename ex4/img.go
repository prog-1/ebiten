package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type RotatingImage struct {
	cx, cy float64
	img    *ebiten.Image
	dw, w  float64 // dw is rad/sec.
}

func NewRectImage(scrWidth, scrHeight int, minW, maxW float64) *RotatingImage {
	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	a := uint8(rand.Intn(255))
	clr := color.RGBA{R: r, G: g, B: b, A: a}

	var maxSize float64
	if scrWidth <= scrHeight {
		maxSize = float64(scrWidth) / math.Sqrt2
	} else {
		maxSize = float64(scrHeight) / math.Sqrt2
	}

	w, h := rand.Float64()*maxSize+1, rand.Float64()*maxSize+1
	m := w
	if h > w {
		m = h
	}
	cx := rand.Float64()*(float64(scrWidth)-m) + float64(m/2)
	cy := rand.Float64()*(float64(scrHeight)-m) + float64(m/2)

	omega := (maxW - minW) * (rand.Float64()*2 - 1)
	if omega >= 0 {
		omega += minW
	} else {
		omega -= minW
	}

	img := ebiten.NewImage(int(w), int(h))
	img.Fill(clr)

	return &RotatingImage{cx: cx, cy: cy, img: img, dw: omega}
}

func (ri *RotatingImage) Update(ms float64) error {
	if ri.w += ri.dw * ms / 1000; ri.w >= 2*math.Pi {
		ri.w -= 2 * math.Pi
	}
	return nil
}

func (ri *RotatingImage) Draw(screen *ebiten.Image) {
	imgCX := float64(ri.img.Bounds().Dx()) / 2
	imgCY := float64(ri.img.Bounds().Dy()) / 2
	var opts ebiten.DrawImageOptions
	opts.GeoM.Translate(-imgCX, -imgCY)
	opts.GeoM.Rotate(ri.w)
	opts.GeoM.Translate(ri.cx, ri.cy)
	screen.DrawImage(ri.img, &opts)
}
