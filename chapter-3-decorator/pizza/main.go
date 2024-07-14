package main

import "fmt"

func main() {
	thickPizza := NewThickcrustPizza()
	thinPizza := NewThincrustPizza()

	finalThickPizza := NewOlives(NewCheese(thickPizza))
	finalThinPizza := NewOlives(NewCheese(thinPizza))
	
	fmt.Println(finalThickPizza.GetDescription(), finalThickPizza.Cost())
	fmt.Println(finalThinPizza.GetDescription(), finalThinPizza.Cost())
}