package main

import "fmt"

func main(){
	var a, b, c int
	var d bool

	fmt.Println("Input first number: ")
	fmt.Scan(&a)

	fmt.Println("Input second number: ")
	fmt.Scan(&b)

	fmt.Println("Input third number: ")
	fmt.Scan(&c)

	d = (b > a && b < c)

	fmt.Println("(", b, ">", a, "&&", b, "<", c, ") is", d)
}