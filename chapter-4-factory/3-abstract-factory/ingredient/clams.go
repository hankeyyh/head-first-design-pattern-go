package ingredient

type Clams interface {
	ToClamsString() string
}

type FreshClams struct {

}

func NewFreshClams() *FreshClams {
	return &FreshClams{}
}

func (f *FreshClams) ToClamsString() string {
	return "Fresh Clams from Long Island Sound"
}

type FrozenClams struct {

}

func NewFrozenClams() *FrozenClams {
	return &FrozenClams{}
}

func (f *FrozenClams) ToClamsString() string {
	return "Frozen Clams from Chesapeake Bay"
}