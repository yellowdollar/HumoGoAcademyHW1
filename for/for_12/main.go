package main

import "fmt"

func main() {
	var n int

	fmt.Print("Input n: ")
	fmt.Scan(&n)

	var result float64 = 1.0

	for i := 1; i <= n; i++ {
		result *= float64(i+10) / 10.0
	}
	fmt.Println(result)
}
