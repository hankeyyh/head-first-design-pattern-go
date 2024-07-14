package main

import "fmt"

func main() {
	person := &PersonImp{
		name: "Alice",
		gender: "female",
		interests: "programming",
	}

	ownerProxy := &OwnerProxy{person: person}
	nonOwnerProxy := &NonOwnerProxy{person: person}

	fmt.Println(ownerProxy.GetName())
	fmt.Println(ownerProxy.GetGender())
	fmt.Println(ownerProxy.GetInterests())
	ownerProxy.SetGeekRating(100)

	fmt.Println(nonOwnerProxy.GetName())
	fmt.Println(nonOwnerProxy.GetGender())
	fmt.Println(nonOwnerProxy.GetInterests())
	nonOwnerProxy.SetName("bob")
	nonOwnerProxy.SetGeekRating(200)

	fmt.Println(person.GetGeekRating())
}