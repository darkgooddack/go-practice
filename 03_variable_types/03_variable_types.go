package main

import (
	"fmt"
)

func main() {

	name := "Americano"
	var price float64 = 2.50
	ready := true
	count := 5

	fmt.Printf("Type of name is: %T\n", name)
	fmt.Printf("Type of price is: %T\n", price)
	fmt.Printf("Type of ready is: %T\n", ready)
	fmt.Printf("Type of count is: %T\n", count)

	var customerName string
	var totalOrderQty int
	var isOrderReady bool

	fmt.Println("Type of customer is:", customerName)
	fmt.Println("Type of totalOrderQty is:", totalOrderQty)
	fmt.Println("Type of isOrderReady is:", isOrderReady)

}
