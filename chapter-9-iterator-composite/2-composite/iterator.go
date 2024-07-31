package main


type Iterator interface {
	Next() MenuComponment
	HasNext() bool
}

type BasicMenuIterator struct {
	menuComponments []MenuComponment
	position        int
}

func NewBasicMenuIterator(menuComponments []MenuComponment) *BasicMenuIterator {
	return &BasicMenuIterator{menuComponments, 0}
}

func (b *BasicMenuIterator) Next() MenuComponment {
	menuComponment := b.menuComponments[b.position]
	b.position++
	return menuComponment
}

func (b *BasicMenuIterator) HasNext() bool {
	if b.position >= len(b.menuComponments) || b.menuComponments[b.position] == nil {
		return false
	}
	return true
}
