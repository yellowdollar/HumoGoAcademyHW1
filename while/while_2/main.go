package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	result := 0

	for a >= b {
		a -= b
		result++
	}

	fmt.Println(result)
}
