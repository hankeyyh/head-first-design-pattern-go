package main

import "container/list"

type Menu interface {
	createIterator() Iterator
}

const DinerMenuMaxItems = 6

type DinerMenu struct {
	menuItems []*MenuItem
	numofItems int
}

func NewDinerMenu() *DinerMenu {
	dm := &DinerMenu{}
	dm.menuItems = make([]*MenuItem, DinerMenuMaxItems)
	dm.numofItems = 0

	dm.addItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	dm.addItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)
	dm.addItem("Soup of the day", "Soup of the day, with a side of potato salad", false, 3.29)
	dm.addItem("Hotdog", "A hot dog, with saurkraut, relish, onions, topped with cheese", false, 3.05)
	return dm
}

func (d *DinerMenu) addItem(name, description string, vegetarian bool, price float64) {
	if d.numofItems >= DinerMenuMaxItems {
		panic("Sorry, menu is full! Can't add item to menu")
	} else {
		d.menuItems[d.numofItems] = NewMenuItem(name, description, vegetarian, price)
		d.numofItems++
	}
}

func (d *DinerMenu) GetMenuItems() []*MenuItem {
	return d.menuItems
}

func (d *DinerMenu) createIterator() Iterator {
	return NewDinerMenuIterator(d.menuItems)
}


type PancakeHouseMenu struct {
	menuItems *list.List
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	pm := &PancakeHouseMenu{
		menuItems: list.New(),
	}
	pm.addItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99)
	pm.addItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99)
	pm.addItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", true, 3.49)
	pm.addItem("Waffles", "Waffles, with your choice of blueberries or strawberries", true, 3.59)

	return pm
}

func (p *PancakeHouseMenu) addItem(name, description string, vegetarian bool, price float64) {
	p.menuItems.PushBack(NewMenuItem(name, description, vegetarian, price))
}

func (p *PancakeHouseMenu) GetMenuItems() *list.List {
	return p.menuItems
}

func (p *PancakeHouseMenu) createIterator() Iterator {
	return NewPancakeHouseMenuIterator(p.menuItems)
}

type CafeMenu struct {
	menuItems map[string]*MenuItem
}

func NewCafeMenu() *CafeMenu {
	cm := &CafeMenu{
		menuItems: make(map[string]*MenuItem),
	}
	cm.addItem("Veggie Burger and Air Fries", "Veggie burger on a whole wheat bun, lettuce, tomato, and fries", true, 3.99)
	cm.addItem("Soup of the day", "A cup of the soup of the day, with a side salad", false, 3.69)
	cm.addItem("Burrito", "A large burrito, with whole pinto beans, salsa, guacamole", true, 4.29)
	return cm
}

func (c *CafeMenu) addItem(name, description string, vegetarian bool, price float64) {
	c.menuItems[name] = NewMenuItem(name, description, vegetarian, price)
}

func (c *CafeMenu) createIterator() Iterator {
	return NewCafeMenuIterator(c.menuItems)
}