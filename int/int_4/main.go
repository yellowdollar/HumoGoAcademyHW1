package main

import "fmt"

func main() {
	var a, b, res int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	if a > b {
		res = a / b
	} else {
		fmt.Println("A should be bigger than B, restart code")
	}

	fmt.Println("In A we have ", res, " B's ")
}
