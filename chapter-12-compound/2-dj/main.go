package main

func main() {
	model := NewBasicBeatModel()
	model.Initialize()

	controller := NewBeatController(model)
	
	view := NewBasicDJView(model)
	view.SetController(controller)
	controller.SetView(view)

	view.CreateWindow("DJ View")
	view.Show()
}