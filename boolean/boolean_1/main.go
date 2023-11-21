package main

import "fmt"

func main(){
	var b bool
	var a  int

	fmt.Println("Input number: ")
	fmt.Scan(&a)

	b = a > 0

	fmt.Println(b)
}