package ingredient

type Cheese interface {
	ToCheeseString() string
}

type MozzarellaCheese struct {

}

func NewMozzarellaCheese() *MozzarellaCheese {
	return &MozzarellaCheese{}
}

func (m *MozzarellaCheese) ToCheeseString() string {
	return "Shredded Mozzarella";
}

type ParmesanCheese struct {

}

func NewParmesanCheese() *ParmesanCheese {
	return &ParmesanCheese{}
}

func (m *ParmesanCheese) toCheeseString() string {
	return "Shredded Parmesan";
}

type ReggianoCheese struct {

}

func NewReggianoCheese() *ReggianoCheese {
	return &ReggianoCheese{}
}

func (m *ReggianoCheese) ToCheeseString() string {
	return "Reggiano Cheese";
}