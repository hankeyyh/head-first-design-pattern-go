package main

type Controller struct {
	model *Model
	view *View
	running bool
}

func NewController(model *Model) *Controller {
	return &Controller{model: model}
}

func (c *Controller) SetView(view *View) {
	c.view = view
}

func (c *Controller) Start() {
	if c.running {
		c.view.DisplayText("Music is already running\n")
		return
	}
	c.view.DisplayText("Music now running\n")
	c.model.On()
	c.running = true
}

func (c *Controller) Stop() {
	if !c.running {
		c.view.DisplayText("Music is already stopped\n")
		return
	}
	c.running = false
	c.view.DisplayText("Music now stopped\n")
}

func (c *Controller) Quit() {
	c.model.Off()
	c.running = false
	c.view.DisplayText("Party Over\n")
}

func (c *Controller) SetBpm(bpm int) {
	c.model.SetBpm(bpm)
}

func (c *Controller) IncreaseBpm(delta int) {
	c.model.SetBpm(c.model.GetBpm() + delta)
}