package vendor

type Light struct {
	location string
}

func NewLight(location string) *Light {
	return &Light{location: location}
}

func (l *Light) TurnOn() {
	println("Turn on light")
}

func (l *Light) TurnOff(){
	println("Turn off light")
}