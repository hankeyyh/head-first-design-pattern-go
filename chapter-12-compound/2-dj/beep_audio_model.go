package main

import (
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

type PlayAudioModel interface {
	Initialize(filename string) error
	Play() error
	Seek(n int) error
	Quit() error
}

type BeepAudioModel struct {
	streamer beep.StreamSeekCloser
	format beep.Format
}
func NewBeepAudioModel() *BeepAudioModel {
	return &BeepAudioModel{}
}

func (b *BeepAudioModel) Initialize(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	b.streamer, b.format, err = mp3.Decode(f)
	if err != nil {
		return err
	}

	err = speaker.Init(b.format.SampleRate, b.format.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}

	return nil
}

func (b *BeepAudioModel) Play() error {
	speaker.PlayAndWait(b.streamer)
	return nil
}

func (b *BeepAudioModel) Seek(n int) error {
	return b.streamer.Seek(n)
}

func (b *BeepAudioModel) Quit() error {
	b.streamer.Close()
	return nil
}
