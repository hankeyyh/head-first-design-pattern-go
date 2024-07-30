package main

/*
Multi-level menu
The Composite Pattern allows you to compose objects into tree structures 
to represent part-whole hierarchies. Composite lets clients tret
individual objects and compositions of objects uniformly.
*/
func main() {
	waitress := NewWaitress(NewPancakeHouseMenu(), NewDinerMenu(), NewCafeMenu())
	waitress.Print()
}