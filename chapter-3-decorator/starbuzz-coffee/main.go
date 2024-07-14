package main

import (
	"fmt"
)

/*
Beverage 必须是interface, 不然condiment没办法嵌套各种beverage
condiment 需要调用 beverage.getSize 决定了interface中要包含getSize方法
beverage interface 不需要包含setSize, 不然condiment也可以调setsize有点奇怪,
setSize可以放在各种concrete beverage里
*/
func main() {
	espresso := NewEspresso()
	fmt.Printf("%s, $%.2f\n", espresso.GetDescription(), espresso.Cost())

	darkRoast := NewDarkRoast()
	fmt.Printf("%s, $%.2f\n", darkRoast.GetDescription(), darkRoast.Cost())
	// TODO value receiver not suitable for embedded pointer receiver
	finalDarkRoast := NewWhip(NewMocha(NewMocha(&darkRoast)))
	fmt.Printf("%s, $%.2f\n", finalDarkRoast.GetDescription(), finalDarkRoast.Cost())

	houseBlend := NewHouseBlend()
	houseBlend.SetSize(VENTI)
	fmt.Printf("%s, $%.2f\n", houseBlend.GetDescription(), houseBlend.Cost())
	finalHouseBlend := NewSoy(NewMocha(NewWhip(&houseBlend)))
	fmt.Printf("%s, $%.2f\n", finalHouseBlend.GetDescription(), finalHouseBlend.Cost())
}