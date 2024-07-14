package ingredient

type Pepperoni interface {
	ToPepperoniString() string
}

type SlicedPepperoni struct {

}

func NewSlicedPepperoni() *SlicedPepperoni {
	return &SlicedPepperoni{}
}

func (s *SlicedPepperoni) ToPepperoniString() string {
	return "Sliced Pepperoni"
}