package main

import "fmt"

func main() {
	var n int

	fmt.Println("Input n: ")
	fmt.Scan(&n)

	k := 1
	for k*k <= n {
		k++
	}

	fmt.Println(k-1, "||", k-1, "*", k-1, " = ", (k-1)*(k-1))
}
