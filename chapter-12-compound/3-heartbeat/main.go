package main

import "github.com/hankeyyh/head-first-design-pattern-go/chapter-12-compound/2-dj/mvc"

func main() {
	model := NewHeartModel()
	
	controller := NewHeatController(model)
	view := mvc.NewBasicDJView(NewHeartAdapter(model))
	controller.SetView(view)
	view.SetController(controller)
	
	view.CreateWindow("Heart Beat View")
	view.Show()
}