package main

func main() {
	waitress := NewWaitress(NewPancakeHouseMenu(), NewDinerMenu())
	waitress.printMenu()
	waitress.printVegetarianMenu()
}