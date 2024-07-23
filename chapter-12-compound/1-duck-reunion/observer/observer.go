package observer

import "reflect"

type Observer interface {
	Update(duck QuackObservable)
}

type Quackologist struct {

}

func NewQuackologist() *Quackologist {
	return &Quackologist{}
}

func (q *Quackologist) Update(duck QuackObservable) {
	t := reflect.ValueOf(duck).Elem().Type()
	println("Quackologist: ", t.Name(), " just quacked.")
}