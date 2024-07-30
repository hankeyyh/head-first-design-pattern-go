package main

import "container/list"

type Iterator interface {
	Next() MenuComponment
	HasNext() bool
	Remove(c MenuComponment)
}

type DinerMenuIterator struct {
	menuItems []MenuComponment
	position  int
}

func NewDinerMenuIterator(menuItems []MenuComponment) *DinerMenuIterator {
	return &DinerMenuIterator{menuItems, 0}
}

func (d *DinerMenuIterator) Next() MenuComponment {
	menuItem := d.menuItems[d.position]
	d.position++
	return menuItem
}

func (d *DinerMenuIterator) HasNext() bool {
	if d.position >= len(d.menuItems) || d.menuItems[d.position] == nil {
		return false
	}
	return true
}

func (d *DinerMenuIterator) Remove(c MenuComponment) {
	for i, v := range d.menuItems {
		if v == c {
			d.menuItems = append(d.menuItems[:i], d.menuItems[i+1:]...)
			break
		}
	}
}

// New!!!!
type DessertMenuIterator struct {
	menuItems []MenuComponment
	position  int
}

func NewDessertMenuIterator(menuItems []MenuComponment) *DessertMenuIterator {
	return &DessertMenuIterator{menuItems, 0}
}

func (d *DessertMenuIterator) Next() MenuComponment {
	menuItem := d.menuItems[d.position]
	d.position++
	return menuItem
}

func (d *DessertMenuIterator) HasNext() bool {
	if d.position >= len(d.menuItems) || d.menuItems[d.position] == nil {
		return false
	}
	return true
}

func (d *DessertMenuIterator) Remove(c MenuComponment) {
	for i, v := range d.menuItems {
		if v == c {
			d.menuItems = append(d.menuItems[:i], d.menuItems[i+1:]...)
			break
		}
	}
}

type PancakeHouseMenuIterator struct {
	menuItems *list.List
	cur       *list.Element
}

func NewPancakeHouseMenuIterator(menuItems *list.List) *PancakeHouseMenuIterator {
	return &PancakeHouseMenuIterator{
		menuItems: menuItems,
	}
}

func (p *PancakeHouseMenuIterator) Next() MenuComponment {
	if p.cur == nil {
		p.cur = p.menuItems.Front()
	} else {
		p.cur = p.cur.Next()
	}
	return p.cur.Value.(MenuComponment)
}

func (p *PancakeHouseMenuIterator) HasNext() bool {
	if p.cur == nil {
		return p.menuItems.Front() != nil
	}
	return p.cur.Next() != nil
}

func (p *PancakeHouseMenuIterator) Remove(c MenuComponment) {
	for e := p.menuItems.Front(); e != nil; e = e.Next() {
		if e.Value == c {
			p.menuItems.Remove(e)
			break
		}
	}
}

type CafeMenuIterator struct {
	menuItems []MenuComponment
	position  int
}

func NewCafeMenuIterator(menuItems map[string]MenuComponment) *CafeMenuIterator {
	items := make([]MenuComponment, 0, len(menuItems))
	for _, item := range menuItems {
		items = append(items, item)
	}
	return &CafeMenuIterator{items, 0}
}

func (c *CafeMenuIterator) Next() MenuComponment {
	menuItem := c.menuItems[c.position]
	c.position++
	return menuItem
}

func (c *CafeMenuIterator) HasNext() bool {
	if c.position >= len(c.menuItems) || c.menuItems[c.position] == nil {
		return false
	}
	return true
}

func (c *CafeMenuIterator) Remove(componment MenuComponment) {
	for i, v := range c.menuItems {
		if v == componment {
			c.menuItems = append(c.menuItems[:i], c.menuItems[i+1:]...)
			break
		}
	}
}
