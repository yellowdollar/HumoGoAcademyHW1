package main

import "fmt"

func main() {
	var n int

	fmt.Println("Input n: ")
	fmt.Scan(&n)

	for n > 1 && n%3 == 0 {
		n /= 3
	}

	if n == 1 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
