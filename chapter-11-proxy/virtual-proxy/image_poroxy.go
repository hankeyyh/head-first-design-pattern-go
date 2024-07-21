package main

import (
	"fmt"
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
	loadingText   *canvas.Text
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

	// TODO 控制生命周期
	go ip.loadImage(name)
}

func (ip *ImageProxy) loadImage(name string) {
	time.Sleep(time.Second * 2)
	err := downloadFile(name, ip.imageMetadata[name].url)
	if err != nil {
		// display text error
		fmt.Println(err)
		ip.imageMetadata[name].loadStatus = UnLoad
		return
	}
	ip.imageMetadata[name].loadStatus = Loaded
	// if still waiting, display the image
	if ip.curImage == name {
		ip.paintImage(name)
	}
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
