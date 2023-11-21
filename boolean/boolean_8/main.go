package main

import "fmt"

func main(){
	var a, b int
	var c bool

	fmt.Println("Input first number: ")
	fmt.Scan(&a)

	fmt.Println("Input second number: ")
	fmt.Scan(&b)

	c = a % 2 != 0 && b % 2 != 0

	fmt.Println("(", a, "is odd and", b, "is odd) is", c)
}