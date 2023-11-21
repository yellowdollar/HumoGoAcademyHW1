package main

import "fmt"

func main() {
	var n float64

	fmt.Println("Input n: ")
	fmt.Scan(&n)

	result := 1.0
	for n > 0 {
		result *= n
		n -= 2
	}

	fmt.Println(result)
}
