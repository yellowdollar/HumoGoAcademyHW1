package main

import "fmt"

func main() {
	var n int

	fmt.Println("Input n: ")
	fmt.Scan(&n)

	var sum float64 = 0.0
	for i := 1; i <= n; i++ {
		sum += 1 / float64(i)
	}

	fmt.Println(sum)

}
