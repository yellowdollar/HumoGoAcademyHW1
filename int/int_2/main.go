package main

import "fmt"

func main() {
	var m, res int

	fmt.Println("Input M: ")
	fmt.Scan(&m)

	res = m / 1000

	fmt.Println(res, "Ton")
}
