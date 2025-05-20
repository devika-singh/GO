package bill

import (
	"fmt"
	"os"
	"time"
)

type Bill struct {
	name       string
	tableNo    int
	dateTime   string
	serverName string
	itemsAmt   map[string]int
	itemsPrice map[string]float64
	total      float64
	tip        float64
}

func (b *Bill) CreateBill(name string, tableNo int, serverName string) {
	b.name = name
	b.tableNo = tableNo
	b.dateTime = time.Now().Format("2006-01-02 15:04:05")
	b.serverName = serverName
	b.itemsAmt = map[string]int{}
	b.itemsPrice = map[string]float64{}
	b.total = 0.0
	b.tip = 0.0
}

func (b *Bill) AddItem(itemName string, price float64, quantity int) {
	b.itemsAmt[itemName] += quantity
	b.itemsPrice[itemName] = price
	b.total += price * float64(quantity)
}

func (b *Bill) RemoveItem(itemName string, price float64, quantity int) {
	if _, exists := b.itemsAmt[itemName]; exists {
		b.itemsAmt[itemName] -= quantity
		if b.itemsAmt[itemName] <= 0 {
			delete(b.itemsAmt, itemName)
		}
		b.total -= price * float64(quantity)
	}
}

func (b *Bill) UpdateTip(tip float64) {
	b.tip = tip
}

func (b *Bill) CalculateTotal() {
	b.total += (b.tip / 100.00) * b.total
}

func (b *Bill) GenerateBill() string {
	billDetails := "Bill Details:\n"
	billDetails += fmt.Sprintf("%v %v\n", "Name: ", b.name)
	billDetails += fmt.Sprintf("%v %v\n", "Server:", b.serverName)
	billDetails += fmt.Sprintf("%v %v\n", "Table No:", b.tableNo)
	billDetails += fmt.Sprintf("%v %v\n", "Date & Time:", b.dateTime)
	billDetails += fmt.Sprintf("%v %v\n", "Server Name:", b.serverName)
	billDetails += fmt.Sprintf("%-25v %-9v %v\n", "Items", "Qty", "Price")
	billDetails += fmt.Sprintf("%-25v%v\n", "-------------------------------------", "----")

	for item, quantity := range b.itemsAmt {
		billDetails += fmt.Sprintf("%-25v %-6v ...$%0.2f\n", item+": ", quantity, float64(quantity)*b.itemsPrice[item])
	}

	billDetails += fmt.Sprintf("%-24v %%%-5v ...$%0.2f\n", "Tip:", b.tip, (b.tip/100.00)*b.total)
	b.CalculateTotal()
	billDetails += fmt.Sprintf("%-32v ...$%0.2f\n", "Total:", b.total)

	return billDetails
}

func (b *Bill) SaveBill() {
	data := []byte(b.GenerateBill())
	err := os.WriteFile("../../bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
}
