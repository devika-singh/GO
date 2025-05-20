package menu

import "fmt"

type MenuItem struct {
	Name   string
	Price  float64
	Status int
}

type Menu struct {
	itemList map[string]MenuItem
}

func (m *Menu) AddItem(name string, price float64) {
	if m.itemList == nil { // lazy init works too
		m.itemList = make(map[string]MenuItem)
	}
	fmt.Println("Before Add item ", m.itemList)
	m.itemList[name] = MenuItem{Name: name, Price: price, Status: 1}
	fmt.Println("After Add item ", m.itemList)
}

func (m *Menu) UpdateStatus(name string) {
	fmt.Println("Update Status ", name)
	//86 the status
}

func (m *Menu) RemoveItem(name string) {

	fmt.Println("Remove item ", name)
	//remove the item from the menu

}
