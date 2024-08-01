package main

import (
	"fmt"
)

func main() {
	espresso := NewEspresso()
	fmt.Printf("%s, $%.2f\n", espresso.GetDescription(), espresso.Cost())

	darkRoast := NewDarkRoast()
	fmt.Printf("%s, $%.2f\n", darkRoast.GetDescription(), darkRoast.Cost())
	finalDarkRoast := NewWhip(NewMocha(NewMocha(&darkRoast)))
	fmt.Printf("%s, $%.2f\n", finalDarkRoast.GetDescription(), finalDarkRoast.Cost())

	houseBlend := NewHouseBlend()
	houseBlend.SetSize(VENTI)
	fmt.Printf("%s, $%.2f\n", houseBlend.GetDescription(), houseBlend.Cost())
	finalHouseBlend := NewSoy(NewMocha(NewWhip(&houseBlend)))
	fmt.Printf("%s, $%.2f\n", finalHouseBlend.GetDescription(), finalHouseBlend.Cost())
}