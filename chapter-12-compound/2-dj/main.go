package main

func main() {
	model := NewBasicBeatModel()
	model.Initialize()
	
	view := NewBasicDJView(model)
	view.CreateWindow("DJ View")
	view.Show()
}