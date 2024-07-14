package main

import (
	"fmt"
)

func main() {
	boiler := NewChocolateBoiler()
	boiler.Fill()
	boiler.Boil()
	boiler.Drain()
	fmt.Printf("boiler1: %p\n", boiler)

	boiler2 := NewChocolateBoiler()
	fmt.Printf("boiler2: %p\n", boiler2)
}
