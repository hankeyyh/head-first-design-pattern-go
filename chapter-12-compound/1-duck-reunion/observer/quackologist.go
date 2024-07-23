package observer

type Quackologist struct {

}

func NewQuackologist() *Quackologist {
	return &Quackologist{}
}

func (q *Quackologist) Update(duck QuackObservable) {
	println("Quackologist: ", duck, " just quacked.")
}