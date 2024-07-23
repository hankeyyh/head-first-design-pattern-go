package main

import "fmt"

type DJView struct {
	model BeatModelInterface
	controller ControllerInterface
}

func NewDJView(controller ControllerInterface, model BeatModelInterface) *DJView {
	result := &DJView{controller: controller, model: model}
	model.RegisterBeatObserver("view", result)
	model.RegisterBpmObserver("view", result)

	return result
}

func (v *DJView) createView() {
	// TODO 创建gui
}

// UpdateBeat 当model每次节拍时调用
func (v *DJView) UpdateBeat() {
	// TODO beatBar set value
	fmt.Println("beat")
}

// UpdateBpm 当model的bpm改变时调用
func (v *DJView) UpdateBpm() {
	bpm := v.model.GetBpm()
	if bpm == 0 {
		// TODO offline
		fmt.Println("offline")
	} else {
		// TODO 显示bpm
		fmt.Println("Current bpm: ", bpm)
	}
}

func (v *DJView) createControls() {
	// TODO 创建控制菜单
}

func (v *DJView) enableStopMenuItem() {
	//TODO 开启停止按钮
}

func (v *DJView) disableStopMenuItem() {
	//TODO 关闭停止按钮
}

func (v *DJView) enableStartMenuItem() {
	//TODO 开启开始按钮
}

func (v *DJView) disableStartMenuItem() {
	//TODO 关闭开始按钮
}

func (v *DJView) actionPerformed() {
	// TODO 设置,增加,减少 bpm
}