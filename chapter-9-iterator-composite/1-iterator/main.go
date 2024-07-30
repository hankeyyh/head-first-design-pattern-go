package main

/*
different menu share the same interface "iterator"
*/
func main() {
	waitress := NewWaitress(NewPancakeHouseMenu(), NewDinerMenu(), NewCafeMenu())
	waitress.printMenu()
	waitress.printVegetarianMenu()
}