package game

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Options struct {
	ScreenWidth  int
	ScreenHeight int
	MinRect      int
	AddRect      func() bool
	NewRect      func() (*ebiten.Image, image.Point)
}

type OptionFunc func(*Options)

func EverySecond() OptionFunc {
	return func(o *Options) {
		var last time.Time
		o.AddRect = func() bool {
			t := time.Now()
			if t.Sub(last).Seconds() < 1 {
				return false
			}
			last = t
			return true
		}
	}
}

func RandomRect() OptionFunc {
	return func(o *Options) {
		o.NewRect = func() (*ebiten.Image, image.Point) {
			img := NewRect(o.ScreenWidth, o.ScreenHeight, o.MinRect)
			w, h := img.Size()
			pos := image.Point{
				X: rand.Intn(o.ScreenWidth - w),
				Y: rand.Intn(o.ScreenHeight - h),
			}
			return img, pos
		}
	}
}

func NewRect(scrW, scrH, min int) *ebiten.Image {
	w := rand.Intn(scrW-min) + min
	h := rand.Intn(scrH-min) + min
	c := color.RGBA{
		R: uint8(rand.Intn(0xff)),
		G: uint8(rand.Intn(0xff)),
		B: uint8(rand.Intn(0xff)),
		A: uint8(rand.Intn(0xff)),
	}
	img := ebiten.NewImage(w, h)
	img.Fill(c)
	return img
}
