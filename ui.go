package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AppState struct {
	isUpdating         bool
	slR, slG, slB      *widget.Slider
	slC, slM, slY, slK *widget.Slider
	slH, slL, slS      *widget.Slider
	enR, enG, enB      *widget.Entry
	enC, enM, enY, enK *widget.Entry
	enH, enL, enS      *widget.Entry
}

func (app *AppState) createChannel(name string, mn, mx float64, onChange func()) (*widget.Slider, *widget.Entry, *fyne.Container) {
	slider := widget.NewSlider(mn, mx)
	slider.Step = 1
	entry := widget.NewEntry()
	entry.SetText(fmt.Sprintf("%.0f", mn))
	label := widget.NewLabel(name)
	slider.OnChanged = func(val float64) {
		entry.SetText(fmt.Sprintf("%.0f", val))
		if !app.isUpdating {
			onChange()
		}
	}
	entry.OnChanged = func(text string) {
		val, err := strconv.ParseFloat(text, 64)
		if err == nil && val >= mn && val <= mx {
			slider.SetValue(val)
		}
	}
	entrySize := container.NewGridWrap(fyne.NewSize(50, 35), entry)
	row := container.NewBorder(nil, nil, label, entrySize, slider)
	return slider, entry, row
}

func (app *AppState) buildRightPanel() *fyne.Container {
	var rowR, rowG, rowB *fyne.Container
	app.slR, app.enR, rowR = app.createChannel("R (Red):", 0, 255, app.updateFromRGB)
	app.slG, app.enG, rowG = app.createChannel("G (Green):", 0, 255, app.updateFromRGB)
	app.slB, app.enB, rowB = app.createChannel("B (Blue):", 0, 255, app.updateFromRGB)

	rgbCard := widget.NewCard("Модель RGB", "Аддитивная модель", container.NewVBox(rowR, rowG, rowB))

	var rowC, rowM, rowY, rowK *fyne.Container
	app.slC, app.enC, rowC = app.createChannel("C (Cyan):", 0, 100, app.updateFromCMYK)
	app.slM, app.enM, rowM = app.createChannel("M (Magenta):", 0, 100, app.updateFromCMYK)
	app.slY, app.enY, rowY = app.createChannel("Y (Yellow):", 0, 100, app.updateFromCMYK)
	app.slK, app.enK, rowK = app.createChannel("K (Key/Black):", 0, 100, app.updateFromCMYK)

	cmykCard := widget.NewCard("Модель CMYK", "Субтрактивная модель", container.NewVBox(rowC, rowM, rowY, rowK))

	var rowH, rowL, rowS *fyne.Container
	app.slH, app.enH, rowH = app.createChannel("H (Hue):", 0, 360, app.updateFromHLS)
	app.slL, app.enL, rowL = app.createChannel("L (Lightness):", 0, 100, app.updateFromHLS)
	app.slS, app.enS, rowS = app.createChannel("S (Saturation):", 0, 100, app.updateFromHLS)

	hlsCard := widget.NewCard("Модель HLS", "Перцепционная модель", container.NewVBox(rowH, rowL, rowS))

	return container.NewVBox(rgbCard, cmykCard, hlsCard)
}

func (app *AppState) updateFromRGB()  {}
func (app *AppState) updateFromCMYK() {}
func (app *AppState) updateFromHLS()  {}
