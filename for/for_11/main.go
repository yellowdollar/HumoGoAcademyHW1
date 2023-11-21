package main

import "fmt"

func main() {
	var n int

	fmt.Println("Input n: ")
	fmt.Scan(&n)

	sum := 0
	for i := 0; i <= n; i++ {
		sum = (n * (n + 1) * (2*n + 1)) / 3
	}

	fmt.Println(sum)
}
