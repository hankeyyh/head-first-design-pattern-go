package main

import (
	"testing"
	"time"
)

func TestBeepAudioModel(t *testing.T) {
	model := NewBeepAudioModel()
	if err := model.Initialize("clap.mp3"); err != nil {
		t.Fatal(err)
	}
	model.Play()
}

func TestModel(t *testing.T) {
	model := NewBasicBeatModel()
	if err := model.Initialize(); err != nil {
		t.Fatal(err)
	}
	model.On()
	<-time.After(5 * time.Second)
}
