package main

import (
	"fmt"
)

func main() {
	var var_1 float64
	var_2 := 3.0

	fmt.Print("Enter Value: ")
	fmt.Scan(&var_1)

	fmt.Println("Second Value:", var_2)
	fmt.Println(var_1 + var_2)

}
