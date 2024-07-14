package main

type Screen struct {
	description string
}

func NewScreen(description string) *Screen {
	return &Screen{description: description}
}

func (s *Screen) Up() {
	println(s.description + " going up")
}

func (s *Screen) Down() {
	println(s.description + " going down")
}

func (s *Screen) ToString() string {
	return s.description
}