package main

import "fmt"

func main() {
	var b, res int

	fmt.Println("Input Bytes: ")
	fmt.Scan(&b)

	res = b / 1024

	fmt.Println(res)
}
