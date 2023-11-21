package main

import "fmt"

func main() {
	var a, x, y int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	x = a % 10
	y = a / 10 % 10

	fmt.Println("Last digit is: ", x, "\nSecond digit is: ", y)
}
