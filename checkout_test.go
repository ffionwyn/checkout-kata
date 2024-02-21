package main

import (
	"testing"
)

func TestScan(t *testing.T) {
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

    // case 2: scan a non-existing item.
	err = checkout.Scan("X")
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
