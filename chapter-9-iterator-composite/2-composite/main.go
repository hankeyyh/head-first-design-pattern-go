package main

/*
Multi-level menu
The Composite Pattern allows you to compose objects into tree structures 
to represent part-whole hierarchies. Composite lets clients tret
individual objects and compositions of objects uniformly.
*/
func main() {
	pancakeHouseMenu := NewBasicMenu("Pancake House Menu", "Breakfast")
	dinerMenu := NewBasicMenu("Diner Menu", "Lunch")
	cafeMenu := NewBasicMenu("Cafe Menu", "Dinner")
	dessertMenu := NewBasicMenu("Desert Menu", "Desert of course!")
	coffeeMenu := NewBasicMenu("Coffee Menu", "Stuff to go with your afternoon coffee")

	allMenus := NewBasicMenu("All Menus", "All menus combined")
	allMenus.Add(pancakeHouseMenu)
	allMenus.Add(dinerMenu)
	allMenus.Add(cafeMenu)

	pancakeHouseMenu.Add(NewMenuItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99))
	pancakeHouseMenu.Add(NewMenuItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99))
	pancakeHouseMenu.Add(NewMenuItem("Blueberry Pancakes", "Pancakes made with fresh blueberries, and blueberry syrup", true, 3.49))
	pancakeHouseMenu.Add(NewMenuItem("Waffles", "Waffles with your choice of blueberries or strawberries", true, 3.59))
	
	dinerMenu.Add(NewMenuItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99))
	dinerMenu.Add(NewMenuItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99))
	dinerMenu.Add(NewMenuItem("Soup of the day", "A bowl of the soup of the day, with a side of potato salad", false, 3.29))
	dinerMenu.Add(NewMenuItem("Hot Dog", "A hot dog, with saurkraut, relish, onions, topped with cheese", false, 3.05))
	dinerMenu.Add(NewMenuItem("Steamed Veggies and Brown Rice", "Steamed vegetables over brown rice", true, 3.99))
	dinerMenu.Add(NewMenuItem("Pasta", "Spaghetti with marinara sauce, and a slice of sourdough bread", true, 3.89))
	dinerMenu.Add(dessertMenu)

	dessertMenu.Add(NewMenuItem("Apple Pie", "Apple pie with a flakey crust, topped with vanilla icecream", true, 1.99))
	dessertMenu.Add(NewMenuItem("Cheesecake", "Creamy New York cheesecake, with a chocolate graham crust", true, 1.99))
	dessertMenu.Add(NewMenuItem("Sorbet", "A scoop of raspberry and a scoop of lime", true, 1.89))
	
	cafeMenu.Add(coffeeMenu)
	coffeeMenu.Add(NewMenuItem("Coffee Cake", "Crumbly cake topped with cinnamon and walnuts", true, 1.59))
	coffeeMenu.Add(NewMenuItem("Bagel", "Flavors include sesame, poppyseed, cinnamon raisin, pumpkin", false, 0.69))
	coffeeMenu.Add(NewMenuItem("Biscotti", "Three almond or hazelnut biscotti cookies", true, 0.89))

	waitress := NewWaitress(allMenus)
	waitress.PrintMenu()

}