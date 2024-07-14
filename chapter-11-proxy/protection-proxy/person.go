package main

type Person interface {
	GetName() string
	GetGender() string
	GetInterests() string
	GetGeekRating() int

	SetName(name string)
	SetGender(gender string)
	SetInterests(interests string)
	SetGeekRating(rating int)
}

type PersonImp struct {
	name       string
	gender     string
	interests  string
	geekRating int
	ratingCnt  int
}

func (p *PersonImp) GetName() string {
	return p.name
}

func (p *PersonImp) GetGender() string {
	return p.gender
}

func (p *PersonImp) GetInterests() string {
	return p.interests
}

func (p *PersonImp) GetGeekRating() int {
	if (p.ratingCnt == 0) {
		return 0
	}
	return p.geekRating / p.ratingCnt
}

func (p *PersonImp) SetName(name string) {
	p.name = name
}

func (p *PersonImp) SetGender(gender string) {
	p.gender = gender
}

func (p *PersonImp) SetInterests(interests string) {
	p.interests = interests
}

func (p *PersonImp) SetGeekRating(rating int) {
	p.geekRating = rating
	p.ratingCnt++
}

