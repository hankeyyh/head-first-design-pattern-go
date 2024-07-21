package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type LoadStatus int

const (
	UnLoad LoadStatus = iota
	Loading
	Loaded
)

type ImageIcon struct {
	filename   string
	url        string
	loadStatus LoadStatus
}

type ImageProxy struct {
	imageMetadata map[string]*ImageIcon
	curImage      string
	image         *canvas.Image
	text          *canvas.Text
}

func NewImageProxy(nameUrlMap map[string]string) *ImageProxy {
	result := new(ImageProxy)
	result.imageMetadata = make(map[string]*ImageIcon)
	for filename, url := range nameUrlMap {
		icon := &ImageIcon{url: url, loadStatus: UnLoad, filename: filename}
		if fileExists(filename) {
			icon.loadStatus = Loaded
		}
		result.imageMetadata[filename] = icon
	}
	result.image = canvas.NewImageFromFile("")
	result.image.FillMode = canvas.ImageFillOriginal
	result.text = canvas.NewText("", color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	result.image.Hide()
	result.text.Hide()

	return result
}

func (ip *ImageProxy) PaintIcon(name string) {
	ip.curImage = name
	// image is loaded, display the image
	if ip.imageMetadata[name].loadStatus == Loaded {
		ip.paintImage(name)
		return
	}
	// image is loading, display loading text
	ip.paintText("Loading...")
	if ip.imageMetadata[name].loadStatus == Loading {
		return
	}
	// iamge is unload, start loading
	ip.imageMetadata[name].loadStatus = Loading

	go ip.loadImage(name)
}

func (ip *ImageProxy) loadImage(name string) {
	// take more time
	time.Sleep(time.Second * 2)
	
	err := downloadFile(name, ip.imageMetadata[name].url)
	if err != nil {
		ip.imageMetadata[name].loadStatus = UnLoad
	} else {
		ip.imageMetadata[name].loadStatus = Loaded
	}
	// if the image is not the current image, do not display it
	if name != ip.curImage {
		return
	}
	// display fail text or image
	if err != nil {
		ip.paintText(err.Error())
	} else {
		ip.paintImage(name)
	}
}

func (ip *ImageProxy) paintImage(name string) {
	ip.image.File = ip.imageMetadata[name].filename
	ip.image.Show()
	ip.text.Hide()
	ip.image.Refresh()
}

func (ip *ImageProxy) paintText(text string) {
	ip.image.Hide()
	ip.text.Text = text
	ip.text.Show()
}

func (ip *ImageProxy) GetContent() fyne.CanvasObject {
	return container.New(layout.NewCenterLayout(), ip.text, ip.image)
}
