package main

type TurkeyAdapter struct {
	turkey Turkey
}

func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey: turkey}
}

func (t *TurkeyAdapter) Quack() {
	t.turkey.Gobble()
}

func (t *TurkeyAdapter) Fly() {
	// Turkey fly shorter than duck, 
	// so fly the turkey five times to make it as long as duck
	for i := 0; i < 5; i++ {
		t.turkey.Fly()
	}
}