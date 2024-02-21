package main

import (
	"reflect"
	"testing"
)

func TestScanHappyPath(t *testing.T) {
    // map for testing.
    items := map[string]Item{
        "A": {
            Name:      "A",
            UnitPrice: 50,
            SpecialPrice: SpecialPrice{
                Quantity: 3,
                Price:    130,
            },
        },
    }

    // new checkout instance.
    checkout := StartCheckout(items)

    // case 1: scan an existing item
    err := checkout.Scan("A")
    if err != nil {
        t.Errorf("Error: expected nil, got: %v", err)
    }
}

func TestScanUnhappyPath(t *testing.T) {
    // map for testing.
    items := map[string]Item{
        "A": {
            Name:      "A",
            UnitPrice: 50,
            SpecialPrice: SpecialPrice{
                Quantity: 3,
                Price:    130,
            },
        },
    }

    // new checkout instance.
    checkout := StartCheckout(items)

    // case 2: scan a non-existing item.
	err := checkout.Scan("X")
    if err == nil {
        t.Errorf("Expected an error for scanning a non-existing item, but got nil")
    } else {
        t.Logf("Expected error for scanning a non-existing item: %v", err)
    }

    // case 3: scanning an item with an empty name.
    err = checkout.Scan("")
    if err == nil {
        t.Errorf("Expected an error for scanning an item with an empty name, but got nil")
    } else {
        t.Logf("Expected error for scanning an item with an empty name: %v", err)
    }
}

func TestGetTotalPriceHappyPath(t *testing.T) {
	// map for testing.
	items := map[string]Item{
		"A": {
			Name:      "A",
			UnitPrice: 50,
			SpecialPrice: SpecialPrice{
				Quantity: 3,
				Price:    130,
			},
		},
		"B": {
			Name:      "B",
			UnitPrice: 30,
			SpecialPrice: SpecialPrice{
				Quantity: 2,
				Price:    45,
			},
		},
		"C": {
			Name:      "C",
			UnitPrice: 20,
		},
		"D": {
			Name:      "D",
			UnitPrice: 15,
		},
	}
	checkout := StartCheckout(items)

	// case 1: no items scanned.
	totalPrice, err := checkout.GetTotalPrice()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if totalPrice != 0 {
		t.Errorf("Expected total price to be 0, got %d", totalPrice)
	}
}

func TestGetTotalPriceUnhappyPath(t *testing.T) {
	// map for testing.
	items := map[string]Item{
		"A": {
			Name:      "A",
			UnitPrice: 50,
			SpecialPrice: SpecialPrice{
				Quantity: 3,
				Price:    130,
			},
		},
		"B": {
			Name:      "B",
			UnitPrice: 30,
			SpecialPrice: SpecialPrice{
				Quantity: 2,
				Price:    45,
			},
		},
		"C": {
			Name:      "C",
			UnitPrice: 20,
		},
		"D": {
			Name:      "D",
			UnitPrice: 15,
		},
	}
	checkout := StartCheckout(items)

	// case 2: scan items and check total price. 
	checkout.Scan("A")
	checkout.Scan("A")
	checkout.Scan("B")
	checkout.Scan("C")
	totalPrice, err := checkout.GetTotalPrice()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedTotal := 50 + 50 + 30 + 20 
	if totalPrice != expectedTotal {
		t.Errorf("Expected total price to be %d, got %d", expectedTotal, totalPrice)
	}
}

func TestGetItemCountsHappyPath(t *testing.T) {
    // map for testing.
    items := map[string]Item{
        "A": {
            Name:      "A",
            UnitPrice: 50,
            SpecialPrice: SpecialPrice{
                Quantity: 3,
                Price:    130,
            },
        },
        "B": {
            Name:      "B",
            UnitPrice: 30,
            SpecialPrice: SpecialPrice{
                Quantity: 2,
                Price:    45,
            },
        },
    }

    checkout := StartCheckout(items)

    // case 1: when there are items in the cart.
    checkout.Scan("A")
    checkout.Scan("B")
    checkout.Scan("B")
    itemCounts := checkout.getItemCounts()
    expectedCounts := map[string]int{"A": 1, "B": 2}
    if !reflect.DeepEqual(itemCounts, expectedCounts) {
        t.Errorf("Expected item counts to be %v, but got %v", expectedCounts, itemCounts)
    }
}

func TestGetItemCountsUnhappyPath(t *testing.T) {
    // map for testing.
    items := map[string]Item{
        "A": {
            Name:      "A",
            UnitPrice: 50,
            SpecialPrice: SpecialPrice{
                Quantity: 3,
                Price:    130,
            },
        },
    }

    checkout := StartCheckout(items)

    // case 2: when there are no items in the cart.
    itemCountsEmpty := checkout.getItemCounts()
    if len(itemCountsEmpty) != 0 {
        t.Errorf("Expected empty item counts map, but got %v", itemCountsEmpty)
    }
}