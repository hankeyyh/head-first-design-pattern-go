package main

import "github.com/hankeyyh/head-first-design-pattern-go/chapter-12-compound/2-dj/mvc"

type HeartAdapter struct {
	heart HeartModel
}

func NewHeartAdapter(heart HeartModel) *HeartAdapter {
	return &HeartAdapter{
		heart: heart,
	}
}

func (h *HeartAdapter) Initialize() error {
	return nil
}

func (h *HeartAdapter) On() {

}

func (h *HeartAdapter) Off() {

}

func (h *HeartAdapter) Quit() error {
	return nil
}

func (h *HeartAdapter) SetBpm(bpm int) {

}

func (h *HeartAdapter) GetBpm() int {
	return h.heart.GetHeartRate()
}

// notified on every beat
func (h *HeartAdapter) RegisterBeatObserver(o mvc.BeatObserver) {
	h.heart.RegisterBeatObserver(o)
}

func (h *HeartAdapter) RemoveBeatObserver(o mvc.BeatObserver) {
	h.heart.RemoveBeatObserver(o)
}

// notified when BPM changes
func (h *HeartAdapter) RegisterBpmObserver(o mvc.BpmObserver) {
	h.heart.RegisterBpmObserver(o)
}

func (h *HeartAdapter) RemoveBpmObserver(o mvc.BpmObserver) {
	h.heart.RemoveBpmObserver(o)
}
