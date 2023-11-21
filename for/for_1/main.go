package main

import "fmt"

func main() {
	var n, k int

	fmt.Println("Input n: ")
	fmt.Scan(&n)

	fmt.Println("Input k: ")
	fmt.Scan(&k)

	fmt.Println("\nOutput \n")
	for i := 1; i <= n; i++ {
		fmt.Println(k)
	}

}
