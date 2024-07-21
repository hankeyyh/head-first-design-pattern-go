package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type LoadStatus int

const (
	UnStarted LoadStatus = iota
	Loading
	Loaded
)

type ImageIcon struct {
	name       string
	filename   string
	url        string
	loadStatus LoadStatus
}

type ImageProxy struct {
	imageMetadata map[string]*ImageIcon
	curImage      string
	image         *canvas.Image
	loadingText   *canvas.Text
}

func NewImageProxy(nameUrlMap map[string]string) *ImageProxy {
	result := new(ImageProxy)
	result.imageMetadata = make(map[string]*ImageIcon)
	for name, url := range nameUrlMap {
		icon := &ImageIcon{name: name, url: url, loadStatus: UnStarted}
		if fileExists(name) {
			icon.filename = name
			icon.loadStatus = Loaded
		}
		result.imageMetadata[name] = icon
	}
	result.image = canvas.NewImageFromFile("")
	result.image.FillMode = canvas.ImageFillOriginal
	result.loadingText = canvas.NewText("Loading...", color.Black)
	result.image.Hide()
	result.loadingText.Hide()
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
	if ip.imageMetadata[name].loadStatus == Loading {
		ip.paintLoadingText()
		return
	}
	// iamge is not start loading, start loading
	ip.imageMetadata[name].loadStatus = Loading
	ip.paintLoadingText()

	go func() {
		fileName, err := downloadImg(name, ip.imageMetadata[name].url)
		if err != nil {
			// display text error
			fmt.Println(err)
			ip.imageMetadata[name].loadStatus = UnStarted
			return
		}
		ip.imageMetadata[name].loadStatus = Loaded
		ip.imageMetadata[name].filename = fileName
		// if still waiting, display the image
		if ip.curImage == name {
			ip.paintImage(name)
		}
	}()
}

func (ip *ImageProxy) paintImage(name string) {
	ip.image.File = ip.imageMetadata[name].filename
	ip.image.Show()
	ip.loadingText.Hide()
	ip.image.Refresh()
}

func (ip *ImageProxy) paintLoadingText() {
	ip.image.Hide()
	ip.loadingText.Show()
}

func (ip *ImageProxy) GetContent() fyne.CanvasObject {
	return container.New(layout.NewCenterLayout(), ip.loadingText, ip.image)
}
