package main

import (
	"fmt"
)

func main() {
	var coffeeName string = "Espresso"

	var size = "Small"

	price := 2.50

	fmt.Println(coffeeName, size, "price is $", price)
	fmt.Printf("%s %s price is $%.2f\n", coffeeName, size, price)

}
