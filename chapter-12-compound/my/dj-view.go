package main

import "fmt"

type Observer interface {
	Update()
}

type View struct {
	model *Model
	controller *Controller
}

func NewView(model *Model, controller *Controller) *View {
	v := &View{model: model, controller: controller}
	return v
}

func (v *View) Start() {
	v.controller.Start()
}

func (v *View) Stop() {
	v.controller.Stop()
}

func (v *View) Quit() {
	v.controller.Quit()
}

func (v *View) Update() {
	v.DisplayText("BPM changed, now is %d\n", v.model.GetBpm())
}

func (v *View) Show() {
	v.DisplayText("Current BPM is %d\n", v.model.GetBpm())
}

func (v *View) SetBpm(bpm int) {
	v.controller.SetBpm(bpm)
}

func (v *View) IncreaseBpm() {
	v.controller.IncreaseBpm(1)
}

func (v *View) DecreaseBpm() {
	v.controller.IncreaseBpm(-1)
}

func (v *View) DisplayText(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}