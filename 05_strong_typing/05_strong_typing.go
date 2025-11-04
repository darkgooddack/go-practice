package main

import "fmt"

func main() {

	price := 4.50

	quntity := 15

	// total := price * quntity - error

	total := float64(quntity) * price
	fmt.Printf("Total is %.2f\n", total)

}
