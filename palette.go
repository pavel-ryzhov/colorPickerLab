package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type CustomPalette struct {
	widget.BaseWidget
	OnSelected func(hls HLS)
}

func NewCustomPalette(onSelected func(hls HLS)) *CustomPalette {
	p := &CustomPalette{OnSelected: onSelected}
	p.ExtendBaseWidget(p)
	return p
}

func (p *CustomPalette) CreateRenderer() fyne.WidgetRenderer {
	raster := canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		hue := float64(x) / float64(w) * 360.0
		lightness := 1.0 - float64(y)/float64(h)

		rgb := HLS{H: uint16(hue), L: uint8(lightness * 100.0), S: 100}.ToRGB()
		return color.RGBA{R: rgb.R, G: rgb.G, B: rgb.B, A: 255}
	})
	return widget.NewSimpleRenderer(raster)
}

func (p *CustomPalette) Tapped(e *fyne.PointEvent) {
	if p.OnSelected == nil {
		return
	}

	w, h := p.Size().Width, p.Size().Height
	hue := e.Position.X / w * 360.0
	lightness := 1.0 - e.Position.Y/h

	if hue < 0 {
		hue = 0
	}
	if hue >= 360 {
		hue = 0
	}
	if lightness < 0 {
		lightness = 0
	}
	if lightness > 1 {
		lightness = 1
	}
	p.OnSelected(HLS{
		H: uint16(hue),
		L: uint8(lightness * 100),
		S: 100,
	})
}
