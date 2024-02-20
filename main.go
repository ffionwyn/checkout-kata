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

// checkout struct to represent the checkout system.
type Checkout struct {
    items map[string]Item
    cart  map[string]int
}

