package main

import "fmt"

func main() {
	var a, x, y int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	x = a / 10
	y = a % 10

	fmt.Println("X + Y = ", x+y)
	fmt.Println("X * Y = ", x*y)
}
