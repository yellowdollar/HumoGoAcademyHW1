package main

import "fmt"

func main() {
	var price float64

	fmt.Println("Input price of 1kg: ")
	fmt.Scan(&price)

	for i := 1.2; i <= 2.0; i += 0.2 {
		fmt.Println(i, "kg = ", price*i)
	}
}
