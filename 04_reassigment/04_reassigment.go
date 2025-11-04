package main

import "fmt"

func main() {

	// Explicit type
	var cupsQty int = 3
	// cupsQty = "four" - Compile-time error
	fmt.Println("Initial cupsQty:", cupsQty)

	// Implicit type
	var wasProcessed = true
	fmt.Println("Was processed:", wasProcessed)

	// Declaration
	var price float64
	// Initialization
	price = 3.50
	fmt.Printf("Price is %.2f\n", price)
}
