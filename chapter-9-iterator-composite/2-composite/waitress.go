package main

type Waitress struct {
	menus []MenuComponment
}

func NewWaitress(menu ...MenuComponment) *Waitress {
	result := &Waitress{}
	result.menus = append(result.menus, menu...)
	return result
}

func (w *Waitress) Print() {
	for _, menu := range w.menus {
		menu.Print()
	}
}

func (w *Waitress) Add(child MenuComponment) {
	w.menus = append(w.menus, child)
}

func (w *Waitress) Remove(child MenuComponment) {
	for i, v := range w.menus {
		if v == child {
			w.menus = append(w.menus[:i], w.menus[i+1:]...)
			break
		}
	}
}

func (w *Waitress) GetChild(n int) MenuComponment {
	return w.menus[n]
}
