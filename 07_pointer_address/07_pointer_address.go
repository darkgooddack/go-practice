package main

import "fmt"

func main() {
	coffee := "Espresso"
	pointer := &coffee

	fmt.Println("Coffee: ", coffee)
	fmt.Println("Memory location", pointer)
	fmt.Printf("Pointer address: %p\n", pointer)

	fmt.Println("-------------------------")

	coffeeCopy := coffee
	pointerCoffeeCopy := &coffeeCopy

	fmt.Println("CoffeeCopy: ", coffeeCopy)
	fmt.Println("Memory location", pointerCoffeeCopy)
	fmt.Printf("Pointer address: %p\n", pointerCoffeeCopy)
}
