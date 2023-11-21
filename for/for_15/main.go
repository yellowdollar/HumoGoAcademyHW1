package main

import "fmt"

func main() {
	var a, n int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input n: ")
	fmt.Scan(&n)

	powa := 1
	for i := 0; i < n; i++ {
		powa *= a
	}

	fmt.Println(powa)
}
