package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	sum := 0
	for i := a; i <= b; i++ {
		sum += (i * i)
	}

	fmt.Println(sum)
}
