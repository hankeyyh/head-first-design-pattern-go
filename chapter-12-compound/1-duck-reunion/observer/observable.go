package observer

type Observable struct {
	observerList []Observer
	// duck QuackObservable
}

func (o *Observable) RegisterObserver(observer Observer) {
	o.observerList = append(o.observerList, observer)
}

func (o *Observable) NotifyObservers() {
	for _, observer := range o.observerList {
		observer.Update(o)
	}
}