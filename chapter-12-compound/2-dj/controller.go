package main

type Controller interface {
	Start()
	Stop()
	Quit()
	SetView(view DJView)
	IncreaseBpm()
	DecreaseBpm()
	SetBpm(bpm int)
}

type BeatController struct {
	model BeatModel
	view  DJView
}

func NewBeatController(model BeatModel) *BeatController {
	result := &BeatController{
		model: model,
	}
	return result
}

func (b *BeatController) SetView(view DJView) {
	b.view = view
}

func (b *BeatController) Start() {
	b.model.On()
	b.view.DisableStartMenuItem()
	b.view.EnableStopMenuItem()
}

func (b *BeatController) Stop() {
	b.model.Off()
	b.view.DisableStopMenuItem()
	b.view.EnableStartMenuItem()
}

func (b *BeatController) Quit() {
	b.model.Quit()
	b.view.Quit()
}

func (b *BeatController) IncreaseBpm() {
	b.model.SetBpm(b.model.GetBpm() + 1)
}

func (b *BeatController) DecreaseBpm() {
	b.model.SetBpm(b.model.GetBpm() - 1)
}

func (b *BeatController) SetBpm(bpm int) {
	b.model.SetBpm(bpm)
}
