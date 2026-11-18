package main

import (
	"fmt"

	"example.com/pack/user"
)

func main() {
	fmt.Println("check user")
	u, err := user.New("jaggi", "shah")
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	fmt.Println("User created:", u)
}
