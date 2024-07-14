package main

func main() {
	amp := NewAmplifier("Amplifier")
	tuner := NewTuner("AM/FM Tuner")
	screen := NewScreen("Theater Screen")

	homeTheater := NewHomeTheaterFacade(amp, tuner, screen)

	homeTheater.WatchMovie()
	homeTheater.EndMovie()
}