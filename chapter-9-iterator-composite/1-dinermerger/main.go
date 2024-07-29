package main

func main() {
	waitress := NewWaitress(NewPancakeHouseMenu(), NewDinerMenu(), NewCafeMenu())
	waitress.printMenu()
	waitress.printVegetarianMenu()
}