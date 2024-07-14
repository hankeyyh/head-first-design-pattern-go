package main

type HomeTheaterFacade struct {
	amp *Amplifier
	tuner *Tuner
	screen *Screen
}

func NewHomeTheaterFacade(amp *Amplifier, tuner *Tuner, screen *Screen) *HomeTheaterFacade {
	return &HomeTheaterFacade{amp: amp, tuner: tuner, screen: screen}
}

func (h *HomeTheaterFacade) WatchMovie() {
	println("Get ready to watch a movie...")
	h.screen.Down()
	h.amp.On()
	h.amp.SetSurroundSound()
	h.amp.SetVolume(5)
	h.tuner.On()
	h.tuner.SetAm()
}

func (h *HomeTheaterFacade) EndMovie() {
	println("Shutting movie theater down...")
	h.screen.Up()
	h.amp.Off()
	h.tuner.Off()
}