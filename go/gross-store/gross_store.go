package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, ok := units[unit]
	if !ok {
		return false
	}

	billQuantity, billCheckQuantity := bill[item]
	if billCheckQuantity {
		bill[item] = billQuantity + quantity
	} else {
		bill[item] = quantity
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	existingQuantity, checkItem := bill[item]
	if !checkItem {
		return false
	}

	quantity, checkUnit := units[unit]
	if !checkUnit {
		return false
	}

	if existingQuantity < quantity {
		return false
	}

	if existingQuantity == quantity {
		delete(bill, item)
	} else {
		bill[item] = existingQuantity - quantity
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, ok := bill[item]
	return quantity, ok
}
