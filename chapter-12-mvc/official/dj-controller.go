package main

type ControllerInterface interface {
	Start()
	Stop()
	IncreaseBPM()
	DecreseBPM()
	SetBPM(bpm int)
}

type BeatController struct {
	model BeatModelInterface
	view *DJView
}

func NewBeatController(model BeatModelInterface) *BeatController {
	result := &BeatController{model: model}
	view := NewDJView(result, model)
	view.createView()
	view.createControls()
	view.disableStopMenuItem()
	view.enableStartMenuItem()
	model.Initialize()

	return result
}

func (c *BeatController) Start() {
	c.model.On()
	c.view.disableStartMenuItem()
	c.view.enableStopMenuItem()
}

func (c *BeatController) Stop() {
	c.model.Off()
	c.view.disableStopMenuItem()
	c.view.enableStartMenuItem()
}

func (c *BeatController) IncreaseBPM() {
	bpm := c.model.GetBpm()
	c.model.SetBpm(bpm + 1)
}

func (c *BeatController) DecreseBPM() {
	bpm := c.model.GetBpm()
	c.model.SetBpm(bpm - 1)
}

func (c *BeatController) SetBPM(bpm int) {
	c.model.SetBpm(bpm)
}