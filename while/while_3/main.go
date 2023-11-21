package main

import "fmt"

func main() {
	var n, k int

	fmt.Println("Input n: ")
	fmt.Scan(&n)

	fmt.Println("Input k: ")
	fmt.Scan(&k)

	s := 0
	for n >= k {
		n -= k
		s++
	}

	fmt.Println("The quotient of the integral division of N by K: ", s)
	fmt.Println("The remainder of an even division of N by K: ", n)
}
