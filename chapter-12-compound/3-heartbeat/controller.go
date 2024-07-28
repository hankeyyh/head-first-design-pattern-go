package main

import "github.com/hankeyyh/head-first-design-pattern-go/chapter-12-compound/2-dj/mvc"

type HeartController struct {
	model HeartModel
	view  mvc.DJView
}

func NewHeatController(model HeartModel) *HeartController {
	result := &HeartController{
		model: model,
	}
	return result
}

func (c *HeartController) Start() {
	c.model.Start()
}

func (c *HeartController) Stop() {
	
}

func (c *HeartController) Quit() {
	c.model.Quit()
	c.view.Quit()
}

func (c *HeartController) SetView(view mvc.DJView) {
	c.view = view
}

func (c *HeartController) IncreaseBpm() {
	
}

func (c *HeartController) DecreaseBpm() {
	
}

func (c *HeartController) SetBpm(bpm int) {
	
}
