package main

/*
view: 显示当前bpm, 输入bpm, 设置, 增加/减少
control:
model: bpm

*/

func main() {
	model := NewModel()
	controller := NewController(model)
	view := NewView(model, controller)
	controller.SetView(view)
	model.Register(view)

	// 正常
	view.Start()
	view.SetBpm(10)
	view.Show()
	view.IncreaseBpm()
	view.Show()
	view.DecreaseBpm()
	view.Show()
	view.Quit()

	// 重复关闭
	view.Stop()

	// 重复打开
	view.Start()
	view.Start()
	view.Stop()
	view.Quit()
}