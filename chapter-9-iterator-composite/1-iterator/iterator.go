package main

import "container/list"

type Iterator interface {
	Next() *MenuItem
	HasNext() bool
}

type DinerMenuIterator struct {
	menuItems []*MenuItem
	position int
}

func NewDinerMenuIterator(menuItems []*MenuItem) *DinerMenuIterator {
	return &DinerMenuIterator{menuItems, 0}
}

func (d *DinerMenuIterator) Next() *MenuItem {
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

type PancakeHouseMenuIterator struct {
	menuItems *list.List
	cur *list.Element
}

func NewPancakeHouseMenuIterator(menuItems *list.List) *PancakeHouseMenuIterator {
	return &PancakeHouseMenuIterator{
		menuItems: menuItems,
	}
}

func (p *PancakeHouseMenuIterator) Next() *MenuItem {
	if p.cur == nil {
		p.cur = p.menuItems.Front()
	} else {
		p.cur = p.cur.Next()
	}
	return p.cur.Value.(*MenuItem)
}

func (p *PancakeHouseMenuIterator) HasNext() bool {
	if p.cur == nil {
		return p.menuItems.Front() != nil
	}
	return p.cur.Next() != nil
}

type CafeMenuIterator struct {
	menuItems []*MenuItem
	position int
}

func NewCafeMenuIterator(menuItems map[string]*MenuItem) *CafeMenuIterator {
	items := make([]*MenuItem, 0, len(menuItems))
	for _, item := range menuItems {
		items = append(items, item)
	}
	return &CafeMenuIterator{items, 0}
}

func (c *CafeMenuIterator) Next() *MenuItem {
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