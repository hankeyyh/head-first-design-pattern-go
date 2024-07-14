package vendor

type Hothub struct {
	on bool
	temperature int
}

func NewHothub() *Hothub {
	return &Hothub{}
}

func (h *Hothub) On() {
	h.on = true
	println("Turn on hothub")
}

func (h *Hothub) Off() {
	h.on = false
	println("Turn off hothub")
}

func (h *Hothub) Circulate() {
	if h.on {
		println("hothub is bubbling!")
	}
}

func (h *Hothub) JetsOn() {
	if h.on {
		println("Hottub jets are on")
	}
}

func (h *Hothub) JetsOff() {
	if h.on {
		println("Hottub jets are off")
	}
}

func (h *Hothub) SetTemperature(temperature int) {
	if temperature > h.temperature {
		println("Hottub is heating to a steaming ", temperature, " degrees")
	} else {
		println("Hottub is cooling to ", temperature, " degrees")
	}
	h.temperature = temperature
}