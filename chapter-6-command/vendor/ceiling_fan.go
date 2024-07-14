package vendor

type CeilingFanLevel int

const (
	Off CeilingFanLevel = iota
	Low
	Medium
	High
)

type CeilingFan struct {
	location string
	level CeilingFanLevel
}

func NewCeilingFan(location string) *CeilingFan {
	return &CeilingFan{location: location, level: Off}
}

func (c *CeilingFan) High() {
	c.level = High
	println(c.location, "ceiling fan is on high")
}

func (c *CeilingFan) Medium() {
	c.level = Medium
	println(c.location, "ceiling fan is on medium")
}

func (c *CeilingFan) Low() {
	c.level = Low
	println(c.location, "ceiling fan is on low")
}

func (c *CeilingFan) Off() {
	c.level = Off
	println(c.location, "ceiling fan is off")
}

func (c *CeilingFan) GetSpeed() CeilingFanLevel {
	return c.level
}