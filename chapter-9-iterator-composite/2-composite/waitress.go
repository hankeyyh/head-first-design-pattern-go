package main

type Waitress struct {
	allmenu MenuComponment
}

func NewWaitress(allmenu MenuComponment) *Waitress {
	result := &Waitress{
		allmenu: allmenu,
	}
	return result
}

func (w *Waitress) PrintMenu() {
	w.allmenu.Print()
}