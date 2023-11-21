package main

import "fmt"

func main() {
	var a, x int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	x = a / 100

	fmt.Println("X = ", x)
}
