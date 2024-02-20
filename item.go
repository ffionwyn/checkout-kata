package main

// item struct to represent each item.
type Item struct {
	Name         string
	UnitPrice    int
	SpecialPrice SpecialPrice
}

// special price struct to represent special pricing for an item.
type SpecialPrice struct {
	Quantity int
	Price    int
}
