package main

import "fmt"

func main() {
	var n int

	fmt.Println("Input n: ")
	fmt.Scan(&n)

	k := 0
	for n > 0 && n%2 == 0 {
		n /= 2
		k++
	}

	fmt.Println(k)
}
