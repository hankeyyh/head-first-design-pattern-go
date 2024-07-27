package viewcomponent

import (
	"math"
	"time"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type BeatBar struct {
	widget.ProgressBar
	value binding.ExternalFloat

	close chan struct{}
	done chan struct{}
}

func NewBeatBar() *BeatBar {
	b := &BeatBar{
		close: make(chan struct{}),
		done: make(chan struct{}),
	}
	b.ExtendBaseWidget(b)

	f := 0.0
	b.value = binding.BindFloat(&f)
	b.Bind(b.value)

	go b.decay()
	
	return b
}

func (b *BeatBar) decay() {
	defer close(b.done)
	for {
		raw, _ := b.value.Get()
		raw = math.Max(raw * 0.75, 0.0)

		b.value.Set(raw)
		time.Sleep(50 * time.Millisecond)

		select {
		case <-b.close:
			return
		default:
		}
	}
}

func (b *BeatBar) Pulse() {
	b.value.Set(1.0)
}

func (b *BeatBar) Close() {
	close(b.close)
	<-b.done
}