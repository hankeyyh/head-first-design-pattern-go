package main

import "github.com/hankeyyh/head-first-design-pattern-go/chapter-12-compound/2-dj/mvc"

func main() {
	model := mvc.NewBasicBeatModel()
	if err := model.Initialize(); err != nil {
		panic(err)
	}

	controller := mvc.NewBeatController(model)
	view := mvc.NewBasicDJView(model)
	view.SetController(controller)
	controller.SetView(view)

	view.CreateWindow("DJ View")
	view.Show()
}