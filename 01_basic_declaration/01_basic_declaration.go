package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	fmt.Println("Hello World")
	color.Yellow("Hello World v2")

	var info string = "Medium Espresso price is $3.00"

	fmt.Println(info)

	info = "Large Espresso price is $5.00"

	color.Green(info)
	color.Green(info)

}
