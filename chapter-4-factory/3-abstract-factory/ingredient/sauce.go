package ingredient

type Sauce interface {
	ToSauceString() string
}

type MarinaraSauce struct {

}

func NewMarinaraSauce() *MarinaraSauce {
	return &MarinaraSauce{}
}

func (m *MarinaraSauce) ToSauceString() string {
	return "Marinara Sauce"
}

type PlumTomatoSauce struct {

}

func NewPlumTomatoSauce() *PlumTomatoSauce {
	return &PlumTomatoSauce{}
}

func (p *PlumTomatoSauce) ToSauceString() string {
	return "Plum Tomato Sauce"
}