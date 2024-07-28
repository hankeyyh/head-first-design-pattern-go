package main

import (
	"math/rand"
	"time"

	"github.com/hankeyyh/head-first-design-pattern-go/chapter-12-compound/2-dj/mvc"
)

type HeartModel interface {
	Start() 
	Quit() 
	GetHeartRate() int
	RegisterBeatObserver(o mvc.BeatObserver)
	RemoveBeatObserver(o mvc.BeatObserver)
	RegisterBpmObserver(o mvc.BpmObserver)
	RemoveBpmObserver(o mvc.BpmObserver)
}

type BasicHeartModel struct {
	time int

	beatObservers map[string]mvc.BeatObserver
	bpmObservers  map[string]mvc.BpmObserver

	close chan struct{}
	done chan struct{}
}

func NewHeartModel() *BasicHeartModel {
	return &BasicHeartModel{
		time: 1000,
		beatObservers: make(map[string]mvc.BeatObserver),
		bpmObservers:  make(map[string]mvc.BpmObserver),
		close: make(chan struct{}),
		done: make(chan struct{}),
	}
}

func (h *BasicHeartModel) Start() {
	go h.loop()
}

func (h *BasicHeartModel) Quit() {
	close(h.close)
	<-h.done
}

func (h *BasicHeartModel) loop() {
	defer close(h.done)
	lastRate := -1
	for {
		change := rand.Intn(10)
		if rand.Intn(2) == 0 {
			change = 0 - change
		}
		rate := 60000 / (h.time + change)
		if rate < 120 && rate > 50 {
			h.time += change
			h.notifyBeatObservers()
			if rate != lastRate {
				lastRate = rate
				h.notifyBpmObservers()
			}
		}
		time.Sleep(time.Duration(h.time) * time.Millisecond)

		select {
			case <-h.close:
				return
			default:
		}
	}
}

func (h *BasicHeartModel) GetHeartRate() int {
	return 60000 / h.time
}

// notified on every beat
func (h *BasicHeartModel) RegisterBeatObserver(o mvc.BeatObserver) {
	h.beatObservers[o.GetBeatObserverName()] = o
}

func (h *BasicHeartModel) RemoveBeatObserver(o mvc.BeatObserver) {
	delete(h.beatObservers, o.GetBeatObserverName())
}

func (h *BasicHeartModel) notifyBeatObservers() {
	for _, o := range h.beatObservers {
		o.UpdateBeat()
	}
}

// notified when BPM changes
func (h *BasicHeartModel) RegisterBpmObserver(o mvc.BpmObserver) {
	h.bpmObservers[o.GetBpmObserverName()] = o
}

func (h *BasicHeartModel) RemoveBpmObserver(o mvc.BpmObserver) {
	delete(h.bpmObservers, o.GetBpmObserverName())
}

func (h *BasicHeartModel) notifyBpmObservers() {
	for _, o := range h.bpmObservers {
		o.UpdateBpm()
	}
}