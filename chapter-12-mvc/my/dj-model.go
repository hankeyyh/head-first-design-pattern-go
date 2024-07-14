package main

type Model struct {
	bpm int
	observers []Observer
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) On() {
	m.bpm = 1
}

func (m *Model) Off() {
	m.bpm = 0
}

func (m *Model) SetBpm(bpm int) {
	m.bpm = bpm
	m.NotifyObservers()
}

func (m *Model) GetBpm() int {
	return m.bpm
}

func (m *Model) NotifyObservers() {
	for _, observer := range m.observers {
		observer.Update()
	}
}

func (m *Model) Register(o Observer) {
	m.observers = append(m.observers, o)
}