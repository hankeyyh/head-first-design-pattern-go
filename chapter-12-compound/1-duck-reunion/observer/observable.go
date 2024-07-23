package observer

type BasicObservable struct {
	observerList []Observer
	Duck QuackObservable
}

func (o *BasicObservable) RegisterObserver(observer Observer) {
	o.observerList = append(o.observerList, observer)
}

func (o *BasicObservable) NotifyObservers() {
	for _, observer := range o.observerList {
		observer.Update(o.Duck)
	}
}
