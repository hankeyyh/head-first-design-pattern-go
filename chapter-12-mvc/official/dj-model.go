package main

import (
	"time"
)

type BeatObserver interface {
	UpdateBeat()
}

type BpmObserver interface {
	UpdateBpm()
}

type BeatModelInterface interface {
	Initialize()
	On()
	Off()
	SetBpm(bpm int)
	GetBpm() int
	// notified on every beat
	RegisterBeatObserver(key string, o BeatObserver)
	RemoveBeatObserver(key string, o BeatObserver)
	// notified when bpm change
	RegisterBpmObserver(key string, o BpmObserver)
	RemoveBpmObserver(key string)
}

type Clip struct {
}

type BeatModel struct {
	bpm           int
	stop          bool
	clip          Clip
	beatObservers map[string]BeatObserver
	bpmObservers  map[string]BpmObserver
}

func NewBeatModel() *BeatModel {
	return &BeatModel{}
}

func (b *BeatModel) Initialize() {
	// audio file
}

func (b *BeatModel) On() {
	b.bpm = 90
	b.notifyBpmObservers()
	b.stop = false
	// goroutine play audio
	go b.loopBeat()
}

func (b *BeatModel) loopBeat() {
	for {
		b.playBeat()
		b.notifyBeatObservers()
		time.Sleep(time.Duration(float64(60) / float64(b.GetBpm())) * time.Second)
	}
}

func (b *BeatModel) Off() {
	b.stopBeat()
	b.stop = true
}

func (b *BeatModel) playBeat() {
	// audio play
}

func (b *BeatModel) stopBeat() {
	// audio stop
}

func (b *BeatModel) SetBpm(bpm int) {
	b.bpm = bpm
	b.notifyBpmObservers()
}

func (b *BeatModel) GetBpm() int {
	return b.bpm
}

func (b *BeatModel) notifyBeatObservers() {
	for _, o := range b.beatObservers {
		o.UpdateBeat()
	}
}

func (b *BeatModel) notifyBpmObservers() {
	for _, o := range b.bpmObservers {
		o.UpdateBpm()
	}
}

func (b *BeatModel) RegisterBeatObserver(key string, o BeatObserver) {
	b.beatObservers[key] = o
}

func (b *BeatModel) RemoveBeatObserver(key string, o BeatObserver) {
	delete(b.beatObservers, key)
}

func (b *BeatModel) RegisterBpmObserver(key string, o BpmObserver) {
	b.bpmObservers[key] = o
}

func (b *BeatModel) RemoveBpmObserver(key string) {
	delete(b.bpmObservers, key)
}