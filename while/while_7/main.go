package main

import "fmt"

func main() {
	var n int

	fmt.Println("Input n: ")
	fmt.Scan(&n)

	k := 0

	for k*k <= n {
		k++
	}

	fmt.Println(k, "||", k, "*", k, " = ", k*k)

}
