package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	for a >= b {
		a -= b
	}

	fmt.Println("Size of free space is: ", a)
}
