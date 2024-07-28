package mvc

type BeatObserver interface {
	GetBeatObserverName() string
	UpdateBeat()
}

type BpmObserver interface {
	GetBpmObserverName() string
	UpdateBpm(bpm int)
}
