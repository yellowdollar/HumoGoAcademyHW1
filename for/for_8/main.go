package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	substr := 1

	for i := a; i <= b; i++ {
		substr *= i
	}

	fmt.Println(substr)
}
