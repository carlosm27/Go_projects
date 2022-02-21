package main

func findGrocery(groceries string) (*Grocery, error) {
	var grocery Grocery
	if result := db.Where(&Grocery{Name: groceries}).First(&grocery); result.Error != nil {
		return nil, result.Error
	}
	return &grocery, nil
}

func findQuantity(quantity int) ([]Grocery, error) {
	var groceries []Grocery
	if result := db.Where("Quantity > ?", quantity).Find(&groceries); result.Error != nil {
		return nil, result.Error
	}
	return groceries, nil
}

func insertGrocery(grocery *Grocery) error {
	if result := db.Create(grocery); result.Error != nil {
		return result.Error
	}
	return nil
}
