package main

import (
	"container/list"
	"fmt"
)

type Menu interface {
	createIterator() Iterator
}

const DinerMenuMaxItems = 6

type DinerMenu struct {
	childs     []MenuComponment
	numofItems int
}

func NewDinerMenu() *DinerMenu {
	dm := &DinerMenu{
		childs:     make([]MenuComponment, DinerMenuMaxItems),
		numofItems: 0,
	}

	dm.Add(NewMenuItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99))
	dm.Add(NewMenuItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99))
	dm.Add(NewMenuItem("Soup of the day", "Soup of the day, with a side of potato salad", false, 3.29))
	dm.Add(NewMenuItem("Hotdog", "A hot dog, with saurkraut, relish, onions, topped with cheese", false, 3.05))
	dm.Add(NewDessertMenu())

	return dm
}

func (c *DinerMenu) Print() {
	fmt.Printf("\nLunch\n")
	iterator := c.createIterator()
	for iterator.HasNext() {
		iterator.Next().Print()
	}
}

func (c *DinerMenu) Add(child MenuComponment) {
	if c.numofItems >= DinerMenuMaxItems {
		panic("Sorry, menu is full! Can't add item to menu")
	} else {
		c.childs[c.numofItems] = child
		c.numofItems++
	}
}

func (c *DinerMenu) Remove(child MenuComponment) {
	for i, v := range c.childs {
		if v == child {
			c.childs = append(c.childs[:i], c.childs[i+1:]...)
			break
		}
	}
}

func (c *DinerMenu) GetChild(n int) MenuComponment {
	return c.childs[n]
}

func (c *DinerMenu) createIterator() Iterator {
	return NewDinerMenuIterator(c.childs)
}

// New!!!!
type DessertMenu struct {
	childs []MenuComponment
}

func NewDessertMenu() *DessertMenu {
	dm := &DessertMenu{
		childs: make([]MenuComponment, 0),
	}

	dm.Add(NewMenuItem("Apple Pie", "Apple pie with a flakey crust, topped with vanilla ice cream", true, 1.59))
	dm.Add(NewMenuItem("Cheesecake", "Creamy New York cheesecake, with a chocolate graham crust", true, 1.99))
	dm.Add(NewMenuItem("Sorbet", "A scoop of raspberry and a scoop of lime", true, 1.89))

	return dm
}

func (d *DessertMenu) Print() {
	fmt.Printf("\nDessert\n")
	iterator := d.createIterator()
	for iterator.HasNext() {
		iterator.Next().Print()
	}
}

func (d *DessertMenu) Add(child MenuComponment) {
	d.childs = append(d.childs, child)
}

func (d *DessertMenu) Remove(child MenuComponment) {
	for i, v := range d.childs {
		if v == child {
			d.childs = append(d.childs[:i], d.childs[i+1:]...)
			break
		}
	}
}

func (d *DessertMenu) GetChild(n int) MenuComponment {
	return d.childs[n]
}

func (c *DessertMenu) createIterator() Iterator {
	return NewDessertMenuIterator(c.childs)
}

type PancakeHouseMenu struct {
	childs *list.List
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	pm := &PancakeHouseMenu{
		childs: list.New(),
	}
	pm.Add(NewMenuItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99))
	pm.Add(NewMenuItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99))
	pm.Add(NewMenuItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", true, 3.49))
	pm.Add(NewMenuItem("Waffles", "Waffles, with your choice of blueberries or strawberries", true, 3.59))

	return pm
}

func (p *PancakeHouseMenu) Print() {
	fmt.Printf("Menu\n---\nBreakfast\n")
	iterator := p.createIterator()
	for iterator.HasNext() {
		iterator.Next().Print()
	}
}

func (p *PancakeHouseMenu) Add(child MenuComponment) {
	p.childs.PushBack(child)
}

func (p *PancakeHouseMenu) Remove(child MenuComponment) {
	for e := p.childs.Front(); e != nil; e = e.Next() {
		if e.Value == child {
			p.childs.Remove(e)
			break
		}
	}
}

func (p *PancakeHouseMenu) GetChild(n int) MenuComponment {
	i := 0
	for e := p.childs.Front(); e != nil; e = e.Next() {
		if i == n {
			return e.Value.(MenuComponment)
		}
		i++
	}
	return nil
}

func (p *PancakeHouseMenu) createIterator() Iterator {
	return NewPancakeHouseMenuIterator(p.childs)
}

type CafeMenu struct {
	childs map[string]MenuComponment
}

func NewCafeMenu() *CafeMenu {
	cm := &CafeMenu{
		childs: make(map[string]MenuComponment),
	}
	cm.addItem("Veggie Burger and Air Fries", "Veggie burger on a whole wheat bun, lettuce, tomato, and fries", true, 3.99)
	cm.addItem("Soup of the day", "A cup of the soup of the day, with a side salad", false, 3.69)
	cm.addItem("Burrito", "A large burrito, with whole pinto beans, salsa, guacamole", true, 4.29)
	return cm
}

func (c *CafeMenu) Print() {
	fmt.Printf("\nDinner\n")
	iterator := c.createIterator()
	for iterator.HasNext() {
		iterator.Next().Print()
	}
}

func (c *CafeMenu) Add(child MenuComponment) {

}

func (c *CafeMenu) Remove(child MenuComponment) {

}

func (c *CafeMenu) GetChild(n int) MenuComponment {
	return nil
}

func (c *CafeMenu) addItem(name, description string, vegetarian bool, price float64) {
	c.childs[name] = NewMenuItem(name, description, vegetarian, price)
}

func (c *CafeMenu) createIterator() Iterator {
	return NewCafeMenuIterator(c.childs)
}
