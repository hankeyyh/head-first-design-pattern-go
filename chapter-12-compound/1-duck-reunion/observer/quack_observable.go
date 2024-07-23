package observer

type QuackObservable interface {
	RegisterObserver(observer Observer)
	NotifyObservers()
}