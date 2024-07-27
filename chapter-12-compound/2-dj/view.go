package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	viewcomponent "github.com/hankeyyh/head-first-design-pattern-go/chapter-12-compound/2-dj/view_component"
)

type DJView interface {
	BeatObserver
	BpmObserver

	CreateWindow(name string)
	Show()
	Quit()

	SetController(controller Controller)

	EnableStartMenuItem()
	DisableStartMenuItem()
	EnableStopMenuItem()
	DisableStopMenuItem()
}

type BasicDJView struct {
	controller Controller
	model      BeatModel

	app      fyne.App
	window   fyne.Window
	pulseBar *viewcomponent.BeatBar
	bpmLabel *widget.Label

	startMenuItem *fyne.MenuItem
	stopMenuItem  *fyne.MenuItem
	quitMenuItem  *fyne.MenuItem
}

func NewBasicDJView(model BeatModel) *BasicDJView {
	result := &BasicDJView{
		model: model,
	}
	model.RegisterBeatObserver(result)
	model.RegisterBpmObserver(result)
	return result
}

func (b *BasicDJView) CreateWindow(name string) {
	b.app = app.New()
	b.window = b.app.NewWindow(name)
	b.window.Resize(fyne.NewSize(300, 200))

	controlPanel := b.createControlPanel()
	bpmPanel := b.createBpmPanel()
	vbox := container.NewVBox(controlPanel, bpmPanel)
	b.window.SetContent(vbox)
	b.window.SetMainMenu(b.createMenu())
}

func (b *BasicDJView) createControlPanel() *fyne.Container {
	// numentry
	entry := viewcomponent.NewNumEntry()
	entry.SetPlaceHolder("Must contain a number")

	form := widget.NewForm(&widget.FormItem{
		Text:   "EnterBpm: ",
		Widget: entry,
	})
	form.OnSubmit = func() {
		bpm, _ := strconv.Atoi(entry.Text)
		b.controller.SetBpm(bpm)
	}

	// increase button
	increaseButton := widget.NewButton("Increase", func() {
		b.controller.IncreaseBpm()
	})

	// decrease button
	decreaseButton := widget.NewButton("Decrease", func() {
		b.controller.DecreaseBpm()
	})

	buttonContainer := container.NewHBox(layout.NewSpacer(), decreaseButton, increaseButton, layout.NewSpacer())

	vbox := container.NewVBox(form, buttonContainer)
	return vbox
}

func (b *BasicDJView) createBpmPanel() *fyne.Container {
	b.pulseBar = viewcomponent.NewBeatBar()
	b.bpmLabel = widget.NewLabel("0")

	hbox := container.NewHBox(widget.NewLabel("Current Bpm: "), b.bpmLabel)
	vbox := container.NewVBox(b.pulseBar, hbox)
	return vbox
}

func (b *BasicDJView) createMenu() *fyne.MainMenu {
	b.startMenuItem = fyne.NewMenuItem("Start", func() {
		b.controller.Start()
	})
	b.stopMenuItem = fyne.NewMenuItem("Stop", func() {
		b.controller.Stop()
	})
	b.quitMenuItem = fyne.NewMenuItem("Quit", func() {
		b.controller.Quit()
	})

	ctrlMenu := fyne.NewMenu("DJ Control", b.startMenuItem, b.stopMenuItem, b.quitMenuItem)
	main := fyne.NewMainMenu(ctrlMenu)
	return main
}

func (b *BasicDJView) Show() {
	b.controller.Start()
	b.window.ShowAndRun()
}

func (b *BasicDJView) Quit() {
	b.pulseBar.Close()
	b.window.Close()
}

func (b *BasicDJView) SetController(controller Controller) {
	b.controller = controller
}

func (b *BasicDJView) UpdateBeat() {
	b.pulseBar.Pulse()
}

func (b *BasicDJView) GetBeatObserverName() string {
	return "BasicDJView"
}

func (b *BasicDJView) UpdateBpm(bpm int) {
	b.bpmLabel.SetText(strconv.Itoa(bpm))
}

func (b *BasicDJView) GetBpmObserverName() string {
	return "BasicDJView"
}

func (b *BasicDJView) EnableStartMenuItem() {
	b.startMenuItem.Disabled = false
}

func (b *BasicDJView) DisableStartMenuItem() {
	b.startMenuItem.Disabled = true
}

func (b *BasicDJView) EnableStopMenuItem() {
	b.stopMenuItem.Disabled = false
}

func (b *BasicDJView) DisableStopMenuItem() {
	b.stopMenuItem.Disabled = true
}
