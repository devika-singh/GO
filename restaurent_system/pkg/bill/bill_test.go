package bill

import (
	"strings"
	"testing"
)

func TestBillGeneration(t *testing.T) {
	var b Bill
	b.CreateBill("Alice", 5, "John")
	b.AddItem("Pasta", 12.50, 2)
	b.AddItem("Coke", 3.00, 1)
	b.AddItem("Salad", 5.00, 1)
	b.RemoveItem("Coke", 3.00, 1)
	b.UpdateTip(10)

	billText := b.GenerateBill()

	// Basic checks
	if !strings.Contains(billText, "Alice") {
		t.Errorf("Expected bill to contain customer name 'Alice'")
	}
	if !strings.Contains(billText, "Pasta") || !strings.Contains(billText, "2") {
		t.Errorf("Expected Pasta item with quantity 2")
	}
	if strings.Contains(billText, "Coke") {
		t.Errorf("Coke should have been removed")
	}
	if !strings.Contains(billText, "Total") || !strings.Contains(billText, "Tip") {
		t.Errorf("Bill missing total or tip")
	}

	b.SaveBill()
}
