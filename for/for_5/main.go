package main

import "fmt"

func main() {
	var price float64

	fmt.Println("Input price of 1kg: ")
	fmt.Scan(&price)

	for i := 0.1; i <= 1.0; i += 0.1 {
		fmt.Println(i, "kg = ", price*i)
	}
}
