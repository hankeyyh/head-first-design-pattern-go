package observer

type Observer interface {
	Update(duck QuackObservable)
}
