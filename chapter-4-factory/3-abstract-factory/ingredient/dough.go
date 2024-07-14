package ingredient

type Dough interface {
	ToDoughString() string
}

type ThinCrustDough struct{

}

func NewThinCrustDough() *ThinCrustDough {
	return &ThinCrustDough{}
}

func (t *ThinCrustDough) ToDoughString() string {
	return "Thin Crust Dough"
}

type ThickCrustDough struct{

}

func NewThickCrustDough() *ThickCrustDough {
	return &ThickCrustDough{}
}

func (t *ThickCrustDough) ToDoughString() string {
	return "Thick Crust Dough"
}
