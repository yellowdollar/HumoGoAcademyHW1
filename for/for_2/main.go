package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	n = b - a + 1

	for i := a; i <= b; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n", n)
}
