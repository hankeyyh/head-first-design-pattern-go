package main

import (
	"sync"
	"time"
)

type BeatModel interface {
	Initialize() error
	On()
	Off()
	Quit() error

	SetBpm(bpm int)
	GetBpm() int

	// notified on every beat
	RegisterBeatObserver(o BeatObserver)
	RemoveBeatObserver(o BeatObserver)
	// notified when BPM changes
	RegisterBpmObserver(o BpmObserver)
	RemoveBpmObserver(o BpmObserver)
}

type BasicBeatModel struct {
	bpm int
	bpmLock sync.RWMutex

	stop bool

	beatObservers map[string]BeatObserver
	bpmObservers  map[string]BpmObserver

	audio PlayAudioModel

	stopChan chan struct{}
	doneChan chan struct{}
}

func NewBasicBeatModel() *BasicBeatModel {
	return &BasicBeatModel{
		beatObservers: make(map[string]BeatObserver),
		bpmObservers: make(map[string]BpmObserver),
		stop: true,
	}
}

func (b *BasicBeatModel) Initialize() error {
	b.audio = NewBeepAudioModel()
	return b.audio.Initialize("clap.mp3")
}

func (b *BasicBeatModel) On() {
	if !b.stop {
		return
	}
	b.SetBpm(90)
	b.stop = false
	b.stopChan = make(chan struct{})
	b.doneChan = make(chan struct{})

	go b.loop()
}

func (b *BasicBeatModel) Off() {
	if b.stop {
		return
	}
	b.stop = true
	close(b.stopChan)
	<-b.doneChan
}

func (b *BasicBeatModel) Quit() error {
	b.Off()
	return b.audio.Quit()
}

func (b *BasicBeatModel) loop() {
	defer close(b.doneChan)
	for {
		b.audio.Play()
		b.audio.Seek(0)
		b.notifyBeatObservers()

		waitDuration := time.Duration(60000/b.GetBpm()) * time.Millisecond
		time.Sleep(waitDuration)

		select {
		case <-b.stopChan:
			return
		default:
		}
	}
}

func (b *BasicBeatModel) SetBpm(bpm int) {
	b.bpmLock.Lock()
	b.bpm = bpm
	b.bpmLock.Unlock()
	b.notifyBpmObservers()
}

func (b *BasicBeatModel) GetBpm() int {
	b.bpmLock.RLock()
	defer b.bpmLock.RUnlock()
	return b.bpm
}

func (b *BasicBeatModel) RegisterBeatObserver(o BeatObserver) {
	b.beatObservers[o.GetBeatObserverName()] = o
}

func (b *BasicBeatModel) RemoveBeatObserver(o BeatObserver) {
	delete(b.beatObservers, o.GetBeatObserverName())
}

func (b *BasicBeatModel) notifyBeatObservers() {
	for _, o := range b.beatObservers {
		o.UpdateBeat()
	}
}

func (b *BasicBeatModel) RegisterBpmObserver(o BpmObserver) {
	b.bpmObservers[o.GetBpmObserverName()] = o
}

func (b *BasicBeatModel) RemoveBpmObserver(o BpmObserver) {
	delete(b.bpmObservers, o.GetBpmObserverName())
}

func (b *BasicBeatModel) notifyBpmObservers() {
	for _, o := range b.bpmObservers {
		o.UpdateBpm(b.bpm)
	}
}