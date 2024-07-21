package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func rawload() {
	a := app.New()
	w := a.NewWindow("Menu Example")
	w.Resize(fyne.NewSize(400, 400))

	text := canvas.NewText("Empty", color.NRGBA{R: 0, G: 0, B: 0, A: 255})
	image := canvas.NewImageFromFile("")

	wselect := widget.NewSelect([]string{"Buddha Bar.jpg", "Ima.jpg", "MCMXC a.D.jpg", "text"}, nil)
	wselect.OnChanged = func(selected string) {
		fmt.Println(selected)
		if selected == "text" {
			image.Hide()
			text.Show()
		} else {
			image.File = selected
			image.FillMode = canvas.ImageFillOriginal
			image.Refresh()
			image.Show()
			text.Hide()
		}
	}

	center := container.New(layout.NewCenterLayout(), text, image)
	vbox := container.New(layout.NewVBoxLayout(), wselect, center)

	w.SetContent(vbox)
	w.ShowAndRun()
}

func proxyload() {
	a := app.New()
	w := a.NewWindow("Menu Example")
	w.Resize(fyne.NewSize(400, 400))

	imageStore := map[string]string{
		"Buddha Bar.jpg":                     "http://images.amazon.com/images/P/B00009XBYK.01.LZZZZZZZ.jpg",
		"Ima.jpg":                            "http://images.amazon.com/images/P/B000005IRM.01.LZZZZZZZ.jpg",
		"MCMXC a.D.jpg":                      "http://images.amazon.com/images/P/B000002URV.01.LZZZZZZZ.jpg",
		"Northern Exposure.jpg":              "http://images.amazon.com/images/P/B000003SFN.01.LZZZZZZZ.jpg",
		"Selected Ambient Works, Vol. 2.jpg": "http://images.amazon.com/images/P/B000002MNZ.01.LZZZZZZZ.jpg",
	}

	imageProxy := NewImageProxy(imageStore)

	keys := make([]string, 0, len(imageStore))
	for k := range imageStore {
		keys = append(keys, k)
	}
	wselect := widget.NewSelect(keys, nil)
	wselect.OnChanged = func(selected string) {
		imageProxy.PaintIcon(selected)
	}

	vbox := container.New(layout.NewVBoxLayout(), wselect, imageProxy.GetContent())

	w.SetContent(vbox)
	w.ShowAndRun()
}

func main() {
	// rawload()
	proxyload()
}