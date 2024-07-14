package main

import "fmt"


type OwnerProxy struct {
	person Person
}

func (op *OwnerProxy) GetName() string {
	return op.person.GetName()
}

func (op *OwnerProxy) GetGender() string {
	return op.person.GetGender()
}

func (op *OwnerProxy) GetInterests() string {
	return op.person.GetInterests()
}

func (op *OwnerProxy) GetGeekRating() int {
	return op.person.GetGeekRating()
}

func (op *OwnerProxy) SetName(name string) {
	op.person.SetName(name)
}

func (op *OwnerProxy) SetGender(gender string) {
	op.person.SetGender(gender)
}

func (op *OwnerProxy) SetInterests(interests string) {
	op.person.SetInterests(interests)
}

func (op *OwnerProxy) SetGeekRating(rating int) {
	fmt.Println("You can't set geek rating for yourself")
}

type NonOwnerProxy struct {
	person Person
}

func (op *NonOwnerProxy) GetName() string {
	return op.person.GetName()
}

func (op *NonOwnerProxy) GetGender() string {
	return op.person.GetGender()
}

func (op *NonOwnerProxy) GetInterests() string {
	return op.person.GetInterests()
}

func (op *NonOwnerProxy) GetGeekRating() int {
	return op.person.GetGeekRating()
}

func (op *NonOwnerProxy) SetName(name string) {
	fmt.Println("You can't set name for others")
}

func (op *NonOwnerProxy) SetGender(gender string) {
	fmt.Println("you can't set gender for others")
}

func (op *NonOwnerProxy) SetInterests(interests string) {
	fmt.Println("You can't set interests for others")
}

func (op *NonOwnerProxy) SetGeekRating(rating int) {
	op.person.SetGeekRating(rating)
}