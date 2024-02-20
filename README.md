Checkout System

Code for a simple checkout system implemented in Go. The system allows users to scan items, calculates the total price, and applies special pricing for specific items.

Files

main.go: This file contains the main function that demonstrates the usage of the checkout system. It initialises items, scans them, calculates the total price, and prints the result.

checkout.go: This file defines the Checkout struct and related methods for managing the checkout process. It includes functionality to start a new checkout session, scan items, and calculate the total price.

item.go: This file defines the Item struct, representing each item available for checkout, along with its unit price and special pricing information.

Usage

To use the checkout system, follow these steps:

Run the main.go file to see the checkout system in action.

Functionality

Scanning Items: The Scan method in the Checkout struct allows users to add items to the cart.

Calculating Total Price: The GetTotalPrice method in the Checkout struct calculates the total price of all items in the cart, considering any special pricing rules.

Special Pricing: Certain items have special pricing rules defined, such as "3 for 130" or "2 for 45".

Error Handling

Error handling has been implemented for invalid item names, missing special price configurations, and negative prices.

Structure

The code is organised into separate files (main.go, checkout.go, and item.go) for better maintainability and readability.
