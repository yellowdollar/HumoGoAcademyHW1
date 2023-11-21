package main

import "fmt"

func main(){
	var a, b int
	var c bool

	fmt.Println("Input first number: ")
	fmt.Scan(&a)

	fmt.Println("Input second number: ")
	fmt.Scan(&b)

	c = a > 2  && b <= 3

	fmt.Println("(", a, "> 2 &&", b, "<= 3 ) is", c)
}